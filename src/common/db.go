package db

import (
	"fmt"
	"os"
	"time"

	"github.com/Kamva/mgm/v2"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ConnectDb is connect mongodb database
func ConnectDb() error {
	mongo := os.Getenv("MONGO_HOST")
	if mongo == "" {
		mongo = "mongodb://localhost:27017"
	}

	db := os.Getenv("MONGO_DB_NAME")
	if db == "" {
		db = "go-book"
	}

	fmt.Println("mongo: ", mongo)

	//  _ = mgm.SetDefaultConfig(&mgm.Config{CtxTimeout: 12 * time.Second}, "go-book", options.Client().ApplyURI("mongodb://root:12345@localhost:27017"))
	err := mgm.SetDefaultConfig(&mgm.Config{CtxTimeout: 12 * time.Second}, db, options.Client().ApplyURI(mongo))
	if err != nil {
		fmt.Println("Connect database error: ", err)
	}

	return err
}
