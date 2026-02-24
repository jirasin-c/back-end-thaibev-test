package main

import (
	"log"
	"os"
	"thaibev-test/db"
	"thaibev-test/handler"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	dbconn := db.Connection()
	defer dbconn.Close()

	r := gin.Default()
	r.Use(cors.Default())

	h := handler.New(dbconn)

	api := r.Group("/api")
	{
		api.GET("/occupations", h.GetOccupations)
		api.POST("/people", h.CreatePerson)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(r.Run(":" + port))
}
