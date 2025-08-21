package grpc

import (
	"net"

	taskpb "github.com/AddonVbs/project-protos/proto/task"
	userpb "github.com/AddonVbs/project-protos/proto/user"

	"github.com/AddonVbs/tasks-service/internal/task"
	"google.golang.org/grpc"
)

func RunGRPC(svc *task.TaskServers, uc userpb.UserServiceClient) error {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		return err
	}

	grpcSrv := grpc.NewServer()
	handler := NewHandler(*svc, uc)
	taskpb.RegisterTaskServiceServer(grpcSrv, handler)

	return grpcSrv.Serve(lis)
}
