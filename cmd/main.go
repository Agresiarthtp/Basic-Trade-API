package main

import (
	"Basic-Trade-API/pkg/config"
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	// Call all configuration
	godotenv.Load()

	// Call configuration
	env, err := config.InitialAllConfig()
	if err != nil {
		log.Panic("error configuration", err)
	}

	// Sql connection
	databaseSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", env.MySqlName, env.MySqlPass, env.MySqlHost, env.MySqlPort, env.MySqlDbName)
	dbConnection, err := sql.Open("mysql", databaseSource)

	fmt.Println("env", env)
	if err != nil {
		log.Panic("error failed mysql", err)
	}
	defer dbConnection.Close()

	// Ping verifies a connection to the database is still alive,
	err = dbConnection.Ping()
	if err != nil {
		log.Panic("not alive!")
	}
}
