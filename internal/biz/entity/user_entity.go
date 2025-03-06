package entity

import "time"

// UserEntity 用户实体
type UserEntity struct {
	ID         int64      `json:"id"`
	Avatar     string     `json:"avatar"`
	Username   string     `json:"username"`
	Password   string     `json:"-"`
	Phone      *string    `json:"phone"`
	Gender     int8       `json:"gender"`
	GenderText string     `json:"genderText"`
	Birthday   *time.Time `json:"birthday,omitempty"`
	Status     int8       `json:"status,omitempty"`
	StatusText string     `json:"statusText"`
	CreatedAt  time.Time  `json:"createdAt"`
	UpdatedAt  time.Time  `json:"updatedAt"`
}
