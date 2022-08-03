package config

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitDB(cfg Config) (*mongo.Client, error) {

	// credentials := options.Credential{
	// 	AuthMechanism: cfg.MongoMechanism,
	// 	Username:      cfg.MongoUser,
	// 	Password:      cfg.MongoPassword,
	// }

	clientOptions := options.Client().ApplyURI(cfg.MongoURL)

	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return nil, err
	}

	client.Connect(context.TODO())

	if err != nil {
		return nil, err
	}

	return client, err

}
