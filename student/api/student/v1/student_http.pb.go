// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.5.0
// - protoc             v3.20.1
// source: api/student/v1/student.proto

package v1

import (
	context "context"

	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var (
	_ = new(context.Context)
	_ = binding.EncodeURL
)

const _ = http.SupportPackageIsVersion1

const (
	OperationStudentCreateStudent = "/api.student.v1.Student/CreateStudent"
	OperationStudentDeleteStudent = "/api.student.v1.Student/DeleteStudent"
	OperationStudentGetStudent    = "/api.student.v1.Student/GetStudent"
	OperationStudentListStudent   = "/api.student.v1.Student/ListStudent"
	OperationStudentUpdateStudent = "/api.student.v1.Student/UpdateStudent"
)

type StudentHTTPServer interface {
	CreateStudent(context.Context, *CreateStudentRequest) (*CreateStudentReply, error)
	DeleteStudent(context.Context, *DeleteStudentRequest) (*DeleteStudentReply, error)
	GetStudent(context.Context, *GetStudentRequest) (*GetStudentReply, error)
	ListStudent(context.Context, *ListStudentRequest) (*ListStudentReply, error)
	UpdateStudent(context.Context, *UpdateStudentRequest) (*UpdateStudentReply, error)
}

func RegisterStudentHTTPServer(s *http.Server, srv StudentHTTPServer) {
	r := s.Route("/")
	r.POST("/student", _Student_CreateStudent0_HTTP_Handler(srv))
	r.PUT("/student/{id}", _Student_UpdateStudent0_HTTP_Handler(srv))
	r.DELETE("/student/{id}", _Student_DeleteStudent0_HTTP_Handler(srv))
	r.GET("/student/{id}", _Student_GetStudent0_HTTP_Handler(srv))
	r.GET("/students", _Student_ListStudent0_HTTP_Handler(srv))
}

func _Student_CreateStudent0_HTTP_Handler(srv StudentHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateStudentRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationStudentCreateStudent)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateStudent(ctx, req.(*CreateStudentRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateStudentReply)
		return ctx.Result(200, reply)
	}
}

func _Student_UpdateStudent0_HTTP_Handler(srv StudentHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateStudentRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationStudentUpdateStudent)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateStudent(ctx, req.(*UpdateStudentRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateStudentReply)
		return ctx.Result(200, reply)
	}
}

func _Student_DeleteStudent0_HTTP_Handler(srv StudentHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteStudentRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationStudentDeleteStudent)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteStudent(ctx, req.(*DeleteStudentRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteStudentReply)
		return ctx.Result(200, reply)
	}
}

func _Student_GetStudent0_HTTP_Handler(srv StudentHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetStudentRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationStudentGetStudent)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetStudent(ctx, req.(*GetStudentRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetStudentReply)
		return ctx.Result(200, reply)
	}
}

func _Student_ListStudent0_HTTP_Handler(srv StudentHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListStudentRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationStudentListStudent)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListStudent(ctx, req.(*ListStudentRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListStudentReply)
		return ctx.Result(200, reply)
	}
}

type StudentHTTPClient interface {
	CreateStudent(ctx context.Context, req *CreateStudentRequest, opts ...http.CallOption) (rsp *CreateStudentReply, err error)
	DeleteStudent(ctx context.Context, req *DeleteStudentRequest, opts ...http.CallOption) (rsp *DeleteStudentReply, err error)
	GetStudent(ctx context.Context, req *GetStudentRequest, opts ...http.CallOption) (rsp *GetStudentReply, err error)
	ListStudent(ctx context.Context, req *ListStudentRequest, opts ...http.CallOption) (rsp *ListStudentReply, err error)
	UpdateStudent(ctx context.Context, req *UpdateStudentRequest, opts ...http.CallOption) (rsp *UpdateStudentReply, err error)
}

type StudentHTTPClientImpl struct {
	cc *http.Client
}

func NewStudentHTTPClient(client *http.Client) StudentHTTPClient {
	return &StudentHTTPClientImpl{client}
}

func (c *StudentHTTPClientImpl) CreateStudent(ctx context.Context, in *CreateStudentRequest, opts ...http.CallOption) (*CreateStudentReply, error) {
	var out CreateStudentReply
	pattern := "/student"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationStudentCreateStudent))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *StudentHTTPClientImpl) DeleteStudent(ctx context.Context, in *DeleteStudentRequest, opts ...http.CallOption) (*DeleteStudentReply, error) {
	var out DeleteStudentReply
	pattern := "/student/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationStudentDeleteStudent))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *StudentHTTPClientImpl) GetStudent(ctx context.Context, in *GetStudentRequest, opts ...http.CallOption) (*GetStudentReply, error) {
	var out GetStudentReply
	pattern := "/student/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationStudentGetStudent))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *StudentHTTPClientImpl) ListStudent(ctx context.Context, in *ListStudentRequest, opts ...http.CallOption) (*ListStudentReply, error) {
	var out ListStudentReply
	pattern := "/students"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationStudentListStudent))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *StudentHTTPClientImpl) UpdateStudent(ctx context.Context, in *UpdateStudentRequest, opts ...http.CallOption) (*UpdateStudentReply, error) {
	var out UpdateStudentReply
	pattern := "/student/{id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationStudentUpdateStudent))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
