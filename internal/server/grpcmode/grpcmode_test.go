// Package grpcmode запускает сервер.
package grpcmode

import (
	"context"
	"gophkeeper/internal/server/configure"
	"gophkeeper/internal/store"
	"gophkeeper/internal/store/mocks"
	"gophkeeper/pkg/proto"
	"reflect"
	"testing"

	"github.com/stretchr/testify/mock"
)

var conf configure.Config

func TestGophKeeperServer_Register(t *testing.T) {
	storeKeeper := &store.StorageContext{}
	mockStore := new(mocks.MockStore)
	mockStore.On("UserRegister", mock.Anything, "test", "test").Return(store.ErrLoginDuplicate)
	mockStore.On("UserRegister", mock.Anything, "test2", "test").Return(nil)
	storeKeeper.SetStorage(mockStore)

	// conf.ReadConfig()
	type fields struct {
		UnimplementedGophKeeperServer proto.UnimplementedGophKeeperServer
		cfg                           configure.Config
	}
	type args struct {
		ctx context.Context
		in  *proto.RegisterRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *proto.RegisterResponse
		wantErr bool
	}{
		{
			name: "Positive #1",
			args: args{
				ctx: context.Background(),
				in: &proto.RegisterRequest{
					Login:    "test",
					Password: "test",
				},
			},
			fields: fields{
				UnimplementedGophKeeperServer: proto.UnimplementedGophKeeperServer{},
				cfg:                           conf,
			},
			want:    &proto.RegisterResponse{},
			wantErr: true,
		},
		{
			name: "Negative #2",
			args: args{
				ctx: context.Background(),
				in: &proto.RegisterRequest{
					Login:    "test2",
					Password: "test",
				},
			},
			fields: fields{
				UnimplementedGophKeeperServer: proto.UnimplementedGophKeeperServer{},
				cfg:                           conf,
			},
			want: &proto.RegisterResponse{
				Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjA3MzI0MDQsInVzZXJuYW1lIjoidGVzdDIifQ.wxicxzfUPXyReBwUQ5FdEMibB0_KNoRT2jOuhRrieI4",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ms := &GophKeeperServer{
				UnimplementedGophKeeperServer: tt.fields.UnimplementedGophKeeperServer,
				storage:                       storeKeeper,
				cfg:                           tt.fields.cfg,
			}
			got, err := ms.Register(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("GophKeeperServer.Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(len(got.Token), len(tt.want.Token)) {
				t.Errorf("GophKeeperServer.Register() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGophKeeperServer_Login(t *testing.T) {

	storeKeeper := &store.StorageContext{}
	mockStore := new(mocks.MockStore)
	mockStore.On("UserLogin", mock.Anything, "test", "test").Return(store.ErrAuthentication)
	mockStore.On("UserLogin", mock.Anything, "test2", "test").Return(nil)
	storeKeeper.SetStorage(mockStore)

	// conf.ReadConfig()

	type fields struct {
		UnimplementedGophKeeperServer proto.UnimplementedGophKeeperServer
		cfg                           configure.Config
	}
	type args struct {
		ctx context.Context
		in  *proto.LoginRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *proto.LoginResponse
		wantErr bool
	}{
		{
			name: "Negtive #1",
			args: args{
				ctx: context.Background(),
				in: &proto.LoginRequest{
					Login:    "test",
					Password: "test",
				},
			},
			fields: fields{
				UnimplementedGophKeeperServer: proto.UnimplementedGophKeeperServer{},
				cfg:                           conf,
			},
			want:    &proto.LoginResponse{},
			wantErr: true,
		},
		{
			name: "Positive #1",
			args: args{
				ctx: context.Background(),
				in: &proto.LoginRequest{
					Login:    "test2",
					Password: "test",
				},
			},
			fields: fields{
				UnimplementedGophKeeperServer: proto.UnimplementedGophKeeperServer{},
				cfg:                           conf,
			},
			want:    &proto.LoginResponse{Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjA3NDM5ODMsIlVzZXJuYW1lIjoidGVzdDIifQ.FW-NqOjZXZluWc9_hNw9D43gq9XJkL5EZAL2V22nDiU"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ms := &GophKeeperServer{
				UnimplementedGophKeeperServer: tt.fields.UnimplementedGophKeeperServer,
				storage:                       storeKeeper,
				cfg:                           tt.fields.cfg,
			}
			got, err := ms.Login(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("GophKeeperServer.Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(len(got.Token), len(tt.want.Token)) {
				t.Errorf("GophKeeperServer.Login() = %v, want %v", got, tt.want)
			}
		})
	}
}
