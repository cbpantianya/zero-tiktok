package model

type User struct {
	ID        int    `gorm:"column:user_id" json:"user_id"`
	Name      string `gorm:"column:name" json:"name"`
	Signature string `gorm:"column:signature" json:"signature"`
	Cover     string `gorm:"column:cover" json:"cover"`
	Avatar    string `gorm:"column:avatar" json:"avatar"`
	Pass      string `gorm:"column:pass" json:"pass"`
	Salt      string `gorm:"column:salt" json:"salt"`
}
