package model

import "time"

type Comment struct {
	VideoID     int64     `gorm:"not null" json:"video_id"`
	UserID      int64     `gorm:"not null" json:"user_id"`
	CommentText string    `gorm:"not null" json:"comment_text"`
	CommentID   int64     `gorm:"not null" json:"comment_id"`
	CreatedAt   time.Time `grom:"not null" json:"created_at"`
}
type Video struct {
	VideoID   int64     `gorm:"primary_key;auto_increment;not null" json:"video_id"`
	PublishAt time.Time `gorm:"not null" json:"publish_at"`
	AuthorID  int64     `gorm:"not null" json:"author_id"`
	Play      string    `gorm:"not null" json:"play"`
	Cover     string    `gorm:"not null" json:"cover"`
	Title     string    `gorm:"not null" json:"title"`
}

type Relation struct {
	UserID     int64 `gorm:"primary_key;not null" json:"user_id"`
	FollowerID int64 `gorm:"not null" json:"follower_id"`
}
