package repository

import (
	"context"
	"github.com/alibekabdrakhman/justcode/lecture9/models"
	"github.com/alibekabdrakhman/justcode/lecture9/repository/postgre"
)

type IStudentRepository interface {
	GetAll(ctx context.Context) (error, []models.Student)
	GetById(ctx context.Context, id int) (error, models.Student)
	Create(ctx context.Context, student models.Student) (error, int)
	Update(ctx context.Context, student models.Student) (error, models.Student)
	Delete(ctx context.Context, id int) error
}
type ITeacherRepository interface {
	GetAll(ctx context.Context) (error, []models.Teacher)
	GetById(ctx context.Context, id int) (error, models.Teacher)
	Create(ctx context.Context, teacher models.Teacher) (error, int)
	Update(ctx context.Context, teacher models.Teacher) (error, models.Teacher)
	Delete(ctx context.Context, id int) error
}
type ISubjectRepository interface {
	GetAll(ctx context.Context) (error, []models.Subject)
	GetById(ctx context.Context, id int) (error, models.Subject)
	Create(ctx context.Context, subject models.Subject) (error, int)
	Update(ctx context.Context, subject models.Subject) (error, models.Subject)
	Delete(ctx context.Context, id int) error
}

type Storage struct {
	Student IStudentRepository
	Teacher ITeacherRepository
	Subject ISubjectRepository
}

func NewStorage(ctx context.Context) (*Storage, error) {
	pgUrl := "host=localhost port=5432 user=postgres password=qwerty dbname=justcode sslmode=disable"
	DB, err := postgre.Dial(ctx, pgUrl)
	if err != nil {
		return nil, err
	}

	storage := Storage{
		Student: postgre.NewStudentRepository(DB),
		Teacher: postgre.NewTeacherRepository(DB),
		Subject: postgre.NewSubjectTeacher(DB),
	}
	return &storage, nil

}
