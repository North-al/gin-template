package service

import (
	"github.com/North-al/gin-template/internal/biz/entity"
	"github.com/North-al/gin-template/internal/biz/repository"
)

type UserService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) Create(user *entity.UserEntity) (uint, error) {
	return s.userRepo.Create(user)
}
