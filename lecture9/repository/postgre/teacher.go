package postgre

import (
	"context"
	"fmt"
	"github.com/alibekabdrakhman/justcode/lecture9/models"
	"gorm.io/gorm"
)

func NewTeacherRepository(db *gorm.DB) *TeacherRepository {
	return &TeacherRepository{
		DB: db,
	}
}

type TeacherRepository struct {
	DB *gorm.DB
}

func (r *TeacherRepository) GetAll(ctx context.Context) (error, []models.Teacher) {
	var resp []models.Teacher
	err := r.DB.WithContext(ctx).Find(&resp)
	return err.Error, resp
}

func (r *TeacherRepository) GetById(ctx context.Context, id int) (error, models.Teacher) {
	var res models.Teacher
	err := r.DB.WithContext(ctx).Where("id = ?", id).Find(&res).Error
	return err, res
}

func (r *TeacherRepository) Create(ctx context.Context, teacher models.Teacher) (error, int) {
	if err := r.DB.WithContext(ctx).Create(&teacher).Error; err != nil {
		return err, -1
	}
	return nil, teacher.Id
}

func (r *TeacherRepository) Update(ctx context.Context, teacher models.Teacher) (error, models.Teacher) {
	var existing models.Teacher
	if err := r.DB.WithContext(ctx).First(&existing, teacher.Id).Error; err != nil {
		fmt.Println("Entity not found")
		return err, models.Teacher{}
	}

	existing.Name = teacher.Name
	r.DB.Save(&existing)
	return nil, existing
}

func (r *TeacherRepository) Delete(ctx context.Context, id int) error {
	return r.DB.WithContext(ctx).Delete(&models.Teacher{}, id).Error
}
