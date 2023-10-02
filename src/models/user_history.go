package models

type User_historys struct {
	ID          int64  `gorm:"primaryKey, not null, index, uniqueIndex, autoIncrement" json:"id"`
	UserId      int64  `json:"user_id"`
	Users       Users  `gorm:"foreignKey:user_id;joinForeignKey:id" json:"users"`
	Action      string `json:"action"`
	Description string `json:"description"`
	CreatedBy   int64  `gorm:"not null" json:"created_by"`
	UpdatedBy   int64  `json:"updated_by"`
	DeletedBy   int64  `json:"deleted_by"`
	CreatedAt   int64  `gorm:"not null, autoCreateTime:milli;" json:"created_at"`
	UpdatedAt   int64  `gorm:"autoCreateTime:milli;" json:"updated_at"`
	DeletedAt   int64  `json:"deleted_at"`
}
