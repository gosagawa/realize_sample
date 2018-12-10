package grpc

import (
	"github.com/andfactory/go-webapp-sample/library/log"

	"github.com/andfactory/go-webapp-sample/registry"
	"google.golang.org/grpc"
)

// grpcRequestHandler GRPCのリクエスト処理を受けつけるハンドラクラス
type grpcRequestHandler interface {
	Register(s *grpc.Server, r registry.Repository)
}

// サーバに登録するRequestHandlerのmap
var registers = make(map[string]func(*grpc.Server, registry.Repository) error)

// RegisterServices サーバにgRPCサービスをまとめて登録する
func RegisterServices(s *grpc.Server) error {

	repo := registry.NewRepository()

	for _, v := range registers {
		if err := v(s, repo); err != nil {
			return err
		}
		log.Logger.Debug("registered grpc handler")
	}

	return nil
}
