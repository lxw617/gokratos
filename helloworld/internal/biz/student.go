package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

// Student is a Student model.
type Student struct {
	Name string
	ID   int64
}

// StudentRepo is a Student repo.
type StudentRepo interface {
	Save(context.Context, *Student) (*Student, error)
	Get(context.Context, *Student) (*Student, error)
}

// StudentUsecase is a Student usecase.
type StudentUsecase struct {
	repo StudentRepo
	log  *log.Helper
}

// NewStudentUsecase new a Student usecase.
func NewStudentUsecase(repo StudentRepo, logger log.Logger) *StudentUsecase {
	return &StudentUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateStudent creates a Student, and returns the new Student.
func (uc *StudentUsecase) CreateStudent(ctx context.Context, g *Student) (*Student, error) {
	uc.log.WithContext(ctx).Infof("CreateStudent: %v", g.Name)
	return uc.repo.Save(ctx, g)
}
