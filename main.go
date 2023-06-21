package main

import (
	"os"

	"github.com/atifali-pm/go-resturant-management/middleware"
	"github.com/atifali-pm/go-resturant-management/routes"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := gin.New()

	router.Use(gin.Logger())

	routes.UserRouter(router)
	router.Use(middleware.Authentication)

	router.Run(":" + port)

}
