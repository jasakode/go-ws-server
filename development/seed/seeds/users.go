package seeds

import (
	"go-ws-server/src/models"

	"gorm.io/gorm"
)

func Users(db *gorm.DB) error {
	user := []*models.Users{
		{
			Name:      "Anton",
			CreatedBy: 1,
		},
		{
			Name:      "Ilham",
			CreatedBy: 1,
		},
		{
			Name:      "Albert",
			CreatedBy: 1,
		},
	}
	rs := models.DB.Create(&user)

	if rs.Error != nil {
		return rs.Error
	}
	return nil
}
