// Package grpcmode запускает сервер.
package grpcmode

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"
	"time"

	"gophkeeper/internal/logger"
	"gophkeeper/internal/server/configure"
	"gophkeeper/internal/store"
	"gophkeeper/internal/store/pg"

	"gophkeeper/pkg/proto"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/jwtauth"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

// UserClaims JWT структура
type UserClaims struct {
	jwt.StandardClaims
	Username string
}

var (
	ErrNotValidToken = errors.New("token not valid") // ErrNotValidToken токен не прошел проверку.
	ErrNotValidData  = errors.New("data not valid")  // ErrNotValidData не верный формат данных.
)

// Run запускает сервер
func Run(cfg configure.Config) {
	pg.Migrations(cfg)

	tlsCreds, err := credentials.NewServerTLSFromFile(cfg.CertFile, cfg.KeyFile)
	if err != nil {
		logger.Log.Panic("Не удалось создать сертификаты", zap.Error(err))
	}

	storeKeeper := &store.StorageContext{}
	storeKeeper.SetStorage(pg.NewDatabase(cfg.DatabaseDsn))

	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.Port))
	if err != nil {
		logger.Log.Panic("Не удалось открыть порт", zap.Error(err))
	}

	// создаём gRPC-сервер без зарегистрированной службы
	s := grpc.NewServer(grpc.Creds(tlsCreds))

	// регистрируем сервис
	metricsServer := GophKeeperServer{
		Storage: storeKeeper,
		Cfg:     cfg,
	}
	proto.RegisterGophKeeperServer(s, &metricsServer)

	// получаем запрос gRPC
	if err := s.Serve(listen); err != nil {
		logger.Log.Panic("Не удалось запустить сервер", zap.Error(err))
	}
}

// GophKeeperServer поддерживает все необходимые методы сервера.
type GophKeeperServer struct {
	proto.UnimplementedGophKeeperServer

	Storage *store.StorageContext
	Cfg     configure.Config
}

// Register регистрирует нового пользователя.
func (ms *GophKeeperServer) Register(ctx context.Context, in *proto.RegisterRequest) (*proto.RegisterResponse, error) {
	var response proto.RegisterResponse

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if in.Login == "" || in.Password == "" {
		return &response, nil
	}

	err := ms.Storage.UserRegister(ctx, in.Login, in.Password)
	if errors.Is(err, store.ErrLoginDuplicate) {
		return &response, err
	} else if err != nil && !errors.Is(err, store.ErrLoginDuplicate) {
		return &response, err
	}

	claims := jwt.MapClaims{
		"username": in.Login,
	}

	jwtauth.SetExpiry(claims, time.Now().Add(time.Minute*5))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	response.Token, err = token.SignedString([]byte(ms.Cfg.SecretKey))
	if err != nil {
		logger.Log.Warn("Ошибка создания токена:", zap.Error(err))
		return &response, err
	}

	logger.Log.Info("Пользователь аутентифицирован")

	logger.Log.Info("Новый пользователь зарегистрирован и атентифицирован")

	return &response, nil
}

