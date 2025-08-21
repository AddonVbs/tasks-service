package main

import (
	"log"

	"github.com/AddonVbs/tasks-service/internal/database"
	"github.com/AddonVbs/tasks-service/internal/task"
	transportgrpc "github.com/AddonVbs/tasks-service/internal/transport/grpc"
)

func main() {
	// 1. Init DB
	database.InitDB()

	// автоматически создать таблицу (dev)
	if err := database.DB.AutoMigrate(&task.Task{}); err != nil {
		log.Fatalf("failed to automigrate: %v", err)
	}

	// 2. Репозиторий + сервис
	repo := task.NewTaskRepository(database.DB)
	svc := task.NewTaskService(repo)

	// 3. Клиент к users-service
	userClient, conn, err := transportgrpc.NewUserClient("localhost:50051")
	if err != nil {
		log.Fatalf("failed to connect to users: %v", err)
	}
	defer conn.Close()

	// 4. Запуск tasks gRPC
	if err := transportgrpc.RunGRPC(&svc, userClient); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
