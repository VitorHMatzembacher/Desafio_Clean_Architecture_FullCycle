package main

import (
	"context"
	"log"
	"net"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"

	"ordersystem/configs"
	"ordersystem/internal/infra/database"
	"ordersystem/internal/infra/grpc/service"
	pb "ordersystem/internal/infra/grpc/pb"
	"ordersystem/internal/infra/web"
	"ordersystem/internal/usecase"
)

func main() {
	// Carrega variáveis de ambiente
	_ = godotenv.Load()

	// Conecta ao banco
	db, err := configs.LoadDB()
	if err != nil {
		log.Fatalf("erro ao conectar BD: %v", err)
	}
	defer db.Close()

	// Repositório e usecase
	repo := database.NewOrderRepository(db)
	uc := usecase.NewListOrdersUseCase(repo)

	// --- HTTP (Gin) ---
	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}
	router := gin.Default()
	httpHandler := web.NewOrderHandler(uc)
	router.GET("/order", httpHandler.ListOrders)

	// Inicia HTTP em goroutine
	go func() {
		log.Printf("HTTP rodando em :%s", httpPort)
		if err := router.Run(":" + httpPort); err != nil {
			log.Fatalf("failed to start HTTP: %v", err)
		}
	}()

	// --- gRPC ---
	grpcPort := os.Getenv("GRPC_PORT")
	if grpcPort == "" {
		grpcPort = "50051"
	}
	lis, err := net.Listen("tcp", ":"+grpcPort)
	if err != nil {
		log.Fatalf("failed to listen gRPC: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterOrderServiceServer(grpcServer, service.NewOrderServiceServer(uc))

	log.Printf("gRPC rodando em :%s", grpcPort)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to start gRPC: %v", err)
	}
}
