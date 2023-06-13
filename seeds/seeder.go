package seeds

import (
	"log"
	"os"

	"github.com/valeelim/mahchat/internal/database"
)

func Seed() {
	var (
		dbUser = os.Getenv("DB_USER")
		dbPort = os.Getenv("DB_PORT")
		dbPass = os.Getenv("DB_PASS")
		dbName = os.Getenv("DB_NAME")
		dbHost = os.Getenv("DB_HOST")
	)

	db, err := database.New(dbUser, dbPort, dbPass, dbName, dbHost)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	channelSeed, err := seedChannel(db, 2)
	if err != nil {
		log.Fatalf("something bad happened, %v", err)
		return 
	}
	seedMessage(db, channelSeed, 3)
}