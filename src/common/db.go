package db

import (
	"fmt"
	"time"

	"github.com/Kamva/mgm/v2"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ConnectDb is connect mongodb database
func ConnectDb() error {
	//  _ = mgm.SetDefaultConfig(&mgm.Config{CtxTimeout: 12 * time.Second}, "go-book", options.Client().ApplyURI("mongodb://root:12345@localhost:27017"))
	err := mgm.SetDefaultConfig(&mgm.Config{CtxTimeout: 12 * time.Second}, "go-book", options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		fmt.Println("Connect database error: ", err)
	}

	return err
}
