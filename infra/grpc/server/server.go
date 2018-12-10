package server

import (
	"net"

	"github.com/gosagawa/realize_sample/adapter/grpc"
	"github.com/gosagawa/realize_sample/library/log"
	"go.uber.org/zap"
	ggrpc "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Config サーバ起動時の設定情報サーバ起動時の設定情報
type Config struct {
	Port string
}

// StartServer gRPCサーバを起動する
func StartServer(c Config) {

	lis, err := net.Listen("tcp", ":"+c.Port)
	if err != nil {
		log.Logger.Panic("failed to listen")
	}
	s := ggrpc.NewServer()

	if err := grpc.RegisterServices(s); err != nil {
		log.Logger.Panic("failed to register services")
	}

	// Register reflection service on gRPC server.
	reflection.Register(s)

	log.Logger.Info("start grpc server", zap.String("port", c.Port))
	if err := s.Serve(lis); err != nil {
		log.Logger.Panic("failed to serve: %v")
	}
}
