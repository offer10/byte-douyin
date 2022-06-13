package model

type User struct {
	BaseModel
	Username      string  `gorm:"size:64;not null;uniqueIndex" json:"username"`
	Nickname      string  `gorm:"size:64" json:"nickname"`
	Password      string  `gorm:"size:255;not null" json:"-"`
	FollowCount   int64   `gorm:"column:follow_count"`
	FollowerCount int64   `gorm:"column:follower_count"`
	Follows       []*User `gorm:"joinForeignKey:UserId;many2many:relations" json:"follow"`
	Followers     []*User `gorm:"joinForeignKey:FollowId;many2many:relations" json:"follower"`
}
