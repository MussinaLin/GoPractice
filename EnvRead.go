package main

import (
	_ "github.com/joho/godotenv/autoload"
	"log"
	"os"
)

func main(){
	//err := godotenv.Load()
	//if err != nil {
	//	log.Fatal("Error loading .env file")
	//}

	env := os.Getenv("env")
	log.Println("env:", env)
}
