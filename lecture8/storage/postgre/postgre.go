package postgre

import (
	"github.com/alibekabdrakhman/justcode/lecture8/model"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Dial(c *gin.Context, url string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		return nil, err
	}
	if db != nil {
		err := db.WithContext(c).AutoMigrate(&model.User{})
		if err != nil {
			return nil, err
		}
	}

	return db, nil
}
