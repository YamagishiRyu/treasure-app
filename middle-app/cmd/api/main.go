package main

import (
	"log"
	"os"

	server "github.com/YamagishiRyu/treasure-app/middle-app"
	"github.com/joho/godotenv"
)

func main() {
	log.SetFlags(log.Ldate + log.Ltime + log.Lshortfile)
	log.SetOutput(os.Stdout)

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error loading .env file. %s", err)
	}

	datasource := os.Getenv("DATABASE_DATASOURCE")
	if datasource == "" {
		log.Fatal("Cannot get datasource for database.")
	}

	s := server.NewServer()
	s.Init(datasource)
	s.Run(os.Getenv("PORT"))
}
