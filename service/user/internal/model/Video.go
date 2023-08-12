package model

import "time"

type Video struct {
	VideoId   int       `gorm:"primary_key;auto_increment;not null" json:"user_id"`
	PublishAt time.Time `gorm:"not null" json:"publish_at"`
	AuthorID  int       `gorm:"not null" json:"author_id"`
	Play      string    `gorm:"not null" json:"play"`
	Cover     string    `gorm:"not null" json:"cover"`
	Title     string    `gorm:"not null" json:"title"`
}

func (v *Video) TableName() string {
	return "videos"
}
