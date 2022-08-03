package features

import (
	"api-points/models"
)

func (feat Feature) SavePoint(point models.Point) error {
	err := feat.CalculatePoint(&point)
	if err != nil {
		return err
	}

	if point.NumerPoints == 0 {
		return nil
	}

	customer, err := feat.db.FindCustomer(point.IDCustomer)

	if err != nil {
		return err
	}

	customer.CurrentPoints = customer.CurrentPoints + point.NumerPoints

	err = feat.UpdateCurrenPoint(customer)

	if err != nil {
		return err
	}

	point.Add = true

	return feat.db.SavePoint(point)
}

func (feat Feature) CalculatePoint(point *models.Point) error {
	rate, err := feat.db.FindRate(point.IDShop)

	if err != nil {
		return err
	}
	point.NumerPoints = point.Amount / rate.MoneyByPoint

	return nil
}

func (feat Feature) UpdateCurrenPoint(customer models.Customer) error {
	return feat.db.UpdateCustomer(customer)
}
