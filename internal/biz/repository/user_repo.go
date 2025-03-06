package repository

import (
	"context"

	"github.com/North-al/gin-template/internal/biz/entity"
	"github.com/North-al/gin-template/internal/data/models"
)

type IUserRepo interface {
	UserRepository
	UserRepositoryGen
}

type UserRepositoryGen interface {
	// SELECT * FROM @@table WHERE id = @id LIMIT 1
	FindById(ctx context.Context, id int64) (*entity.UserEntity, error)

	//  SELECT * FROM @@table WHERE phone = @phone LIMIT 1
	FindByPhone(ctx context.Context, phone string) (*entity.UserEntity, error)
}

type UserRepository interface {
	Save(ctx context.Context, user *models.User) error
}
