package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/North-al/gin-template/internal/biz/entity"
	"github.com/North-al/gin-template/internal/biz/repository"
	"github.com/North-al/gin-template/internal/data/models"
	"github.com/North-al/gin-template/internal/pkg"
	"github.com/North-al/gin-template/internal/pkg/utils"
	"gorm.io/gorm"
)

type AuthService struct {
	userRepo repository.IUserRepo
}

func NewAuthService(userRepo repository.IUserRepo) *AuthService {
	return &AuthService{userRepo: userRepo}
}

func (s *AuthService) Register(ctx context.Context, req entity.RegisterRequest) (int64, error) {
	_, err := s.userRepo.FindByPhone(ctx, req.Phone)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// 说明这个用户没有注册
		// 这里可以做一些其他的校验，比如密码强度、用户名是否合法等
		if req.Password != req.ConfirmPassword {
			return 0, errors.New("两次输入的密码不一致")
		}

		// 这里可以做一些其他的校验，比如验证码是否正确等
		if req.Captcha != "000000" {
			return 0, errors.New("验证码错误")
		}

		// 密码加盐
		encryptPassword, err := utils.EncryptPassword(req.Password)
		if err != nil {
			return 0, err
		}

		// 创建用户
		user := &models.User{
			Username: req.Username,
			Phone:    &req.Phone,
			Password: string(encryptPassword),
		}

		// 保存用户
		if err := s.userRepo.Save(ctx, user); err != nil {
			return 0, err
		}

		return user.ID, nil

	}

	return 0, errors.New("用户已经注册")
}

func (s *AuthService) Login(ctx context.Context, req entity.LoginRequest) (string, error) {
	user, err := s.userRepo.FindByPhone(ctx, req.Phone)
	fmt.Println(req)
	if err != nil {
		return "", errors.New("用户名或密码错误")
	}

	if req.Type == "code" {
		fmt.Println("code")
		if req.Captcha != "000000" {
			return "", errors.New("验证码错误")
		}
	} else if req.Type == "password" {
		// 密码校验
		fmt.Println("password")
		err = utils.VerifyPassword(req.Password, user.Password)

		if err != nil {
			return "", errors.New("用户名或密码错误")
		}

	} else {
		return "", errors.New("异常登录类型")
	}

	// 下发 token
	token, err := pkg.GenerateToken(user.Username, user.ID)
	// TODO: 缓存 token

	return token, nil
}
