package main

import (
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"log"

	_ "tchh.lucpham/docs"
	"tchh.lucpham/pkg/db"
	"tchh.lucpham/pkg/http"
)

// @title	Kul API
// @description Document API server
// @version 1.0
// @host localhost:8000
// @basepath /
// @contact.name   kul
// @contact.email  tchh.lucpham@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

// @schemes http https
func main() {

	err := db.ConnectMongoDb()
	defer db.DisconnectMongoDb()
	if err != nil {
		log.Fatal(err.Error())
	}
	router := http.SetupRouter()

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":8000")
}
