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
	graph "github.com/VitorHMatzembacher/Desafio_Clean_Architecture_FullCycle/internal/infra/graph"
	pb "github.com/VitorHMatzembacher/Desafio_Clean_Architecture_FullCycle/internal/infra/grpc/pb"
	"github.com/VitorHMatzembacher/Desafio_Clean_Architecture_FullCycle/internal/infra/grpc/service"
	"github.com/VitorHMatzembacher/Desafio_Clean_Architecture_FullCycle/internal/infra/web"
	"github.com/VitorHMatzembacher/Desafio_Clean_Architecture_FullCycle/internal/usecase"
)

func main() {
	_ = godotenv.Load()

	// conecta ao banco
	db, err := configs.LoadDB()
	if err != nil {
		log.Fatalf("erro ao conectar BD: %v", err)
	}
	defer db.Close()

	// monta o repositório e os use cases
	repo := database.NewOrderRepository(db)
	createUc := usecase.NewCreateOrderUseCase(repo)
	listUc := usecase.NewListOrdersUseCase(repo)

	// instancia o handler REST com ambos os use cases
	handler := web.NewOrderHandler(createUc, listUc)

	// 🚀 REST
	router := gin.Default()
	router.POST("/order", handler.CreateOrder)
	router.GET("/order", handler.ListOrders)

	// 🚀 GraphQL
	schemaSDL := graph.MustLoadSchema()
	schema := graphql.MustParseSchema(
		schemaSDL,
		&graph.Resolver{ListOrdersUseCase: listUc},
	)
	router.POST("/query", gin.WrapH(&relay.Handler{Schema: schema}))
	router.GET("/playground", gin.WrapH(&relay.Handler{Schema: schema}))

	// HTTP listener
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

	// 🚀 gRPC
	grpcPort := os.Getenv("GRPC_PORT")
	if grpcPort == "" {
		grpcPort = "50051"
	}
	lis, err := net.Listen("tcp", ":"+grpcPort)
	if err != nil {
		log.Fatalf("failed to listen gRPC: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterOrderServiceServer(grpcServer, service.NewOrderServiceServer(listUc))

	log.Printf("gRPC rodando em :%s", grpcPort)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to start gRPC: %v", err)
	}
}
