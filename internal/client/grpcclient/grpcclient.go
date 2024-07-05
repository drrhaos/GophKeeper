// Package grpcclient реализует интерфейс взаимодействия с GRPC сервером.
package grpcclient

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
	"path/filepath"

	"gophkeeper/internal/client/configure"
	"gophkeeper/internal/crypt"
	"gophkeeper/internal/server/grpcmode"

	"gophkeeper/pkg/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

// GRPCClient хранит соединение с сервером.
type GRPCClient struct {
	client proto.GophKeeperClient
	token  string
	cfg    configure.Config
	secret string
}

// Connect устанавливает соединение с сервером.
func Connect(cfg configure.Config, user string, password string) (*GRPCClient, error) {
	conn, err := grpc.NewClient(cfg.Address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client := proto.NewGophKeeperClient(conn)
	req := proto.LoginRequest{
		Login:    user,
		Password: password,
	}
	res, err := client.Login(context.Background(), &req)
	if err != nil {
		return nil, err
	}
	token := res.GetToken()

	hash := md5.Sum([]byte(cfg.Secret))
	secret := hex.EncodeToString(hash[:])

	gClient := &GRPCClient{
		client: client,
		token:  token,
		cfg:    cfg,
		secret: secret,
	}
	return gClient, nil
}

// Reg регистрирует нового польщователя.
func Reg(cfg configure.Config, user string, password string) (*GRPCClient, error) {
	conn, err := grpc.NewClient(cfg.Address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client := proto.NewGophKeeperClient(conn)
	req := proto.RegisterRequest{
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
func (client *GRPCClient) GetListFields() *proto.ListFielsdKeepResponse {
	md := metadata.New(map[string]string{"Authorization": client.token})
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	respEnc, err := client.client.ListFields(ctx, &proto.ListFieldsKeepRequest{})
	if err != nil {
		return nil
	}

	respDec := &proto.ListFielsdKeepResponse{
		Data: make(map[string]*proto.FieldKeep),
	}

	for key, ddd := range respEnc.Data {
		respDec.GetData()[key] = crypt.DecField(ddd, client.secret)
	}

	return respDec
}

// SaveField сохранение записи на сервере.
func (client *GRPCClient) SaveField(field *proto.EditFieldKeepRequest) (*proto.EditFieldKeepResponse, error) {
	md := metadata.New(map[string]string{"Authorization": client.token})
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	fieldEnc := &proto.EditFieldKeepRequest{
		Uuid: field.Uuid,
		Data: crypt.EncField(field.Data, client.secret),
	}

	respEnc, err := client.client.EditField(ctx, fieldEnc)
	if err != nil {
		return nil, err
	}

	respDec := &proto.EditFieldKeepResponse{
		Data: crypt.DecField(respEnc.GetData(), client.secret),
	}

	return respDec, nil
}

// AddField добавление записи на сервере.
func (client *GRPCClient) AddField(field *proto.AddFieldKeepRequest) (*proto.AddFieldKeepResponse, error) {
	md := metadata.New(map[string]string{"Authorization": client.token})
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	fieldEnc := &proto.AddFieldKeepRequest{
		Data: crypt.EncField(field.Data, client.secret),
	}

	respEnc, err := client.client.AddField(ctx, fieldEnc)
	if err != nil {
		return nil, err
	}

	respDec := &proto.AddFieldKeepResponse{
		Uuid: respEnc.Uuid,
		Data: crypt.DecField(respEnc.GetData(), client.secret),
	}
	return respDec, nil
}

// DelField удаление записи на сервере.
func (client *GRPCClient) DelField(field *proto.DeleteFieldKeepRequest) (*proto.DeleteFieldKeepResponse, error) {
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

		if err := stream.Send(&proto.FileUploadRequest{FileName: filepath.Base(filePath), Chunk: chunk}); err != nil {
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

	stream, err := client.client.Download(ctx, &proto.FileDownRequest{Uuid: uuid, FileName: fileName})
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
