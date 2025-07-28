package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"ordersystem/configs"
	"ordersystem/internal/infra/database"
	"ordersystem/internal/infra/web"
	"ordersystem/internal/usecase"
)

func main() {
	// Carregar variáveis do .env
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  .env não encontrado, usando variáveis de ambiente")
	}

	// Conectar ao banco
	db, err := configs.LoadDB()
	if err != nil {
		log.Fatalf("❌ erro ao conectar ao banco: %v", err)
	}
	defer db.Close()

	// Setup do repositório e usecase
	repo := database.NewOrderRepository(db)
	usecase := usecase.NewListOrdersUseCase(repo)
	handler := web.NewOrderHandler(usecase)

	// Iniciar servidor HTTP
	r := gin.Default()
	r.GET("/order", handler.ListOrders)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("🚀 Servidor rodando em http://localhost:%s/order", port)
	r.Run(":" + port)
}
