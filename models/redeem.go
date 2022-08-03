package models

type Redeem struct {
	IDCustomer string `json:"idCustomer" bson:"id_customer"`
	IDAward    string `json:"idAward" bson:"id_award"`
}
