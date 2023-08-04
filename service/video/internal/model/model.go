package model

import "time"

type Video struct {
	VideoID   int64     `gorm:"primary_key;auto_increment;not null" json:"video_id"`
	PublishAt time.Time `gorm:"not null" json:"publish_at"`
	AuthorID  int64     `gorm:"not null" json:"author_id"`
	Play      string    `gorm:"not null" json:"play"`
	Cover     string    `gorm:"not null" json:"cover"`
	Title     string    `gorm:"not null" json:"title"`
}

type Favorite struct {
	VideoID int64 `gorm:"primary_key;not null" json:"video_id"`
	UserID  int64 `gorm:"primary_key;not null" json:"author_id"`
}

type Comment struct {
	VideoID     int64
	UserID      int64
	CommentText string
	CommentID   int64
}
