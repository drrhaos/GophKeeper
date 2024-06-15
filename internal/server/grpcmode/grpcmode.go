// Package grpcmode запускает сервер.
package grpcmode

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"time"

	"gophkeeper/internal/logger"
	"gophkeeper/internal/server/configure"
	"gophkeeper/internal/store"
	"gophkeeper/internal/store/pg"

	pb "gophkeeper/pkg/proto"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/jwtauth"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)


type UserClaims struct {
    jwt.StandardClaims
    Username string
}

var ErrNotValidToken = errors.New("token not valid") // ErrNotValidToken токен не прошел проверку.

// Run запускает сервер
func Run(cfg configure.Config) {
	pg.Migrations(cfg.DatabaseDsn)

	storeKeeper := &store.StorageContext{}
	storeKeeper.SetStorage(pg.NewDatabase(cfg.DatabaseDsn))

	listen, err := net.Listen("tcp", cfg.Address)
	if err != nil {
		log.Fatal(err)
	}

	// создаём gRPC-сервер без зарегистрированной службы
	s := grpc.NewServer()

	// регистрируем сервис
	metricsServer := GophKeeperServer{
		storage: storeKeeper,
		cfg:     cfg,
	}
	pb.RegisterGophKeeperServer(s, &metricsServer)

	// получаем запрос gRPC
	if err := s.Serve(listen); err != nil {
		log.Fatal(err)
	}
}

// GophKeeperServer поддерживает все необходимые методы сервера.
type GophKeeperServer struct {
	pb.UnimplementedGophKeeperServer

	storage *store.StorageContext
	cfg     configure.Config
}

// Register регистрирует нового пользователя.
func (ms *GophKeeperServer) Register(ctx context.Context, in *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	var response pb.RegisterResponse

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if in.Login == "" || in.Password == "" {
		return &response, nil
	}

	err := ms.storage.UserRegister(ctx, in.Login, in.Password)
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

	response.Token, err = token.SignedString([]byte(ms.cfg.SecretKey))
	if err != nil {
		logger.Log.Warn("Ошибка создания токена:", zap.Error(err))
		return &response, err
	}

	logger.Log.Info("Новый пользователь зарегистрирован и атентифицирован")

	return &response, nil
}

// Login аутентифицирует нового пользователя.
func (ms *GophKeeperServer) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	var response pb.LoginResponse

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if in.Login == "" || in.Password == "" {
		return &response, nil
	}

	err := ms.storage.UserLogin(ctx, in.Login, in.Password)

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
    response.Token, err = token.SignedString([]byte(ms.cfg.SecretKey))
	
	if err != nil {
		logger.Log.Warn("Ошибка создания токена:", zap.Error(err))
		return &response, err
	}

	logger.Log.Info("Пользователь аутентифицирован")

	return &response, nil
}

// SyncData синхронизация данных.
func (ms *GophKeeperServer) SyncData(ctx context.Context, in *pb.SyncRequest) (*pb.SyncResponse, error) {
	var response pb.SyncResponse

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	
    claims := &UserClaims{}
    parsedToken, err := jwt.ParseWithClaims(in.GetToken(), claims,
    func(t *jwt.Token) (interface{}, error) {
        if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
        }
        return []byte(ms.cfg.SecretKey), nil
    })
	
	if err != nil || !parsedToken.Valid {
		logger.Log.Warn("Недействительный JWT-токен")
		return &response, ErrNotValidToken
	}

	ms.storage.SyncFields(ctx, claims.Username, in.GetData())

	logger.Log.Info("Данные синхронизированны")

	return &response, nil
}
