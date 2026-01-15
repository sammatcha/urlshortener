package models

import "time"

type Url struct {
	URL_ID		uint		`gorm:"primaryKey;column:url_id"`
	CreatedAt	time.Time	`gorm:"column:created_at"`
	LongURL 	string		`gorm:"column:long_url"`
	ShortURL 	string		`gorm:"column:short_url"`
}