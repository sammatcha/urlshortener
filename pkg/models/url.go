package models

import "time"

type Url struct {
	URL_ID		uint		`gorm:"primaryKey"`
	CreatedAt	time.Time
	LongURL 	string
	ShortURL 	string
}