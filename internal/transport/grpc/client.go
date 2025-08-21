package grpc

import (
	userpb "github.com/AddonVbs/project-protos/proto/user"
	"google.golang.org/grpc"
)

func NewUserClient(addr string) (userpb.UserServiceClient, *grpc.ClientConn, error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure()) // устанавливаем соединение
	if err != nil {
		return nil, nil, err
	}
	client := userpb.NewUserServiceClient(conn)
	return client, conn, nil
}
