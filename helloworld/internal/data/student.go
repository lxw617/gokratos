package data

import (
	"context"

	"helloworld/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type studentRepo struct {
	data *Data
	log  *log.Helper
}

// NewStudentRepo .
func NewStudentRepo(data *Data, logger log.Logger) biz.StudentRepo {
	return &studentRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *studentRepo) Save(ctx context.Context, s *biz.Student) (*biz.Student, error) {
	return s, nil
}

func (r *studentRepo) Get(ctx context.Context, s *biz.Student) (*biz.Student, error) {
	return s, nil
}
