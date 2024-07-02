// Package grpcclient реализует интерфейс взаимодействия с GRPC сервером.
package grpcclient

import (
	"context"
	"io"
	"os"
	"path/filepath"

	"gophkeeper/internal/client/configure"
	"gophkeeper/internal/server/grpcmode"

	pb "gophkeeper/pkg/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

// GRPCClient хранит соединение с сервером.
type GRPCClient struct {
	client pb.GophKeeperClient
	token  string
	cfg    configure.Config
}

// Connect устанавливает соединение с сервером.
func Connect(cfg configure.Config, user string, password string) (*GRPCClient, error) {
	conn, err := grpc.NewClient(cfg.Address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client := pb.NewGophKeeperClient(conn)
	req := pb.LoginRequest{
		Login:    user,
		Password: password,
	}
	res, err := client.Login(context.Background(), &req)
	if err != nil {
		return nil, err
	}
	token := res.GetToken()

	return &GRPCClient{client: client, token: token, cfg: cfg}, nil
}

// Reg регистрирует нового польщователя.
func Reg(cfg configure.Config, user string, password string) (*GRPCClient, error) {
	conn, err := grpc.NewClient(cfg.Address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client := pb.NewGophKeeperClient(conn)
	req := pb.RegisterRequest{
		Login:    user,
		Password: password,
	}
	res, err := client.Register(context.Background(), &req)
	if err != nil {
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

// SaveField сохранение записи на сервере.
func (client *GRPCClient) SaveField(field *pb.EditFieldKeepRequest) (*pb.EditFieldKeepResponse, error) {
	md := metadata.New(map[string]string{"Authorization": client.token})
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	return client.client.EditField(ctx, field)
}

// AddField добавление записи на сервере.
func (client *GRPCClient) AddField(field *pb.AddFieldKeepRequest) (*pb.AddFieldKeepResponse, error) {
	md := metadata.New(map[string]string{"Authorization": client.token})
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	return client.client.AddField(ctx, field)
}

// DelField удаление записи на сервере.
func (client *GRPCClient) DelField(field *pb.DeleteFieldKeepRequest) (*pb.DeleteFieldKeepResponse, error) {
	md := metadata.New(map[string]string{"Authorization": client.token})
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	return client.client.DelField(ctx, field)
}

// Upload загрузка файла на сервер.
func (client *GRPCClient) Upload(ctx context.Context, filePath string) error {
	md := metadata.New(map[string]string{"Authorization": client.token})
	ctx = metadata.NewOutgoingContext(ctx, md)
	stream, err := client.client.Upload(ctx)
	if err != nil {
		return err
	}
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	buf := make([]byte, 1024)
	batchNumber := 1

	for {
		num, err := file.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		chunk := buf[:num]

		if err := stream.Send(&pb.FileUploadRequest{FileName: filepath.Base(filePath), Chunk: chunk}); err != nil {
			return err
		}
		batchNumber++

	}
	_, err = stream.CloseAndRecv()
	if err != nil {
		return err
	}
	return nil
}

// Download загрузка файла с сервера.
func (client *GRPCClient) Download(ctx context.Context, uuid string, fileName string) error {
	md := metadata.New(map[string]string{"Authorization": client.token})
	ctx = metadata.NewOutgoingContext(ctx, md)

	stream, err := client.client.Download(ctx, &pb.FileDownRequest{Uuid: uuid, FileName: fileName})
	if err != nil {
		return err
	}

	file := grpcmode.NewFile()
	var fileSize uint32
	fileSize = 0
	defer func() {
		file.Close()
	}()
	for {
		req, err := stream.Recv()
		if file.FilePath == "" {
			file.SetFile(uuid, client.cfg.StaticPath)
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		chunk := req.GetChunk()
		fileSize += uint32(len(chunk))
		if err := file.Write(chunk); err != nil {
			return err
		}
	}
	return nil
}
