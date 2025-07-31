package main

import (
	"log"
	"net"
	"os"

	"github.com/gin-gonic/gin"
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"

	"github.com/VitorHMatzembacher/Desafio_Clean_Architecture_FullCycle/configs"
	"github.com/VitorHMatzembacher/Desafio_Clean_Architecture_FullCycle/internal/infra/database"
	pb "github.com/VitorHMatzembacher/Desafio_Clean_Architecture_FullCycle/internal/infra/grpc/pb"
	"github.com/VitorHMatzembacher/Desafio_Clean_Architecture_FullCycle/internal/infra/grpc/service"
	"github.com/VitorHMatzembacher/Desafio_Clean_Architecture_FullCycle/internal/infra/web"
	"github.com/VitorHMatzembacher/Desafio_Clean_Architecture_FullCycle/internal/usecase"

	graph "github.com/VitorHMatzembacher/Desafio_Clean_Architecture_FullCycle/internal/infra/graph"
)

func main() {
	_ = godotenv.Load()

	db, err := configs.LoadDB()
	if err != nil {
		log.Fatalf("erro ao conectar BD: %v", err)
	}
	defer db.Close()

	repo := database.NewOrderRepository(db)
	uc := usecase.NewListOrdersUseCase(repo)

	router := gin.Default()
	router.GET("/order", web.NewOrderHandler(uc).ListOrders)

	schemaSDL := graph.MustLoadSchema()

	schema := graphql.MustParseSchema(
		schemaSDL,
		&graph.Resolver{ListOrdersUseCase: uc},
	)

	router.POST("/query", gin.WrapH(&relay.Handler{Schema: schema}))

	router.GET("/playground", gin.WrapH(&relay.Handler{Schema: schema}))

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}
	go func() {
		log.Printf("HTTP rodando em :%s", httpPort)
		if err := router.Run(":" + httpPort); err != nil {
			log.Fatalf("failed to start HTTP: %v", err)
		}
	}()

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
