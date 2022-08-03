package repository

import (
	"api-points/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r Repository) FindCustomer(id string) (models.Customer, error) {
	opt := options.Find()

	filter := bson.M{"id_customer": id}

	rows, err := r.db.Collection("customers").Find(context.TODO(), filter, opt)
	if err != nil {
		return models.Customer{}, err
	}
	var customer models.Customer
	for rows.Next(context.TODO()) {
		err = rows.Decode(&customer)
		if err != nil {
			return models.Customer{}, err
		}
	}

	return customer, nil
}

func (r Repository) UpdateCustomer(customer models.Customer) error {

	filter := bson.M{"id_customer": customer.IDCustomer}

	update := bson.M{"$set": bson.M{
		"current_points": customer.CurrentPoints,
	}}

	_, err := r.db.Collection("customers").UpdateOne(
		context.Background(),
		filter,
		update,
	)

	return err
}
