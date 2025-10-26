package main

import (
	"go-test/backend/bootstrap"
	"go-test/backend/handlers"
	"go-test/backend/repository"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	// Configure CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Register custom validators
	bootstrap.RegisterCustomValidators()

	s := repository.NewStore()
	h := handlers.NewItemsHandler(s)

	r.GET("/items", h.GetAll)
	r.GET("/items/:guid", h.GetByGUID)
	r.POST("/items", h.Create)
	r.PUT("/items/:guid", h.Update)
	r.DELETE("/items/:guid", h.Delete)

	err := r.Run()
	if err != nil {
		return
	}
}
