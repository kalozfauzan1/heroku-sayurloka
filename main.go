package main

import (
	"./config"
	"./controllers"
	"./controllers/home"
	"./controllers/list_product"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

func main() {
	db := config.DBInit()
	inDB := &controllers.InDB{DB: db}
	home := home.InDB{DB: db}
	list_product := list_product.InDB{DB: db}
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	router := gin.Default()

	router.POST("/api/authorization", inDB.GenerateAuth)
	//promotion
	router.GET("/api/get_promotion", inDB.GetPromotion)

	//categories
	router.GET("/api/get_categories", inDB.GetCategorie)

	//order
	router.GET("/api/my_order/:customer_id", inDB.ListOrder)

	//product
	router.GET("/api/list_product", inDB.ListProduct)

	// page home
	router.GET("/api/list_home", home.ListHome)

	// page list_product
	router.GET("/api/list_product/:type", list_product.ListProduct)

	router.Run(":" + port)
}
