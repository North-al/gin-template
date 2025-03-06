package models

import (
	"time"
	
	"gorm.io/gorm"
)

// User 用户实体
type User struct {
	ID        int64          `gorm:"primaryKey;autoIncrement;comment:用户ID"`
	Username  string         `gorm:"unique;size:50;not null;comment:用户名"`
	Password  string         `gorm:"size:255;not null;comment:加密存储的密码"`
	Phone     *string        `gorm:"unique;size:20;comment:手机号"`
	Avatar    string         `gorm:"size:255;comment:头像URL"`
	Gender    int8           `gorm:"default:0;comment:性别 0:未知 1:男 2:女"`
	Birthday  *time.Time     `gorm:"comment:生日"`
	Status    int8           `gorm:"default:1;comment:状态 1:正常 0:禁用"`
	CreatedAt time.Time      `gorm:"autoCreateTime;comment:账户创建时间"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime;comment:账户更新时间"`
	DeletedAt gorm.DeletedAt `gorm:"index;comment:软删除标记"`
}
