package main

import (
	"log"
	"net"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"

	"ordersystem/configs"
	"ordersystem/internal/infra/database"
	pb "ordersystem/internal/infra/grpc/pb"
	"ordersystem/internal/infra/grpc/service"
	"ordersystem/internal/infra/web"
	"ordersystem/internal/usecase"

	graph "ordersystem/internal/infra/graph"
	"ordersystem/internal/infra/graph/generated"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

func main() {
	_ = godotenv.Load()

	db, err := configs.LoadDB()
	if err != nil {
		log.Fatalf(" erro ao conectar BD: %v", err)
	}
	defer db.Close()

	repo := database.NewOrderRepository(db)
	uc := usecase.NewListOrdersUseCase(repo)

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	router := gin.Default()

	httpHandler := web.NewOrderHandler(uc)
	router.GET("/order", httpHandler.ListOrders)

	srv := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{Resolvers: &graph.Resolver{ListOrdersUseCase: uc}},
		),
	)
	router.POST("/query", gin.WrapH(srv))
	router.GET("/playground", gin.WrapH(playground.Handler("GraphQL Playground", "/query")))

	go func() {
		log.Printf(" HTTP rodando em :%s", httpPort)
		if err := router.Run(":" + httpPort); err != nil {
			log.Fatalf(" failed to start HTTP: %v", err)
		}
	}()

	grpcPort := os.Getenv("GRPC_PORT")
	if grpcPort == "" {
		grpcPort = "50051"
	}

	lis, err := net.Listen("tcp", ":"+grpcPort)
	if err != nil {
		log.Fatalf(" failed to listen gRPC: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterOrderServiceServer(grpcServer, service.NewOrderServiceServer(uc))

	log.Printf(" gRPC rodando em :%s", grpcPort)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf(" failed to start gRPC: %v", err)
	}
}
