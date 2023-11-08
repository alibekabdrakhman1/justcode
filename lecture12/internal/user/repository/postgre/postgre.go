package postgre

import (
	"context"
	"fmt"
	"github.com/alibekabdrakhman/justcode/lecture12/internal/user/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Dial(ctx context.Context, url string) (*gorm.DB, error) {
	fmt.Println(url)
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		return nil, err
	}
	if db != nil {
		err := db.WithContext(ctx).AutoMigrate(&model.User{})
		if err != nil {
			return nil, err
		}
	}

	return db, nil
}
