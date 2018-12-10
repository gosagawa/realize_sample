package main

import (
	"flag"

	"github.com/gosagawa/realize_sample/infra/grpc/gateway"
)

var port = flag.String("port", "3000", "port to listen")
var endpoint = flag.String("endpoint", "localhost:4000", "grpc server endpoint")

func init() {
	flag.Parse()
}

func main() {

	c := gateway.Config{
		Port:     *port,
		Endpoint: *endpoint,
	}

	gateway.StartServer(c)
}
