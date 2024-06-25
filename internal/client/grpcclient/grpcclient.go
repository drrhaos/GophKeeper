// Package grpcclient реализует интерфейс взаимодействия с GRPC сервером.
package grpcclient

import (
	"context"

	"gophkeeper/internal/client/configure"
	"gophkeeper/internal/logger"

	pb "gophkeeper/pkg/proto"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

// GRPCClient хранит соединение с сервером.
type GRPCClient struct {
	client pb.GophKeeperClient
	token  string
}

// NewGRPCClient устанавливает соединение с сервером.
func NewGRPCClient(cfg configure.Config, user string, password string) (*GRPCClient, error) {
	conn, err := grpc.NewClient(cfg.Address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Log.Warn("Не удалось установить соединение с сервером", zap.Error(err))
		return nil, err
	}

	client := pb.NewGophKeeperClient(conn)
	req := pb.LoginRequest{
		Login:    user,
		Password: password,
	}
	res, err := client.Login(context.Background(), &req)
	if err != nil {
		logger.Log.Warn("Не удалось авторизоваться", zap.Error(err))
		return nil, err
	}
	token := res.GetToken()

	return &GRPCClient{client: client, token: token}, nil
}

// GetListFields получает с сервера список запсисей.
func (client *GRPCClient) GetListFields() *pb.ListFielsdKeepResponse {
	md := metadata.New(map[string]string{"Authorization": client.token})
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	resp, err := client.client.ListFields(ctx, &pb.ListFieldsKeepRequest{})
	if err != nil {
		return nil
	}
	return resp
}

// SaveField сохранение хаписи на сервере.
func (client *GRPCClient) SaveField(field *pb.EditFieldKeepRequest) (*pb.EditFieldKeepResponse, error) {
	md := metadata.New(map[string]string{"Authorization": client.token})
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	return client.client.EditField(ctx, field)
}
