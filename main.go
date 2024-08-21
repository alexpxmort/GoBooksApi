package main

import (
	"gobooks/internal/config"
	"gobooks/internal/repository"
	"gobooks/internal/service"
	"gobooks/internal/web"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Inicializa a configuração do banco de dados
	config.InitDB()

	// Cria instâncias do repositório e do serviço
	bookRepo := &repository.BookRepository{}
	bookService := service.NewBookService(bookRepo)
	bookHandlers := web.NewBookHandlers(bookService)

	// Inicializa o roteador HTTP e define as rotas
	router := gin.Default()

	booksGroup := router.Group("/books")
	{
		booksGroup.POST("", bookHandlers.CreateBook)                    // POST /books
		booksGroup.GET("", bookHandlers.GetBooks)                       // GET /books
		booksGroup.GET("/:id", bookHandlers.GetBookById)                // GET /books/:id
		booksGroup.PUT("/:id", bookHandlers.UpdateBook)                 // PUT /books/:id
		booksGroup.PATCH("markAsRead/:id", bookHandlers.MarkBookAsRead) // PATCH /markAsRead/:id
	}

	// Inicia o servidor HTTP
	log.Println("Iniciando servidor na porta 8080...")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
