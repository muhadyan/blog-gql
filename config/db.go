package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *gorm.DB
var err error

func InitDB() (*gorm.DB, error) {
	configuration := GetConfig()
	connect_string := fmt.Sprintf("postgres://%s:%s@%s/%s", configuration.DB_USERNAME, configuration.DB_PASSWORD, configuration.DB_HOST, configuration.DB_NAME)
	db, err = gorm.Open("postgres", connect_string)
	if err != nil {
		return nil, err
	}
	// db.AutoMigrate(
	// 	&model.User{},
	// 	&model.Article{},
	// 	&model.Comment{},
	// 	&model.ChildComment{},
	// 	&model.Like{},
	// )

	return db, nil
}

func DbManager() *gorm.DB {
	return db
}

func ConnectMongo() *mongo.Client {
	configuration := GetConfig()
	client, err := mongo.NewClient(options.Client().ApplyURI(configuration.MONGOURI))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	//ping the database
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")
	return client
}

// Client instance
var MongoDB *mongo.Client = ConnectMongo()

// getting database collections
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("blogGQL").Collection(collectionName)
	return collection
}
