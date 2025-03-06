package repository

import (
	"context"

	"github.com/North-al/gin-template/internal/biz/entity"
	"github.com/North-al/gin-template/internal/biz/repository"
	"github.com/North-al/gin-template/internal/data/models"
	"github.com/North-al/gin-template/internal/data/query"
	"gorm.io/gorm"
)

type userRepoImpl struct {
	db    *gorm.DB
	query *query.Query
}

func NewUserRepository(db *gorm.DB) repository.IUserRepo {
	return &userRepoImpl{db: db, query: query.Use(db)}
}

func (u *userRepoImpl) FindById(ctx context.Context, id int64) (*entity.UserEntity, error) {
	return u.query.User.FindById(ctx, id)
}

func (u *userRepoImpl) FindByPhone(ctx context.Context, phone string) (*entity.UserEntity, error) {
	return u.query.User.WithContext(ctx).FindByPhone(ctx, phone)
}

func (u *userRepoImpl) Save(ctx context.Context, user *models.User) error {
	// gen 默认生成的方法，没有ctx传递
	return u.query.User.WithContext(ctx).Create(user)
}
