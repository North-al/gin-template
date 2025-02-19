package entity

// User 用户实体
type UserEntity struct {
	Id       uint   `json:"id"`
	UserName string `json:"username"`
	Password string `json:"-"`
}
