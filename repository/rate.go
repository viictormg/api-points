package repository

import (
	"api-points/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r Repository) FindRate(id string) (models.Rate, error) {
	opt := options.Find()
	filter := bson.M{"id_shop": id}

	rows, err := r.db.Collection("rates").Find(context.TODO(), filter, opt)
	if err != nil {
		return models.Rate{}, err
	}
	var pickup models.Rate
	for rows.Next(context.TODO()) {
		err = rows.Decode(&pickup)
		if err != nil {
			return models.Rate{}, err
		}
	}

	return pickup, nil
}
