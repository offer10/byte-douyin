package model

type User struct {
	BaseModel
	Username string  `gorm:"size:64;not null;uniqueIndex" json:"username"`
	Nickname string  `gorm:"size:64" json:"nickname"`
	Password string  `gorm:"size:255;not null" json:"-"`
	Follow   []*User `gorm:"many2many:user_follow" json:"follow"`
	Follower []*User `gorm:"many2many:user_follower" json:"follower"`
}
