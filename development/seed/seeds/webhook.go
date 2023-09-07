package seeds

import (
	"go-ws-server/src/models"

	"gorm.io/gorm"
)

func Webhook(db *gorm.DB) error {
	user := []*models.Webhooks{
		{
			Name:      "Order",
			Value:     "order",
			CreatedBy: 1,
		},
		{
			Name:      "Pesanan",
			Value:     "pesanan",
			CreatedBy: 1,
		},
		{
			Name:      "Notifikasi",
			Value:     "notifikasi",
			CreatedBy: 1,
		},
	}
	rs := models.DB.Create(&user)

	if rs.Error != nil {
		return rs.Error
	}
	return nil
}
