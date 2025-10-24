package main

import (
	"go-test/backend/handlers"
	"go-test/backend/models/validators"
	"go-test/backend/store"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
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

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err := v.RegisterValidation("itemtype", validators.ValidateItemType)
		if err != nil {
			log.Fatal("Failed to register itemtype validator:", err)
		}
		err = v.RegisterValidation("itemstatus", validators.ValidateItemStatus)
		if err != nil {
			log.Fatal("Failed to register itemstatus validator:", err)
		}
		err = v.RegisterValidation("sortcode", validators.ValidateSortCode)
		if err != nil {
			log.Fatal("Failed to register sortcode validator:", err)
		}
	}

	s := store.NewStore()
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
