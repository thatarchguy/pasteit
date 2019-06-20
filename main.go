package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/thatarchguy/pasteit/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/xo/dburl"
)

var db *gorm.DB

var config struct {
	address  string
	port     string
	hostname string
	database string
}

func main() {
	loadConfig()
	configureDatabase()

	r := gin.Default()

	r.POST("/", createPost)
	r.GET("/", fetchAllPost)
	r.GET("/:id", fetchSinglePost)
	log.Printf("Hostname %v", config.hostname)

	r.Run(config.address + ":" + config.port)
}

func loadConfig() {

	flag.StringVar(&config.address, "address", "0.0.0.0", "Address to listen on")
	flag.StringVar(&config.port, "port", "8080", "Port to listen on")
	flag.StringVar(&config.hostname, "hostname", "0.0.0.0", "Hostname of Server")

	flag.StringVar(&config.database, "database", "sqlite://:memory:", "Database String")
	flag.Parse()
}

func configureDatabase() {
	//open a db connection
	u, err := dburl.Parse(config.database)
	if err != nil {
		fmt.Print(err)
		panic("Invalid Database Schema")
	}

	//GORM handles sqlite in a special way
	if u.Scheme == "sqlite3" || u.Scheme == "sqlite" {
		db, err = gorm.Open("sqlite3", u.Opaque)
	} else {
		db, err = gorm.Open(u.Scheme, u.String())
	}
	if err != nil {
		fmt.Print(err)
		panic("failed to connect database")
	}
	//Migrate the schema
	db.AutoMigrate(&models.Post{})
}
