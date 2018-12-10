package gateway

import (
	"context"

	"github.com/gosagawa/realize_sample/adapter/grpc/proto"
	"github.com/gosagawa/realize_sample/library/log"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

var serviceHandlers = make(map[string]func(context.Context, *runtime.ServeMux, *grpc.ClientConn) error)

func init() {
	serviceHandlers["User"] = proto.RegisterUserServiceHandler
}

// RegisterServiceHandlers サービスハンドラをまとめて登録する
func RegisterServiceHandlers(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {

	err := proto.RegisterUserServiceHandler(ctx, mux, conn)
	if err != nil {
		return err
	}

	for _, v := range serviceHandlers {
		if err := v(ctx, mux, conn); err != nil {
			return err
		}
		log.Logger.Debug("registered service handler")
	}
	return nil
}
