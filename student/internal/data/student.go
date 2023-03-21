package data

import (
	"context"

	"gorm.io/gorm"

	"student/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type studentRepo struct {
	data *Data
	log  *log.Helper
}

// 初始化 studentRepo
func NewStudentRepo(data *Data, logger log.Logger) biz.StudentRepo {
	return &studentRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *studentRepo) GetStudent(ctx context.Context, id int32) (*biz.Student, error) {
	var stu biz.Student
	if err := r.data.gormDB.Where("id = ?", id).First(&stu).Error; err != nil {
		return nil, err
	}
	r.log.WithContext(ctx).Info("gormDB: GetStudent, id: ", id)
	return &biz.Student{
		ID:        stu.ID,
		Name:      stu.Name,
		Status:    stu.Status,
		Info:      stu.Info,
		UpdatedAt: stu.UpdatedAt,
		CreatedAt: stu.CreatedAt,
	}, nil
}

func (r *studentRepo) ListStudent(name string) ([]*biz.Student, error) {
	lists := []*biz.Student{}
	scopes := make([]func(*gorm.DB) *gorm.DB, 0)

	if name != "" {
		scopes = append(scopes, func(db *gorm.DB) *gorm.DB {
			return r.data.gormDB.Where("name like ?", "%"+name+"%")
		})
	}
	if err := r.data.gormDB.Model(&biz.Student{}).Scopes(scopes...).Scan(&lists).Error; err != nil {
		return nil, err
	}
	return lists, nil
}

func (r *studentRepo) DeleteStudent(id int32) error {
	if err := r.data.gormDB.Model(&biz.Student{}).Where("id = ?", id).Delete(&biz.Student{}).Error; err != nil {
		return err
	}
	return nil
}

func (r *studentRepo) UpdateStudent(id int32, stu *biz.Student) (*biz.Student, error) {
	_student := map[string]interface{}{
		"name":   stu.Name,
		"info":   stu.Info,
		"status": stu.Status,
	}
	if err := r.data.gormDB.Model(&biz.Student{}).Where(id).Updates(_student).Scan(stu).Error; err != nil {
		return nil, err
	}
	return stu, nil
}

func (r *studentRepo) CreateStudent(stu *biz.Student) (*biz.Student, error) {
	if err := r.data.gormDB.Model(&biz.Student{}).Create(&stu).Error; err != nil {
		return nil, err
	}
	return stu, nil
}