// Login аутентифицирует нового пользователя.
func (ms *GophKeeperServer) Login(ctx context.Context, in *proto.LoginRequest) (*proto.LoginResponse, error) {
	var response proto.LoginResponse

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if in.Login == "" || in.Password == "" {
		return &response, nil
	}

	err := ms.Storage.UserLogin(ctx, in.Login, in.Password)

	if errors.Is(err, store.ErrAuthentication) {
		return &response, err
	} else if err != nil && !errors.Is(err, store.ErrLoginDuplicate) {
		return &response, err
	}

	claims := &UserClaims{
		Username: in.GetLogin(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 3).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	response.Token, err = token.SignedString([]byte(ms.Cfg.SecretKey))
	if err != nil {
		logger.Log.Warn("Ошибка создания токена:", zap.Error(err))
		return &response, err
	}

	logger.Log.Info("Пользователь аутентифицирован")

	return &response, nil
}

// AddField добавдяет запись в хранилище.
func (ms *GophKeeperServer) AddField(ctx context.Context, in *proto.AddFieldKeepRequest) (*proto.AddFieldKeepResponse, error) {
	var response proto.AddFieldKeepResponse
	claims, err := ms.checkToken(ctx)
	if err != nil {
		return &response, err
	}

	var ok bool
	response.Uuid, response.Data, ok = ms.Storage.AddField(ctx, claims.Username, in.GetData())
	if !ok {
		return &response, ErrNotValidData
	}

	logger.Log.Info("Данные добавлены")

	return &response, nil
}

// EditField изменяет запись в хранилище.
func (ms *GophKeeperServer) EditField(ctx context.Context, in *proto.EditFieldKeepRequest) (*proto.EditFieldKeepResponse, error) {
	var response proto.EditFieldKeepResponse
	claims, err := ms.checkToken(ctx)
	if err != nil {
		return &response, err
	}
	var ok bool
	response.Data, ok = ms.Storage.EditField(ctx, claims.Username, in.GetUuid(), in.GetData())
	if !ok {
		return &response, ErrNotValidData
	}
	logger.Log.Info("Данные изменены")

	return &response, nil
}

// DelField удаляет запись из хранилища.
func (ms *GophKeeperServer) DelField(ctx context.Context, in *proto.DeleteFieldKeepRequest) (*proto.DeleteFieldKeepResponse, error) {
	var response proto.DeleteFieldKeepResponse
	claims, err := ms.checkToken(ctx)
	if err != nil {
		return &response, err
	}

	uuid, ok := ms.Storage.DelField(ctx, claims.Username, in.GetUuid())
	if !ok {
		return &response, ErrNotValidData
	}
	response.Uuid = uuid
	logger.Log.Info("Запись удалена")

	return &response, nil
}

// ListFields возвращает список записей.
func (ms *GophKeeperServer) ListFields(ctx context.Context, _ *proto.ListFieldsKeepRequest) (*proto.ListFielsdKeepResponse, error) {
	var response *proto.ListFielsdKeepResponse
	claims, err := ms.checkToken(ctx)
	if err != nil {
		return response, err
	}
	var ok bool
	response, ok = ms.Storage.ListFields(ctx, claims.Username)
	if !ok {
		return response, ErrNotValidData
	}

	logger.Log.Info("Данные получены")

	return response, nil
}

func (ms *GophKeeperServer) checkToken(ctx context.Context) (*UserClaims, error) {
	var token string
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		values := md.Get("Authorization")
		if len(values) > 0 {
			token = values[0]
		}
	}

	claims := &UserClaims{}
	parsedToken, err := jwt.ParseWithClaims(token, claims,
		func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}
			return []byte(ms.Cfg.SecretKey), nil
		})

	if err != nil || !parsedToken.Valid {
		logger.Log.Warn("Недействительный JWT-токен")
		return claims, ErrNotValidToken
	}
	return claims, nil
}

// Upload загрузка файла на сервер
func (ms *GophKeeperServer) Upload(stream proto.GophKeeper_UploadServer) error {
	_, err := ms.checkToken(stream.Context())
	if err != nil {
		return err
	}

	file := NewFile()
	var fileSize uint32
	fileSize = 0
	defer func() {
		if err := file.Close(); err != nil {
			logger.Log.Warn("Ошибка при закрытии файла", zap.Error(err))
		}
	}()
	for {
		req, err := stream.Recv()
		if file.FilePath == "" {
			errSetFile := file.SetFile(req.GetFileName(), ms.Cfg.WorkPath)
			if errSetFile != nil {
				return errSetFile
			}
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		chunk := req.GetChunk()
		fileSize += uint32(len(chunk))
		// g.l.Debug("received a chunk with size: %d", fileSize)
		if err := file.Write(chunk); err != nil {
			return err
		}
	}
	fileName := filepath.Base(file.FilePath)
	logger.Log.Info(fmt.Sprintf("saved file: %s, size: %d", fileName, fileSize))
	return stream.SendAndClose(&proto.FileUploadResponse{FileName: fileName, Size: fileSize})
}

// Download выгрузка файла с сервера.
func (ms *GophKeeperServer) Download(req *proto.FileDownRequest, stream proto.GophKeeper_DownloadServer) error {
	_, err := ms.checkToken(stream.Context())
	if err != nil {
		return err
	}

	fileName := req.GetFileName()
	uuid := req.GetUuid()
	path := filepath.Join(ms.Cfg.WorkPath, uuid)

	fileInfo, err := os.Stat(path)
	if err != nil {
		return err
	}
	fileSize := fileInfo.Size()

	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	var totalBytesStreamed int64

	for totalBytesStreamed < fileSize {
		shard := make([]byte, 1024)
		bytesRead, err := f.Read(shard)
		if err == io.EOF {
			logger.Log.Info(fmt.Sprintf("return file: %s, size: %d", fileName, fileSize))
			break
		}

		if err != nil {
			return err
		}

		if err := stream.Send(&proto.FileDownResponse{
			Chunk: shard,
		}); err != nil {
			return err
		}
		totalBytesStreamed += int64(bytesRead)
	}
	return nil
}
