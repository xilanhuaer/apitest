package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID          uint   `gorm:"primary_key" json:"id"`
	Account     string `gorm:"unique;not null" json:"account"`
	Password    string `gorm:"not null" json:"password"`
	Name        string `gorm:"unique;not null" json:"name"`
	Avatar      string `json:"avatar"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Description string `json:"description"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func (User) TableName() string {
	return "user"
}
func (u *User) AfterFind(tx *gorm.DB) (err error) {
	if u.Phone != "" && len(u.Phone) == 11 {
		u.Phone = u.Phone[:3] + "****" + u.Phone[:7]
	}
	return
}
