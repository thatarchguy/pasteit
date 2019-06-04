package main

import (
	"fmt"
	"log"

	"github.com/thatarchguy/pasteit/utils"

	"github.com/gin-gonic/gin"
	"github.com/thatarchguy/pasteit/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB
var address, port, hostname string

func init() {
	//open a db connection
	var err error
	db, err = gorm.Open("postgres", utils.ConfigurePostgres())
	if err != nil {
		fmt.Print(err)
		panic("failed to connect database")
	}
	//Migrate the schema
	db.AutoMigrate(&models.Post{})
}

func main() {

	r := gin.Default()

	address, port, hostname = utils.ConfigureHost()

	r.POST("/", createPost)
	r.GET("/", fetchAllPost)
	r.GET("/:id", fetchSinglePost)
	log.Printf("Hostname %v", hostname)

	r.Run(address + port)
}
