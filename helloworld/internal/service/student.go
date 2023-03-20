package service

import (
	"context"
	"helloworld/internal/biz"

	pb "helloworld/api/helloworld/v1"
)

type StudentService struct {
	pb.UnimplementedStudentServer
	uc *biz.StudentUsecase
}

func NewStudentService(uc *biz.StudentUsecase) *StudentService {
	return &StudentService{uc: uc}
}

func (s *StudentService) CreateStudent(ctx context.Context, req *pb.CreateStudentRequest) (*pb.CreateStudentReply, error) {
	return &pb.CreateStudentReply{}, nil
}
func (s *StudentService) UpdateStudent(ctx context.Context, req *pb.UpdateStudentRequest) (*pb.UpdateStudentReply, error) {
	return &pb.UpdateStudentReply{}, nil
}
func (s *StudentService) DeleteStudent(ctx context.Context, req *pb.DeleteStudentRequest) (*pb.DeleteStudentReply, error) {
	return &pb.DeleteStudentReply{}, nil
}
func (s *StudentService) GetStudent(ctx context.Context, req *pb.GetStudentRequest) (*pb.GetStudentReply, error) {
	return &pb.GetStudentReply{}, nil
}
func (s *StudentService) ListStudent(ctx context.Context, req *pb.ListStudentRequest) (*pb.ListStudentReply, error) {
	return &pb.ListStudentReply{}, nil
}
func (s *StudentService) HelloStudent(ctx context.Context, req *pb.HelloStudentRequest) (*pb.HelloStudentReply, error) {
	g, err := s.uc.CreateStudent(ctx, &biz.Student{Name: req.Name, ID: req.Id})
	if err != nil {
		return nil, err
	}
	return &pb.HelloStudentReply{Name: "I am student " + g.Name + ", my number is " + string(g.ID)}, nil
}
