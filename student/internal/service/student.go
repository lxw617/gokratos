package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"student/internal/biz"

	pb "student/api/student/v1"
)

type StudentService struct {
	pb.UnimplementedStudentServer

	student *biz.StudentUsecase
	log     *log.Helper
}

// 初始化
func NewStudentService(stu *biz.StudentUsecase, logger log.Logger) *StudentService {
	return &StudentService{
		student: stu,
		log:     log.NewHelper(logger),
	}
}

func (s *StudentService) CreateStudent(ctx context.Context, req *pb.CreateStudentRequest) (*pb.CreateStudentReply, error) {
	stu, err := s.student.Create(ctx, &biz.Student{
		Name:   req.Name,
		Info:   req.Info,
		Status: req.Status,
	})
	if err != nil {
		return nil, err
	}
	return &pb.CreateStudentReply{
		Id:     stu.ID,
		Status: stu.Status,
		Name:   stu.Name,
		Info:   stu.Info,
	}, nil
}

func (s *StudentService) UpdateStudent(ctx context.Context, req *pb.UpdateStudentRequest) (*pb.UpdateStudentReply, error) {
	stu, err := s.student.Update(ctx, req.Id, &biz.Student{
		ID:     req.Id,
		Name:   req.Name,
		Info:   req.Info,
		Status: req.Status,
	})
	if err != nil {
		return nil, err
	}
	return &pb.UpdateStudentReply{
		Id:     stu.ID,
		Status: stu.Status,
		Name:   stu.Name,
		Info:   stu.Info,
	}, nil
}

func (s *StudentService) DeleteStudent(ctx context.Context, req *pb.DeleteStudentRequest) (*pb.DeleteStudentReply, error) {
	err := s.student.Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteStudentReply{}, nil
}

// 获取学生信息
func (s *StudentService) GetStudent(ctx context.Context, req *pb.GetStudentRequest) (*pb.GetStudentReply, error) {
	stu, err := s.student.Get(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.GetStudentReply{
		Id:     stu.ID,
		Status: stu.Status,
		Name:   stu.Name,
		Info:   stu.Info,
	}, nil
}

func (s *StudentService) ListStudent(ctx context.Context, req *pb.ListStudentRequest) (*pb.ListStudentReply, error) {
	studentList, err := s.student.List(ctx, req.Name)
	if err != nil {
		return nil, err
	}
	resultList := make([]*pb.ListStudentReply_Result, 0)
	for _, stu := range studentList {
		resultList = append(resultList, &pb.ListStudentReply_Result{
			Id:     stu.ID,
			Status: stu.Status,
			Name:   stu.Name,
			Info:   stu.Info,
		})
	}
	return &pb.ListStudentReply{
		Results: resultList,
	}, nil
}
