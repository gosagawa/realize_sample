package gateway

import (
	"context"
	"net/http"

	"github.com/andfactory/uraraca-webapp/domain/errors"
	"github.com/gosagawa/realize_sample/adapter/grpc/gateway"
	"github.com/gosagawa/realize_sample/library/log"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// Config サーバ起動時の設定情報
type Config struct {
	Port     string
	Endpoint string
}

// StartServer realize_sample-gatewayのサーバを起動する
func StartServer(c Config) {

	opts := []runtime.ServeMuxOption{
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{OrigName: true, EmitDefaults: true}),
	}

	if err := run(c, opts...); err != nil {
		log.Logger.Panic("run realize_sample-gateway failed", zap.Error(errors.WithStack(err)))
	}
}

func run(c Config, opts ...runtime.ServeMuxOption) error {

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	gw, err := newGateway(ctx, c.Endpoint, opts...)
	if err != nil {
		return err
	}

	handler := gw

	log.Logger.Info("start grapc-gateway", zap.String("port", c.Port), zap.String("endpoint", c.Endpoint))
	return http.ListenAndServe(":"+c.Port, handler)
}

func newGateway(ctx context.Context, endpoint string, opts ...runtime.ServeMuxOption) (http.Handler, error) {

	mux := runtime.NewServeMux(opts...)
	dialOpts := []grpc.DialOption{grpc.WithInsecure()}

	conn, err := grpc.Dial(endpoint, dialOpts...)
	if err != nil {
		return nil, err
	}

	err = gateway.RegisterServiceHandlers(ctx, mux, conn)
	if err != nil {
		return nil, err
	}

	return mux, nil
}
