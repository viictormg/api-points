package models

type Rate struct {
	ID           string `json:"id" bson:"_id"`
	IDShop       string `json:"idShop" bson:"id_shop"`
	MoneyByPoint int    `json:"moneyByPoint" bson:"money_by_point"`
}
