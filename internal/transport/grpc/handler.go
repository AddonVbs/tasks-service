package grpc

import (
	"context"
	"fmt"

	pb "github.com/AddonVbs/project-protos/proto/task"
	userpb "github.com/AddonVbs/project-protos/proto/user"
	"github.com/AddonVbs/tasks-service/internal/task"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Handler struct {
	pb.UnimplementedTaskServiceServer
	svc task.TaskServers
	uc  userpb.UserServiceClient
}

func NewHandler(svc task.TaskServers, uc userpb.UserServiceClient) *Handler {
	return &Handler{svc: svc, uc: uc}
}

// CreateTask gRPC handler
func (h *Handler) CreateTask(ctx context.Context, req *pb.CreateTaskRequest) (*pb.CreateTaskResponse, error) {
	// 1. Проверить пользователя:
	if _, err := h.uc.GetUser(ctx, &userpb.GetUserRequest{Id: uint32(req.UserID)}); err != nil {
		return nil, fmt.Errorf("user %d not found: %w", req.UserID, err)
	}
	// 2. Внутренняя логика:
	t, err := h.svc.CreateTask(&task.Task{
		UserID: int(req.UserID),
		Task:   req.Title,
	})

	if err != nil {
		return nil, err
	}
	// 3. Ответ:
	return &pb.CreateTaskResponse{Task: &pb.Task{Id: uint32(t.ID), UserID: uint32(t.UserID), Title: t.Task}}, nil
}

func (h *Handler) ListTasks(ctx context.Context, _ *emptypb.Empty) (*pb.ListTasksResponse, error) {
	tasks, err := h.svc.GetAllTask()
	if err != nil {
		return nil, err
	}

	var protoTasks []*pb.Task
	for _, t := range tasks {
		protoTasks = append(protoTasks, &pb.Task{
			Id:     uint32(t.ID),
			Title:  t.Task,
			UserID: uint32(t.UserID),
		})
	}

	return &pb.ListTasksResponse{Tasks: protoTasks}, nil
}

func (h *Handler) GetTask(ctx context.Context, req *pb.Task) (*pb.Task, error) {
	t, err := h.svc.GetTaskByID(int(req.Id))
	if err != nil {
		return nil, err
	}

	return &pb.Task{
		Id:     uint32(t.ID),
		Title:  t.Task,
		UserID: uint32(t.UserID),
	}, nil
}

func (h *Handler) GetTasksByUserID(ctx context.Context, req *pb.GetTasksByUserIDRequest) (*pb.GetTasksByUserIDResponse, error) {
	tasks, err := h.svc.GetTasksForUser(int(req.UserId))
	if err != nil {
		return nil, err
	}

	var protoTasks []*pb.Task
	for _, t := range tasks {
		protoTasks = append(protoTasks, &pb.Task{
			Id:     uint32(t.ID),
			Title:  t.Task,
			UserID: uint32(t.UserID),
		})
	}

	return &pb.GetTasksByUserIDResponse{Tasks: protoTasks}, nil
}

func (h *Handler) DeleteTask(ctx context.Context, req *pb.DeleteTaskRequest) (*emptypb.Empty, error) {
	id := int(req.Id)

	if id == 0 {
		return nil, fmt.Errorf("нет указание по айди, какую задачу удалить ?")
	}

	err := h.svc.DeleteTask(id)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (h *Handler) UpdateTask(ctx context.Context, req *pb.UpdateTaskRequest) (*pb.Task, error) {
	t, err := h.svc.UpdataTask(int(req.Id), req.Title)
	if err != nil {
		return nil, err
	}

	return &pb.Task{
		Id:     uint32(t.ID),
		Title:  t.Task,
		UserID: uint32(t.UserID),
	}, nil
}
