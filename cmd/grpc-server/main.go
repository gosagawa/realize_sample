package main

import (
	"flag"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gosagawa/realize_sample/infra/grpc/server"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var port = flag.String("port", "4000", "port to listen")

func init() {
	flag.Parse()
}

func main() {

	c := server.Config{
		Port: *port,
	}

	server.StartServer(c)
}
