package models

type Webhooks struct {
	ID    int64  `gorm:"primaryKey, not null, index, uniqueIndex, autoIncrement" json:"id"`
	Name  string `json:"name"`
	Value string `gorm:"primaryKey, not null" json:"value"`

	CreatedBy int64 `gorm:"not null" json:"created_by"`
	UpdatedBy int64 `json:"updated_by"`
	DeletedBy int64 `json:"deleted_by"`
	CreatedAt int64 `gorm:"not null, autoCreateTime:milli;" json:"created_at"`
	UpdatedAt int64 `gorm:"autoCreateTime:milli;" json:"updated_at"`
	DeletedAt int64 `json:"deleted_at"`
}

func (filter Webhooks) FindAll() (*findAllWebhook, error) {
	var respon findAllWebhook
	err := DB.Model(Webhooks{}).Count(&respon.Count).Find(&respon.Webhooks).Error
	if err != nil {
		return nil, err
	}
	return &respon, nil
}
