package repository

import (
	"api-points/models"
	"context"
)

func (r Repository) SavePoint(point models.Point) error {
	col := r.db.Collection("history_points")

	_, err := col.InsertOne(context.Background(), point)

	return err

}
