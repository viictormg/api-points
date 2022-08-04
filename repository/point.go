package repository

import (
	"api-points/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r Repository) SavePoint(point models.Point) error {
	col := r.db.Collection("history_points")

	_, err := col.InsertOne(context.Background(), point)

	return err

}

func (r Repository) GetPoints(idCustomer string) ([]models.Point, error) {
	var poinst []models.Point
	opt := options.Find()

	filter := bson.M{"id_customer": idCustomer}

	rows, err := r.db.Collection("history_points").Find(context.TODO(), filter, opt)
	if err != nil {
		return poinst, err
	}
	var point models.Point
	for rows.Next(context.TODO()) {
		err = rows.Decode(&point)

		if err != nil {
			return poinst, err
		}

		poinst = append(poinst, point)

	}

	return poinst, nil
}
