package postgre

import (
	"context"
	"fmt"
	"github.com/alibekabdrakhman/justcode/lecture9/models"
	"gorm.io/gorm"
)

func NewSubjectTeacher(db *gorm.DB) *SubjectRepository {
	return &SubjectRepository{
		DB: db,
	}
}

type SubjectRepository struct {
	DB *gorm.DB
}

func (r *SubjectRepository) GetAll(ctx context.Context) (error, []models.Subject) {
	var resp []models.Subject
	err := r.DB.WithContext(ctx).Find(&resp)
	return err.Error, resp
}

func (r *SubjectRepository) GetById(ctx context.Context, id int) (error, models.Subject) {
	var res models.Subject
	err := r.DB.WithContext(ctx).Where("id = ?", id).Find(&res).Error
	return err, res
}

func (r *SubjectRepository) Create(ctx context.Context, subject models.Subject) (error, int) {
	if err := r.DB.WithContext(ctx).Create(&subject).Error; err != nil {
		return err, -1
	}
	return nil, subject.Id
}

func (r *SubjectRepository) Update(ctx context.Context, subject models.Subject) (error, models.Subject) {
	var existing models.Subject
	if err := r.DB.WithContext(ctx).First(&existing, subject.Id).Error; err != nil {
		fmt.Println("Entity not found")
		return err, models.Subject{}
	}

	existing.Teacher = subject.Teacher
	r.DB.Save(&existing)
	return nil, existing
}

func (r *SubjectRepository) Delete(ctx context.Context, id int) error {
	return r.DB.WithContext(ctx).Delete(&models.Subject{}, id).Error
}
