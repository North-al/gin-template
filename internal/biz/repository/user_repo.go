package repository

import "github.com/North-al/go-gateway/internal/biz/entity"

type UserRepository interface {
	Create(user *entity.UserEntity) (uint, error)
}
