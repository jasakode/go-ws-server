package middleware

import "gorm.io/gorm"

var CreateGormQuery = func(keys []string, obj interface{}) (*gorm.DB, error) {
	var a gorm.DB

	return &a, nil
}
