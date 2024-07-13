// Package grpcmode запускает сервер.
package grpcmode

import (
	"context"
	"reflect"
	"testing"

	"gophkeeper/internal/server/configure"
	"gophkeeper/internal/store"
	"gophkeeper/internal/store/mocks"
	"gophkeeper/pkg/proto"

	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc/metadata"
)

var conf configure.Config

func TestGophKeeperServer_Register(t *testing.T) {
	storeKeeper := &store.StorageContext{}
	mockStore := new(mocks.MockStore)
	mockStore.On("UserRegister", mock.Anything, "test", "test").Return(store.ErrLoginDuplicate)
	mockStore.On("UserRegister", mock.Anything, "test2", "test").Return(nil)
	storeKeeper.SetStorage(mockStore)

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

func TestGophKeeperServer_AddField(t *testing.T) {
	storeKeeper := &store.StorageContext{}
	mockStore := new(mocks.MockStore)
	mockStore.On("UserLogin", mock.Anything, "test", "test").Return(nil)
	prot := proto.FieldKeep{Login: "user", Password: "password"}
	mockStore.On("AddField", mock.Anything, "test", &prot).Return("63636363636636", &prot, true)
	storeKeeper.SetStorage(mockStore)

	type fields struct {
		UnimplementedGophKeeperServer proto.UnimplementedGophKeeperServer
		cfg                           configure.Config
	}
	type args struct {
		ctx context.Context
		in  *proto.AddFieldKeepRequest
	}
	tests := []struct {
		name     string
		user     string
		password string
		fields   fields
		args     args
		want     *proto.AddFieldKeepResponse
		wantErr  bool
	}{
		{
			name:     "Positive test #1",
			user:     "test",
			password: "test",
			fields: fields{
				UnimplementedGophKeeperServer: proto.UnimplementedGophKeeperServer{},
				cfg:                           conf,
			},
			args: args{
				ctx: context.Background(),
				in: &proto.AddFieldKeepRequest{
					Data: &proto.FieldKeep{
						Login:    "user",
						Password: "password",
					},
				},
			},
			want: &proto.AddFieldKeepResponse{
				Uuid: "63636363636636",
				Data: &proto.FieldKeep{
					Login:    "user",
					Password: "password",
				},
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

			resp, errLogin := ms.Login(tt.args.ctx, &proto.LoginRequest{Login: tt.user, Password: tt.password})
			if errLogin != nil {
				t.Errorf("Login error = %v", errLogin)
			}

			md := metadata.New(map[string]string{"Authorization": resp.Token})
			ctx := metadata.NewIncomingContext(context.Background(), md)

			got, err := ms.AddField(ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("GophKeeperServer.AddField() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.Uuid, tt.want.Uuid) {
				t.Errorf("GophKeeperServer.AddField() = %v, want %v", got.Uuid, tt.want.Uuid)
			}
		})
	}
}

func TestGophKeeperServer_EditField(t *testing.T) {
	storeKeeper := &store.StorageContext{}
	mockStore := new(mocks.MockStore)
	mockStore.On("UserLogin", mock.Anything, "test", "test").Return(nil)
	prot := proto.FieldKeep{Login: "user", Password: "password1"}
	mockStore.On("EditField", mock.Anything, "test", "6666", &prot).Return(&prot, true)
	storeKeeper.SetStorage(mockStore)

	type fields struct {
		UnimplementedGophKeeperServer proto.UnimplementedGophKeeperServer
		cfg                           configure.Config
	}
	type args struct {
		ctx context.Context
		in  *proto.EditFieldKeepRequest
	}
	tests := []struct {
		name     string
		user     string
		password string
		fields   fields
		args     args
		want     *proto.EditFieldKeepResponse
		wantErr  bool
	}{
		{
			name:     "Positive test #1",
			user:     "test",
			password: "test",
			fields: fields{
				UnimplementedGophKeeperServer: proto.UnimplementedGophKeeperServer{},
				cfg:                           conf,
			},
			args: args{
				ctx: context.Background(),
				in: &proto.EditFieldKeepRequest{
					Uuid: "6666",
					Data: &proto.FieldKeep{
						Login:    "user",
						Password: "password1",
					},
				},
			},
			want: &proto.EditFieldKeepResponse{
				Data: &proto.FieldKeep{
					Login:    "user",
					Password: "password1",
				},
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

			resp, errLogin := ms.Login(tt.args.ctx, &proto.LoginRequest{Login: tt.user, Password: tt.password})
			if errLogin != nil {
				t.Errorf("Login error = %v", errLogin)
			}

			md := metadata.New(map[string]string{"Authorization": resp.Token})
			ctx := metadata.NewIncomingContext(context.Background(), md)

			got, err := ms.EditField(ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("GophKeeperServer.EditField() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.Data.GetPassword(), tt.want.Data.GetPassword()) {
				t.Errorf("GophKeeperServer.EditField() = %v, want %v", got.Data.GetPassword(), tt.want.Data.GetPassword())
			}
		})
	}
}

func TestGophKeeperServer_DelField(t *testing.T) {
	storeKeeper := &store.StorageContext{}
	mockStore := new(mocks.MockStore)
	mockStore.On("UserLogin", mock.Anything, "test", "test").Return(nil)
	mockStore.On("DelField", mock.Anything, "test", "6666").Return("6666", true)
	storeKeeper.SetStorage(mockStore)

	type fields struct {
		UnimplementedGophKeeperServer proto.UnimplementedGophKeeperServer
		cfg                           configure.Config
	}
	type args struct {
		ctx context.Context
		in  *proto.DeleteFieldKeepRequest
	}
	tests := []struct {
		name     string
		user     string
		password string
		fields   fields
		args     args
		want     *proto.DeleteFieldKeepResponse
		wantErr  bool
	}{
		{
			name:     "Positive test #1",
			user:     "test",
			password: "test",
			fields: fields{
				cfg:                           conf,
				UnimplementedGophKeeperServer: proto.UnimplementedGophKeeperServer{},
			},
			args: args{
				ctx: context.Background(),
				in: &proto.DeleteFieldKeepRequest{
					Uuid: "6666",
				},
			},
			want:    &proto.DeleteFieldKeepResponse{Uuid: "6666"},
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
			resp, errLogin := ms.Login(tt.args.ctx, &proto.LoginRequest{Login: tt.user, Password: tt.password})
			if errLogin != nil {
				t.Errorf("Login error = %v", errLogin)
			}

			md := metadata.New(map[string]string{"Authorization": resp.Token})
			ctx := metadata.NewIncomingContext(context.Background(), md)

			got, err := ms.DelField(ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("GophKeeperServer.DelField() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.Uuid, tt.want.Uuid) {
				t.Errorf("GophKeeperServer.DelField() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGophKeeperServer_ListFields(t *testing.T) {
	storeKeeper := &store.StorageContext{}
	mockStore := new(mocks.MockStore)
	mockStore.On("UserLogin", mock.Anything, "test", "test").Return(nil)
	field := proto.FieldKeep{
		Login:    "user",
		Password: "password",
	}
	resp := &proto.ListFielsdKeepResponse{
		Data: make(map[string]*proto.FieldKeep),
	}

	resp.Data["6666"] = &field
	mockStore.On("ListFields", mock.Anything, "test").Return(resp, true)
	storeKeeper.SetStorage(mockStore)

	type fields struct {
		UnimplementedGophKeeperServer proto.UnimplementedGophKeeperServer
		cfg                           configure.Config
	}
	type args struct {
		ctx context.Context
		in1 *proto.ListFieldsKeepRequest
	}
	tests := []struct {
		name     string
		user     string
		password string
		fields   fields
		args     args
		want     *proto.ListFielsdKeepResponse
		wantErr  bool
	}{
		{
			name:     "Positive test #1",
			user:     "test",
			password: "test",
			fields: fields{
				cfg:                           conf,
				UnimplementedGophKeeperServer: proto.UnimplementedGophKeeperServer{},
			},
			args: args{
				ctx: context.Background(),
				in1: &proto.ListFieldsKeepRequest{},
			},
			want:    resp,
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

			resp, errLogin := ms.Login(tt.args.ctx, &proto.LoginRequest{Login: tt.user, Password: tt.password})
			if errLogin != nil {
				t.Errorf("Login error = %v", errLogin)
			}

			md := metadata.New(map[string]string{"Authorization": resp.Token})
			ctx := metadata.NewIncomingContext(context.Background(), md)

			got, err := ms.ListFields(ctx, tt.args.in1)
			if (err != nil) != tt.wantErr {
				t.Errorf("GophKeeperServer.ListFields() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got.GetData(), tt.want.GetData()) {
				t.Errorf("GophKeeperServer.ListFields() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGophKeeperServer_checkToken(t *testing.T) {
	storeKeeper := &store.StorageContext{}
	mockStore := new(mocks.MockStore)
	mockStore.On("UserLogin", mock.Anything, "test", "test").Return(nil)
	mockStore.On("UserLogin", mock.Anything, "test", "test2").Return(store.ErrAuthentication)
	storeKeeper.SetStorage(mockStore)

	type fields struct {
		UnimplementedGophKeeperServer proto.UnimplementedGophKeeperServer
		cfg                           configure.Config
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name     string
		user     string
		password string
		fields   fields
		args     args
		want     *UserClaims
		wantErr  bool
	}{
		{
			name:     "Positive test #1",
			user:     "test",
			password: "test",
			fields: fields{
				UnimplementedGophKeeperServer: proto.UnimplementedGophKeeperServer{},
				cfg:                           conf,
			},
			args: args{ctx: context.Background()},
			want: &UserClaims{
				Username: "test",
			},
			wantErr: false,
		},
		{
			name:     "Negative test #1",
			user:     "test",
			password: "test2",
			fields: fields{
				UnimplementedGophKeeperServer: proto.UnimplementedGophKeeperServer{},
				cfg:                           conf,
			},
			args:    args{ctx: context.Background()},
			want:    &UserClaims{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ms := &GophKeeperServer{
				UnimplementedGophKeeperServer: tt.fields.UnimplementedGophKeeperServer,
				storage:                       storeKeeper,
				cfg:                           tt.fields.cfg,
			}

			resp, _ := ms.Login(tt.args.ctx, &proto.LoginRequest{Login: tt.user, Password: tt.password})

			md := metadata.New(map[string]string{"Authorization": resp.Token})
			ctx := metadata.NewIncomingContext(context.Background(), md)

			got, err := ms.checkToken(ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GophKeeperServer.checkToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.Username, tt.want.Username) {
				t.Errorf("GophKeeperServer.checkToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func TestGophKeeperServer_Upload(t *testing.T) {
// 	storeKeeper := &store.StorageContext{}
// 	mockStore := new(mocks.MockStore)
// 	mockStore.On("UserLogin", mock.Anything, "test", "test").Return(nil)
// 	storeKeeper.SetStorage(mockStore)

// 	type fields struct {
// 		UnimplementedGophKeeperServer proto.UnimplementedGophKeeperServer
// 		cfg                           configure.Config
// 	}
// 	type args struct {
// 		stream proto.GophKeeper_UploadServer
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		args    args
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			ms := &GophKeeperServer{
// 				UnimplementedGophKeeperServer: tt.fields.UnimplementedGophKeeperServer,
// 				storage:                       storeKeeper,
// 				cfg:                           tt.fields.cfg,
// 			}
// 			if err := ms.Upload(tt.args.stream); (err != nil) != tt.wantErr {
// 				t.Errorf("GophKeeperServer.Upload() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }