package main

import (
	"fmt"

	"github.com/North-al/gin-template/config"
	"github.com/North-al/gin-template/internal/biz/repository"
	"github.com/North-al/gin-template/internal/data/models"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "internal/data/query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})

	dbConfig := config.GetConfig().Database

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)

	gormDB, _ := gorm.Open(mysql.New(mysql.Config{
		DSN:               dsn,
		DefaultStringSize: 256,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "t_", // 设置表名前缀
		},
	})

	g.UseDB(gormDB)

	g.ApplyBasic(models.User{})

	g.ApplyInterface(func(repository.UserRepositoryGen) {}, models.User{})
	// g.ApplyInterface(func(repository.RoleRepository) {}, model.Role{})
	// g.ApplyInterface(func(repository.DepartmentRepository) {}, model.Department{})

	g.Execute()
}
