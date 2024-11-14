package config

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/ichtrojan/thoth"
	_ "github.com/joho/godotenv/autoload"
	"log"
	"os"
)

func Database() *sql.DB {
	logger, _ := thoth.Init("log")

	user, exist := os.LookupEnv("MYSQL_USER")

	if !exist {
		logger.Log(errors.New("MYSQL_USER not set in .env"))
		log.Fatal("MYSQL_USER not set in .env")
	}

	pass, exist := os.LookupEnv("MYSQL_PASSWORD")

	if !exist {
		logger.Log(errors.New("MYSQL_PASSWORD not set in .env"))
		log.Fatal("MYSQL_PASSWORD not set in .env")
	}

	host, exist := os.LookupEnv("MYSQL_HOST")

	if !exist {
		logger.Log(errors.New("MYSQL_HOST not set in .env"))
		log.Fatal("MYSQL_HOST not set in .env")
	}

	port, exist := os.LookupEnv("MYSQL_PORT")

	if !exist {
		logger.Log(errors.New("MYSQL_PORT not set in .env"))
		log.Fatal("MYSQL_PORT not set in .env")
	}

	credentials := fmt.Sprintf("%s:%s@(%s:%s)/?charset=utf8&parseTime=True", user, pass, host, port)

	database, err := sql.Open("mysql", credentials)

	if err != nil {
		logger.Log(err)
		log.Fatal(err)
	} else {
		fmt.Println("Database Connection Successful")
	}

	_, err = database.Exec(`CREATE DATABASE gotodo`)

	if err != nil {
		fmt.Println(err)
	}

	_, err = database.Exec(`USE gotodo`)

	if err != nil {
		fmt.Println(err)
	}

	_, err = database.Exec(`
		CREATE TABLE todos (
		    id INT AUTO_INCREMENT,
		    item TEXT NOT NULL,
		    completed BOOLEAN DEFAULT FALSE,
		    PRIMARY KEY (id)
		);
	`)

	if err != nil {
		fmt.Println(err)
	}

	return database
}
