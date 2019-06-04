package utils

import "os"

//ConfigureHost sets up the listening address
func ConfigureHost() (string, string, string) {
	port := os.Getenv("PORT")
	if port != "" {
		port = ":" + port
	} else {
		port = ":8080"
	}

	address := os.Getenv("ADDRESS")
	if address == "" {
		address = "0.0.0.0"
	}

	hostname := os.Getenv("HOSTNAME")
	if hostname == "" {
		hostname = address + port
	}
	return address, port, hostname
}

// ConfigurePostgres returns connection string
func ConfigurePostgres() string {
	var connection string
	address := os.Getenv("PG_ADDR")
	if address == "" {
		address = "127.0.0.1"
	}
	connection = "host=" + address

	port := os.Getenv("PG_PORT")
	if port == "" {
		port = "5432"
	}
	connection += " port=" + port

	user := os.Getenv("PG_USER")
	if user == "" {
		user = "postgres"
	}
	connection += " user=" + user

	password := os.Getenv("PG_PASS")
	if password == "" {
		password = "postgres"
	}
	connection += " password=" + password

	database := os.Getenv("PG_DBNAME")
	if database == "" {
		database = "postgres"
	}
	connection += " dbname=" + database

	sslmode := os.Getenv("PG_SSL")
	if sslmode == "" {
		sslmode = "disable"
	}
	connection += " sslmode=" + sslmode

	return connection
}
