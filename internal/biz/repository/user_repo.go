package repository

import "github.com/North-al/gin-template/internal/biz/entity"

type UserRepository interface {
	Create(user *entity.UserEntity) (uint, error)
}
