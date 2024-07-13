// Package pg реализует взаимодействие с базой данной Postgres.
package pg

import (
	"context"
	"fmt"
	"gophkeeper/internal/server/configure"
	"gophkeeper/pkg/proto"
	"reflect"
	"testing"
	"time"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestDatabase_ListFields(t *testing.T) {
	const (
		usr      = "usr"
		password = "pass"
		dbName   = "gophkeeper"
	)

	ctx := context.Background()

	postgresContainer, err := postgres.Run(ctx,
		"docker.io/postgres:16-alpine",
		postgres.WithDatabase(dbName),
		postgres.WithUsername(usr),
		postgres.WithPassword(password),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(10*time.Second)),
	)
	if err != nil {
		t.Errorf("failed to start container: %s", err)
	}

	defer func() {
		if err := postgresContainer.Terminate(ctx); err != nil {
			t.Errorf("failed to terminate container: %s", err)
		}
	}()

	host, _ := postgresContainer.Host(context.Background())
	port, _ := postgresContainer.MappedPort(context.Background(), "5432")

	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		usr,
		password,
		host,
		port.Port(),
		dbName,
	)

	err = Migrations(configure.Config{DatabaseDsn: dsn})
	if err != nil {
		t.Error(err)
	}
	db := NewDatabase(dsn)

	err = db.UserRegister(context.Background(), "test", "test")
	if err != nil {
		t.Error(err)
	}

	db.AddField(context.Background(), "test", &proto.FieldKeep{
		Login:    "test",
		Password: "test2",
	})
	db.AddField(context.Background(), "test", &proto.FieldKeep{
		Login:    "test3",
		Password: "test2",
	})

	type args struct {
		ctx  context.Context
		user string
	}
	tests := []struct {
		name   string
		args   args
		wantOk bool
	}{
		{
			name: "Positive test #1",
			args: args{
				ctx:  context.Background(),
				user: "test",
			},
			wantOk: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, gotOk := db.ListFields(tt.args.ctx, tt.args.user)

			if gotOk != tt.wantOk {
				t.Errorf("Database.ListFields() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func TestDatabase_UserRegister(t *testing.T) {
	const (
		usr      = "usr"
		password = "pass"
		dbName   = "gophkeeper"
	)

	ctx := context.Background()

	postgresContainer, err := postgres.Run(ctx,
		"docker.io/postgres:16-alpine",
		postgres.WithDatabase(dbName),
		postgres.WithUsername(usr),
		postgres.WithPassword(password),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(10*time.Second)),
	)
	if err != nil {
		t.Errorf("failed to start container: %s", err)
	}

	defer func() {
		if err := postgresContainer.Terminate(ctx); err != nil {
			t.Errorf("failed to terminate container: %s", err)
		}
	}()

	host, _ := postgresContainer.Host(context.Background())
	port, _ := postgresContainer.MappedPort(context.Background(), "5432")

	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		usr,
		password,
		host,
		port.Port(),
		dbName,
	)

	err = Migrations(configure.Config{DatabaseDsn: dsn})
	if err != nil {
		t.Error(err)
	}
	db := NewDatabase(dsn)

	type args struct {
		ctx      context.Context
		login    string
		password string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Positive test #1",
			args: args{
				ctx:      context.Background(),
				login:    "test",
				password: "test",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := db.UserRegister(tt.args.ctx, tt.args.login, tt.args.password); (err != nil) != tt.wantErr {
				t.Errorf("Database.UserRegister() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDatabase_UserLogin(t *testing.T) {
	const (
		usr      = "usr"
		password = "pass"
		dbName   = "gophkeeper"
	)

	ctx := context.Background()

	postgresContainer, err := postgres.Run(ctx,
		"docker.io/postgres:16-alpine",
		postgres.WithDatabase(dbName),
		postgres.WithUsername(usr),
		postgres.WithPassword(password),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(10*time.Second)),
	)
	if err != nil {
		t.Errorf("failed to start container: %s", err)
	}

	defer func() {
		if err := postgresContainer.Terminate(ctx); err != nil {
			t.Errorf("failed to terminate container: %s", err)
		}
	}()

	host, _ := postgresContainer.Host(context.Background())
	port, _ := postgresContainer.MappedPort(context.Background(), "5432")

	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		usr,
		password,
		host,
		port.Port(),
		dbName,
	)

	err = Migrations(configure.Config{DatabaseDsn: dsn})
	if err != nil {
		t.Error(err)
	}
	db := NewDatabase(dsn)

	err = db.UserRegister(context.Background(), "test", "test")
	if err != nil {
		t.Error(err)
	}

	type args struct {
		ctx      context.Context
		login    string
		password string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Positive test #1",
			args: args{
				ctx:      context.Background(),
				login:    "test",
				password: "test",
			},
			wantErr: false,
		},
		{
			name: "Negative test #1",
			args: args{
				ctx:      context.Background(),
				login:    "test",
				password: "test2",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := db.UserLogin(tt.args.ctx, tt.args.login, tt.args.password); (err != nil) != tt.wantErr {
				t.Errorf("Database.UserLogin() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDatabase_AddField(t *testing.T) {
	const (
		usr      = "usr"
		password = "pass"
		dbName   = "gophkeeper"
	)

	ctx := context.Background()

	postgresContainer, err := postgres.Run(ctx,
		"docker.io/postgres:16-alpine",
		postgres.WithDatabase(dbName),
		postgres.WithUsername(usr),
		postgres.WithPassword(password),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(10*time.Second)),
	)
	if err != nil {
		t.Errorf("failed to start container: %s", err)
	}

	defer func() {
		if err := postgresContainer.Terminate(ctx); err != nil {
			t.Errorf("failed to terminate container: %s", err)
		}
	}()

	host, _ := postgresContainer.Host(context.Background())
	port, _ := postgresContainer.MappedPort(context.Background(), "5432")

	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		usr,
		password,
		host,
		port.Port(),
		dbName,
	)

	err = Migrations(configure.Config{DatabaseDsn: dsn})
	if err != nil {
		t.Error(err)
	}
	db := NewDatabase(dsn)

	err = db.UserRegister(context.Background(), "test", "test")
	if err != nil {
		t.Error(err)
	}

	type args struct {
		ctx  context.Context
		user string
		data *proto.FieldKeep
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 *proto.FieldKeep
		want2 bool
	}{
		{
			name: "Positive test #1",
			args: args{
				ctx:  context.Background(),
				user: "test",
				data: &proto.FieldKeep{
					Login:    "test",
					Password: "test",
				},
			},
			want: "66323904-cbbc-46cf-bc3b-bd84362c426a",
			want1: &proto.FieldKeep{
				Login:    "test",
				Password: "test",
			},
			want2: true,
		},
		{
			name: "Negative test #1",
			args: args{
				ctx:  context.Background(),
				user: "test2",
				data: &proto.FieldKeep{
					Login:    "test",
					Password: "test",
				},
			},
			want:  "",
			want1: nil,
			want2: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := db.AddField(tt.args.ctx, tt.args.user, tt.args.data)
			if got2 != tt.want2 {
				t.Errorf("Database.AddField() got2 = %v, want %v", got2, tt.want2)
			}
			if len([]rune(got)) != len([]rune(tt.want)) {
				t.Errorf("Database.AddField() got = %v", got)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Database.AddField() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestDatabase_EditField(t *testing.T) {
	const (
		usr      = "usr"
		password = "pass"
		dbName   = "gophkeeper"
	)

	ctx := context.Background()

	postgresContainer, err := postgres.Run(ctx,
		"docker.io/postgres:16-alpine",
		postgres.WithDatabase(dbName),
		postgres.WithUsername(usr),
		postgres.WithPassword(password),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(10*time.Second)),
	)
	if err != nil {
		t.Errorf("failed to start container: %s", err)
	}

	defer func() {
		if err := postgresContainer.Terminate(ctx); err != nil {
			t.Errorf("failed to terminate container: %s", err)
		}
	}()

	host, _ := postgresContainer.Host(context.Background())
	port, _ := postgresContainer.MappedPort(context.Background(), "5432")

	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		usr,
		password,
		host,
		port.Port(),
		dbName,
	)

	err = Migrations(configure.Config{DatabaseDsn: dsn})
	if err != nil {
		t.Error(err)
	}
	db := NewDatabase(dsn)

	err = db.UserRegister(context.Background(), "test", "test")
	if err != nil {
		t.Error(err)
	}

	uuid1, _, _ := db.AddField(context.Background(), "test", &proto.FieldKeep{
		Login:    "test",
		Password: "test2",
	})
	uuid2, _, _ := db.AddField(context.Background(), "test", &proto.FieldKeep{
		Login:    "test3",
		Password: "test2",
	})

	type args struct {
		ctx  context.Context
		user string
		uuid string
		data *proto.FieldKeep
	}
	tests := []struct {
		name  string
		args  args
		want  *proto.FieldKeep
		want1 bool
	}{
		{
			name: "Positive test #1",
			args: args{
				ctx:  context.Background(),
				user: "test",
				uuid: uuid1,
				data: &proto.FieldKeep{
					Login:    "test2",
					Password: "test2",
				},
			},
			want: &proto.FieldKeep{
				Login:    "test2",
				Password: "test2",
			},
			want1: true,
		},
		{
			name: "Negative test #1",
			args: args{
				ctx:  context.Background(),
				user: "test2",
				uuid: uuid2,
				data: &proto.FieldKeep{
					Login:    "test3",
					Password: "test3",
				},
			},
			want:  nil,
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := db.EditField(tt.args.ctx, tt.args.user, tt.args.uuid, tt.args.data)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Database.EditField() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Database.EditField() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestDatabase_getUserID(t *testing.T) {
	const (
		usr      = "usr"
		password = "pass"
		dbName   = "gophkeeper"
	)

	ctx := context.Background()

	postgresContainer, err := postgres.Run(ctx,
		"docker.io/postgres:16-alpine",
		postgres.WithDatabase(dbName),
		postgres.WithUsername(usr),
		postgres.WithPassword(password),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(10*time.Second)),
	)
	if err != nil {
		t.Errorf("failed to start container: %s", err)
	}

	defer func() {
		if err := postgresContainer.Terminate(ctx); err != nil {
			t.Errorf("failed to terminate container: %s", err)
		}
	}()

	host, _ := postgresContainer.Host(context.Background())
	port, _ := postgresContainer.MappedPort(context.Background(), "5432")

	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		usr,
		password,
		host,
		port.Port(),
		dbName,
	)

	err = Migrations(configure.Config{DatabaseDsn: dsn})
	if err != nil {
		t.Error(err)
	}
	db := NewDatabase(dsn)

	err = db.UserRegister(context.Background(), "test", "test")
	if err != nil {
		t.Error(err)
	}

	type args struct {
		ctx  context.Context
		user string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Positive test #1",
			args: args{
				ctx:  context.Background(),
				user: "test",
			},
			wantErr: false,
		},
		{
			name: "Negative test #1",
			args: args{
				ctx:  context.Background(),
				user: "test2",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := db.getUserID(tt.args.ctx, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("Database.getUserID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestDatabase_DelField(t *testing.T) {
	const (
		usr      = "usr"
		password = "pass"
		dbName   = "gophkeeper"
	)

	ctx := context.Background()

	postgresContainer, err := postgres.Run(ctx,
		"docker.io/postgres:16-alpine",
		postgres.WithDatabase(dbName),
		postgres.WithUsername(usr),
		postgres.WithPassword(password),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(10*time.Second)),
	)
	if err != nil {
		t.Errorf("failed to start container: %s", err)
	}

	defer func() {
		if err := postgresContainer.Terminate(ctx); err != nil {
			t.Errorf("failed to terminate container: %s", err)
		}
	}()

	host, _ := postgresContainer.Host(context.Background())
	port, _ := postgresContainer.MappedPort(context.Background(), "5432")

	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		usr,
		password,
		host,
		port.Port(),
		dbName,
	)

	err = Migrations(configure.Config{DatabaseDsn: dsn})
	if err != nil {
		t.Error(err)
	}
	db := NewDatabase(dsn)

	err = db.UserRegister(context.Background(), "test", "test")
	if err != nil {
		t.Error(err)
	}

	uuid1, _, _ := db.AddField(context.Background(), "test", &proto.FieldKeep{
		Login:    "test",
		Password: "test2",
	})
	uuid2, _, _ := db.AddField(context.Background(), "test", &proto.FieldKeep{
		Login:    "test3",
		Password: "test2",
	})

	type args struct {
		ctx  context.Context
		user string
		uuid string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 bool
	}{
		{
			name: "Positive test #1",
			args: args{
				ctx:  context.Background(),
				user: "test",
				uuid: uuid1,
			},
			want:  uuid1,
			want1: true,
		},
		{
			name: "Negative test #1",
			args: args{
				ctx:  context.Background(),
				user: "test2",
				uuid: uuid2,
			},
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := db.DelField(tt.args.ctx, tt.args.user, tt.args.uuid)
			if got != tt.want {
				t.Errorf("Database.DelField() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Database.DelField() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
