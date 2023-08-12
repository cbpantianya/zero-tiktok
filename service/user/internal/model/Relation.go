package model

type Relation struct {
	UserId     int `gorm:"primary_key;not null" json:"user_id"`
	FollowerId int `gorm:"primary_key;not null" json:"follower_id"`
}

func (r *Relation) TableName() string {
	return "relations"
}
