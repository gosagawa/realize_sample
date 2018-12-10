package grpc

import (
	"context"

	"github.com/andfactory/go-webapp-sample/registry"
	"github.com/gosagawa/realize_sample/adapter/grpc/proto"
	"google.golang.org/grpc"
)

func init() {
	registers["User"] = registerUser
}

// registerUser サーバにユーザ系の処理を登録する
func registerUser(s *grpc.Server, r registry.Repository) error {
	u := newUser(r)
	proto.RegisterUserServiceServer(s, u)
	return nil
}

// User ユーザ系の処理実施用の構造体
type User struct {
	repo registry.Repository
}

// newUser ユーザ系の処理実施用の構造体を生成する
func newUser(repo registry.Repository) *User {
	return &User{repo}
}

// Get ユーザ情報を取得する
func (u *User) Get(ctx context.Context, in *proto.GetUserRequest) (*proto.User, error) {

	user := &proto.User{
		Id:   in.Id,
		Name: "John",
		Age:  18,
	}

	return user, nil
}
