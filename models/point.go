package models

type Point struct {
	IDShop      string `json:"idShop" bson:"id_shop"`
	IDCustomer  string `json:"idCustomer" bson:"id_customer"`
	Amount      int    `json:"amount" bson:"amount"`
	NumerPoints int    `json:"numerPoints" bson:"number_points"`
	Add         bool   `json:"add" bson:"add"`
	Award       string `json:"award,omitempty" bson:"award,omitempty"`
}

type ResponseHistoryPoints struct {
	Customer Customer `json:"customer"`
	Points   []Point  `json:"poins"`
}

type Resposnse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
