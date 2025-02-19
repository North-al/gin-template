package repository

import (
	"github.com/North-al/go-gateway/internal/biz/entity"
	"github.com/North-al/go-gateway/internal/biz/repository"
	"gorm.io/gorm"
)

type userRepoImpl struct {
	db *gorm.DB
}

// Create implements repository.UserRepository.
func (u *userRepoImpl) Create(user *entity.UserEntity) (uint, error) {
	panic("unimplemented")
}

func NewUserRepoImpl(db *gorm.DB) repository.UserRepository {
	return &userRepoImpl{db: db}
}
