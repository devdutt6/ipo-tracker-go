package mongoutil

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	clientConnection *mongo.Client
	Database         = GetConnection().Database("ipo")
)

func GetConnection() *mongo.Client {
	if clientConnection != nil {
		return clientConnection
	}

	godotenv.Load(".env")

	var (
		ctx = context.Background()
		err error
	)

	if clientConnection, err = mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("DB_URI"))); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Connection to DB Established")
	}

	return clientConnection
}
