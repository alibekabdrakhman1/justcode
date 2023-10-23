package postgre

import (
	"context"
	"fmt"
	"github.com/alibekabdrakhman/justcode/lecture9/models"
	"gorm.io/gorm"
)

func NewStudentRepository(db *gorm.DB) *StudentRepository {
	return &StudentRepository{
		DB: db,
	}
}

type StudentRepository struct {
	DB *gorm.DB
}

func (r *StudentRepository) GetAll(ctx context.Context) (error, []models.Student) {
	var resp []models.Student
	err := r.DB.WithContext(ctx).Find(&resp)
	return err.Error, resp
}

func (r *StudentRepository) GetById(ctx context.Context, id int) (error, models.Student) {
	var res models.Student
	err := r.DB.WithContext(ctx).Where("id = ?", id).Find(&res).Error
	return err, res
}

func (r *StudentRepository) Create(ctx context.Context, student models.Student) (error, int) {
	if err := r.DB.WithContext(ctx).Create(&student).Error; err != nil {
		return err, -1
	}
	return nil, student.Id
}

// kazhetsya neochen sdelal((((
func (r *StudentRepository) Update(ctx context.Context, student models.Student) (error, models.Student) {
	var existing models.Student
	if err := r.DB.WithContext(ctx).First(&existing, student.Id).Error; err != nil {
		fmt.Println("Entity not found")
		return err, models.Student{}
	}

	existing.Name = student.Name
	r.DB.Save(&existing)
	return nil, existing
}

func (r *StudentRepository) Delete(ctx context.Context, id int) error {
	return r.DB.WithContext(ctx).Delete(&models.Student{}, id).Error
}
