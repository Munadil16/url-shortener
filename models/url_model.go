package models

import "gorm.io/gorm"

type Url struct {
	gorm.Model
	LongUrl    string
	ShortUrlId string
	Visits     string
}
