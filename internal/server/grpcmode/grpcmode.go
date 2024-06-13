// Package grpcmode запускает сервер.
package grpcmode

import (
	"gophkeeper/internal/server/configure"
	"gophkeeper/internal/store"
	"gophkeeper/internal/store/pg"

	pb "gophkeeper/pkg/proto"
)

// Run запускает сервер
func Run(cfg configure.Config){
	pg.Migrations(cfg.DatabaseDsn)

	// storeKeeper := &store.StorageContext{}
	// storeKeeper.SetStorage(pg.NewDatabase(cfg.DatabaseDsn))

	// listen, err := net.Listen("tcp", cfg.Address)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // создаём gRPC-сервер без зарегистрированной службы
	// s := grpc.NewServer()

	// // регистрируем сервис
	// metricsServer := GophKeeperServer{
	// 	storage: storeKeeper,
	// }
	// pb.RegisterGophKeeperServer(s, &metricsServer)

	// // получаем запрос gRPC
	// if err := s.Serve(listen); err != nil {
	// 	log.Fatal(err)
	// }
}

// GophKeeperServer поддерживает все необходимые методы сервера.
type GophKeeperServer struct {
	pb.UnimplementedGophKeeperServer

	storage *store.StorageContext
}