package features

import "api-points/models"

func (feat Feature) GetAwards() ([]models.Award, error) {
	return feat.db.FindAwards()
}
func (feat Feature) GetAward(id string) (models.Award, error) {
	return feat.db.FindAward(id)
}
