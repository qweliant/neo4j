package driver

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/qweliant/neo4j/api/controllers"
)

var server = controllers.Server{}

func init() {
	// loads values from .env into the system
	if err := godotenv.Load()
	err != nil {
		log.Print("sad .env file found")
	}
}

func Run() {

	var err error
	// err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	}

	server.Database(os.Getenv(os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"))

	// This is for testing, when done, do well to comment
	// seed.Load(server.DB)

	apiPort := fmt.Sprintf(":%d", os.Getenv("API_PORT"))


	server.Run(apiPort)
	fmt.Printf("Listening to port %d", apiPort)

}
