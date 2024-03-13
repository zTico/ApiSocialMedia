package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	ConectDb = ""
	Port     = 0
)

func Loading() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	Port, err = strconv.Atoi(os.Getenv("API_PORT"))

	if err != nil {
		Port = 9000
	}

	ConectDb = fmt.Sprintf(
		"%s:%s@%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USUARIO"),
		os.Getenv("DB_SENHA"),
		os.Getenv("DB_NAME"),
	)
}
