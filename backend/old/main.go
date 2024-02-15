package main

import (
	restFul "backend/Restful"
	"path/filepath"

	"github.com/joho/godotenv"
	"backend/Model"
)


func main() {
	
	if godotenv.Load(filepath.Join("./", ".env")) != nil {
		panic("error loading .env file")
	}
	Model.SetUp()
	restFul.SetUp()
}
