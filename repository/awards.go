package repository

import (
	"api-points/models"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r Repository) FindAwards() ([]models.Award, error) {
	var awards []models.Award

	rows, err := r.db.Collection("awards").Find(context.TODO(), bson.M{})

	if err != nil {
		return awards, err
	}
	for rows.Next(context.TODO()) {
		var award models.Award
		err := rows.Decode(&award)
		if err != nil {
			return awards, err
		}

		awards = append(awards, models.Award{
			ID:     award.ID,
			Name:   award.Name,
			Points: award.Points,
		})
	}

	return awards, nil
}

func (r Repository) FindAward(id string) (models.Award, error) {
	opt := options.Find()
	objID, err := primitive.ObjectIDFromHex(id)

	filter := bson.M{"_id": bson.M{"$eq": objID}}

	rows, err := r.db.Collection("awards").Find(context.TODO(), filter, opt)
	if err != nil {
		return models.Award{}, err
	}
	var customer models.Award
	for rows.Next(context.TODO()) {
		err = rows.Decode(&customer)
		if err != nil {
			return models.Award{}, err
		}
	}

	return customer, nil
}
