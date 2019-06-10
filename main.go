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
)

var db *gorm.DB

type server struct {
	address  string
	port     string
	hostname string
}

type database struct {
	engine  string
	address string
	port    string
	user    string
	pass    string
	name    string
	ssl     string
}

var config struct {
	server   server
	database database
}

func main() {
	loadConfig()
	configureDatabase()

	r := gin.Default()

	r.POST("/", createPost)
	r.GET("/", fetchAllPost)
	r.GET("/:id", fetchSinglePost)
	log.Printf("Hostname %v", config.server.hostname)

	r.Run(config.server.address + ":" + config.server.port)
}

func loadConfig() {

	flag.StringVar(&config.server.address, "address", "0.0.0.0", "Address to listen on")
	flag.StringVar(&config.server.port, "port", "8080", "Port to listen on")
	flag.StringVar(&config.server.hostname, "hostname", "0.0.0.0", "Hostname of Server")

	flag.StringVar(&config.database.engine, "db_engine", "postgres", "postgres or mysql")
	flag.StringVar(&config.database.address, "db_addr", "127.0.0.1", "Address to listen on")
	flag.StringVar(&config.database.port, "db_port", "5432", "Port to listen on")
	flag.StringVar(&config.database.user, "db_user", "postgres", "Hostname of Server")
	flag.StringVar(&config.database.pass, "db_pass", "postgres", "Address to listen on")
	flag.StringVar(&config.database.name, "db_name", "postgres", "Port to listen on")
	flag.StringVar(&config.database.ssl, "db_ssl", "disable", "Hostname of Server")
	flag.Parse()
}

func configureDatabase() {
	//open a db connection
	var err error

	if config.database.engine == "postgres" {
		connection := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			config.database.address, config.database.port, config.database.user,
			config.database.pass, config.database.name, config.database.ssl)
		db, err = gorm.Open("postgres", connection)
	}
	if config.database.engine == "mysql" {
		connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True",
			config.database.user, config.database.pass, config.database.address,
			config.database.port, config.database.name)
		db, err = gorm.Open("mysql", connection)
	}
	if err != nil {
		fmt.Print(err)
		panic("failed to connect database")
	}
	//Migrate the schema
	db.AutoMigrate(&models.Post{})
}
