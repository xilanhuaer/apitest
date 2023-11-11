// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameUser = "user"

// User mapped from table <user>
type User struct {
	ID          int32          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Account     string         `gorm:"column:account;not null;comment:用户账号" json:"account"`                        // 用户账号
	Password    string         `gorm:"column:password;not null;comment:用户密码" json:"password"`                      // 用户密码
	Name        string         `gorm:"column:name;not null;comment:用户名" json:"name"`                               // 用户名
	Avatar      string         `gorm:"column:avatar;comment:头像" json:"avatar"`                                     // 头像
	Email       string         `gorm:"column:email;comment:邮箱" json:"email"`                                       // 邮箱
	Phone       string         `gorm:"column:phone;comment:手机号" json:"phone"`                                      // 手机号
	Description string         `gorm:"column:description;comment:描述" json:"description"`                           // 描述
	CreatedAt   time.Time      `gorm:"column:created_at;default:CURRENT_TIMESTAMP;comment:创建时间" json:"created_at"` // 创建时间
	UpdatedAt   time.Time      `gorm:"column:updated_at;default:CURRENT_TIMESTAMP;comment:更新时间" json:"updated_at"` // 更新时间
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;comment:删除时间" json:"deleted_at"`                           // 删除时间
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}