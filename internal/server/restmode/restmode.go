// Package restmode REST сервер.
package restmode

import (
	"context"
	"log"
	"net/http"

	"gophkeeper/internal/server/configure"

	pb "gophkeeper/pkg/proto"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Run запуск REST сервера.
func Run(cfg configure.Config) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := pb.RegisterGophKeeperHandlerFromEndpoint(ctx, mux, cfg.Address, opts)
	if err != nil {
		panic(err)
	}
	log.Printf("server listening at 8081")
	if err := http.ListenAndServe(":8081", mux); err != nil {
		panic(err)
	}
}
