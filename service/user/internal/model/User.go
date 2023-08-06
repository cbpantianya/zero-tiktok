package model

type User struct {
	UserId    int    `gorm:"primary_key;auto_increment;not null" json:"user_id"`
	Name      string `gorm:"not null" json:"name"`
	Signature string `gorm:"default '这个人很懒，什么都没有留下';not null" json:"signature"`
	Cover     string `gorm:"not null;comment '用户个人页顶部大图'" json:"cover"`
	Avatar    string `gorm:"not null" json:"avatar"`
	Pass      string `gorm:"not null" json:"pass"`
}
