// Package grpcmode запускает сервер.
package grpcmode

import (
	"context"
	"errors"
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
	"google.golang.org/grpc"
)

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
}

// Register обновляет значение пачки метрики
func (ms *GophKeeperServer) Register(ctx context.Context, in *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	var response pb.RegisterResponse
	var tokenAuth *jwtauth.JWTAuth

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
		"exp":      time.Now().Add(time.Hour).Unix(),
	}

	_, response.Token, err = tokenAuth.Encode(claims)
	if err != nil {
		logger.Log.Warn("Произошла ошибка генерации токена")
		return &response, err
	}

	logger.Log.Info("Новый пользователь зарегистрирован и атентифицирован")

	return &response, nil
}
