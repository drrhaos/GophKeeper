// Package grpcclient реализует интерфейс взаимодействия с GRPC сервером.
package grpcclient

import (
	"context"
	"net"
	"os"
	"path/filepath"
	"testing"

	cClient "gophkeeper/internal/client/configure"
	cServer "gophkeeper/internal/server/configure"
	"gophkeeper/internal/server/grpcmode"
	"gophkeeper/internal/store"
	"gophkeeper/internal/store/mocks"
	"gophkeeper/pkg/proto"

	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestConnect(t *testing.T) {
	dirPath := "/tmp/"
	storeKeeper := &store.StorageContext{}
	mockStore := new(mocks.MockStore)
	mockStore.On("UserLogin", mock.Anything, "test", "test").Return(nil)
	storeKeeper.SetStorage(mockStore)

	server := grpc.NewServer()
	proto.RegisterGophKeeperServer(server, &grpcmode.GophKeeperServer{
		Storage: storeKeeper,
		Cfg:     cServer.Config{WorkPath: dirPath},
	})

	go func() {
		listen, _ := net.Listen("tcp", ":50051")
		server.Serve(listen)
	}()
	defer server.Stop()

	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	_, err = Connect(cClient.Config{Address: "127.0.0.1:50051"}, "test", "test")
	if err != nil {
		t.Error(err)
	}
}

func TestGRPCClient_GetListFields(t *testing.T) {
	dirPath := "/tmp/"
	storeKeeper := &store.StorageContext{}
	mockStore := new(mocks.MockStore)
	mockStore.On("UserLogin", mock.Anything, "test", "test").Return(nil)
	field := proto.FieldKeep{
		Login:    "Qq+R8v/cVxTOZ9WWSr8tFYukEN+UMtOrEvcSWl9Cz+Gy",
		Password: "ezARybygDwVMLCQ8Gsbker3jZ+bcwpUDJjzqxtPYIzLBZLxG",
	}
	resp := &proto.ListFielsdKeepResponse{
		Data: make(map[string]*proto.FieldKeep),
	}

	resp.Data["6666"] = &field
	mockStore.On("ListFields", mock.Anything, "test").Return(resp, true)
	storeKeeper.SetStorage(mockStore)

	server := grpc.NewServer()
	proto.RegisterGophKeeperServer(server, &grpcmode.GophKeeperServer{
		Storage: storeKeeper,
		Cfg:     cServer.Config{WorkPath: dirPath},
	})

	go func() {
		listen, _ := net.Listen("tcp", ":50051")
		err := server.Serve(listen)
		if err != nil {
			t.Error(err)
		}
	}()
	defer server.Stop()

	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	gClient, err := Connect(cClient.Config{Address: "127.0.0.1:50051", Secret: "test"}, "test", "test")
	if err != nil {
		t.Error(err)
	}
	respList := gClient.GetListFields()
	if respList.GetData()["6666"].Login != "login" {
		t.Errorf("No valid data %s != %s", respList.GetData()["6666"].Login, "")
	}
}

func TestGRPCClient_Upload(t *testing.T) {
	fileTmp := "/tmp/6666"
	defer os.Remove(fileTmp)
	fileTmpOut := "/tmp/66667"
	defer os.Remove(fileTmpOut)
	f, err := os.Create(fileTmp)
	if err != nil {
		t.Error(err)
	}
	defer f.Close()
	for i := 0; i < 1000; i++ {
		_, err = f.Write([]byte("test "))
		if err != nil {
			t.Error(err)
		}
	}

	dirPath, _ := filepath.Split(fileTmp)

	storeKeeper := &store.StorageContext{}
	mockStore := new(mocks.MockStore)
	mockStore.On("UserLogin", mock.Anything, "test", "test").Return(nil)
	storeKeeper.SetStorage(mockStore)

	server := grpc.NewServer()
	proto.RegisterGophKeeperServer(server, &grpcmode.GophKeeperServer{
		Storage: storeKeeper,
		Cfg:     cServer.Config{WorkPath: dirPath},
	})

	go func() {
		listen, _ := net.Listen("tcp", ":50051")
		err = server.Serve(listen)
		if err != nil {
			t.Error(err)
		}
	}()
	defer server.Stop()

	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	gClient, err := Connect(cClient.Config{Address: "127.0.0.1:50051", Secret: "test"}, "test", "test")
	if err != nil {
		t.Error(err)
	}
	err = gClient.Upload(context.Background(), fileTmp)
	if err != nil {
		t.Error(err)
	}
}

func TestGRPCClient_Download(t *testing.T) {
	fileTmp := "/tmp/6666"
	defer os.Remove(fileTmp)
	fileTmpOut := "/tmp/66667"
	defer os.Remove(fileTmpOut)
	f, err := os.Create(fileTmp)
	if err != nil {
		t.Error(err)
	}
	defer f.Close()
	for i := 0; i < 1000; i++ {
		_, err := f.Write([]byte("test "))
		if err != nil {
			t.Error(err)
		}
	}

	dirPath, _ := filepath.Split(fileTmp)

	storeKeeper := &store.StorageContext{}
	mockStore := new(mocks.MockStore)
	mockStore.On("UserLogin", mock.Anything, "test", "test").Return(nil)
	storeKeeper.SetStorage(mockStore)

	server := grpc.NewServer()
	proto.RegisterGophKeeperServer(server, &grpcmode.GophKeeperServer{
		Storage: storeKeeper,
		Cfg:     cServer.Config{WorkPath: dirPath},
	})

	go func() {
		listen, _ := net.Listen("tcp", ":50051")
		err = server.Serve(listen)
		if err != nil {
			t.Error(err)
		}
	}()
	defer server.Stop()

	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	gClient, err := Connect(cClient.Config{Address: "127.0.0.1:50051", Secret: "test", StaticPath: dirPath}, "test", "test")
	if err != nil {
		t.Error(err)
	}
	err = gClient.Download(context.Background(), "6666", "test")
	if err != nil {
		t.Error(err)
	}
}
