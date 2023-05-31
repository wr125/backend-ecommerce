package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/wr125/backend-ecommerce/controllers"
	"github.com/wr125/backend-ecommerce/database"
	"github.com/wr125/backend-ecommerce/middleware"
	"github.com/wr125/backend-ecommerce/routes"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("could not load the .env file")
		os.Exit(1)
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	secert := os.Getenv("SECRET_KEY")
	if secert == "" {
		fmt.Printf("could not read the sercert key %v from .env file", secert)
	}
	app := controllers.NewApplication(database.ProductData(database.Client, "Products"), database.UserData(database.Client, "Users"))

	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)
	router.Use(middleware.Authentication())
	router.GET("/addtocart", app.AddToCart())
	router.GET("/removeitem", app.RemoveItem())
	router.GET("/listcart", controllers.GetItemFromCart())
	router.POST("/addaddress", controllers.AddAddress())
	router.PUT("/edithomeaddress", controllers.EditHomeAddress())
	router.PUT("/editworkaddress", controllers.EditWorkAddress())
	router.GET("/deleteaddresses", controllers.DeleteAddress())
	router.GET("/cartcheckout", app.BuyFromCart())
	router.GET("/instantbuy", app.InstantBuy())
	log.Fatal(router.Run(":" + port))

}
