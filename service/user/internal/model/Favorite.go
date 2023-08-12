package model

type Favorite struct {
	VideoId int `gorm:"primary_key;not null" json:"video_id"`
	UserId  int `gorm:"primary_key;not null" json:"user_id"`
}

func (f *Favorite) TableName() string {
	return "favorites"
}
