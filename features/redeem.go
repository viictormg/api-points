package features

import (
	"api-points/models"
	"fmt"
)

func (feat Feature) RedeemPoints(redeem models.Redeem) error {
	award, err := feat.GetAward(redeem.IDAward)
	if err != nil {
		return err
	}

	customer, err := feat.db.FindCustomer(redeem.IDCustomer)
	if err != nil {
		return err
	}

	if award.Points <= customer.CurrentPoints {
		customer.CurrentPoints = customer.CurrentPoints - award.Points
		err = feat.UpdateCurrenPoint(customer)
		if err != nil {
			return err
		}

		point := models.Point{
			IDShop:      "",
			IDCustomer:  redeem.IDCustomer,
			NumerPoints: award.Points,
			Add:         false,
		}

		return feat.db.SavePoint(point)
	}

	return fmt.Errorf("no tiene suficientes puntos")

}
