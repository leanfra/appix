// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.8.2
// - protoc             v3.12.4
// source: opspillar/v1/envs.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationEnvsCreateEnvs = "/api.opspillar.v1.Envs/CreateEnvs"
const OperationEnvsDeleteEnvs = "/api.opspillar.v1.Envs/DeleteEnvs"
const OperationEnvsGetEnvs = "/api.opspillar.v1.Envs/GetEnvs"
const OperationEnvsListEnvs = "/api.opspillar.v1.Envs/ListEnvs"
const OperationEnvsUpdateEnvs = "/api.opspillar.v1.Envs/UpdateEnvs"

type EnvsHTTPServer interface {
	CreateEnvs(context.Context, *CreateEnvsRequest) (*CreateEnvsReply, error)
	DeleteEnvs(context.Context, *DeleteEnvsRequest) (*DeleteEnvsReply, error)
	GetEnvs(context.Context, *GetEnvsRequest) (*GetEnvsReply, error)
	ListEnvs(context.Context, *ListEnvsRequest) (*ListEnvsReply, error)
	UpdateEnvs(context.Context, *UpdateEnvsRequest) (*UpdateEnvsReply, error)
}

func RegisterEnvsHTTPServer(s *http.Server, srv EnvsHTTPServer) {
	r := s.Route("/")
	r.POST("/api/v1/envs/create", _Envs_CreateEnvs0_HTTP_Handler(srv))
	r.POST("/api/v1/envs/update", _Envs_UpdateEnvs0_HTTP_Handler(srv))
	r.POST("/api/v1/envs/delete", _Envs_DeleteEnvs0_HTTP_Handler(srv))
	r.GET("/api/v1/envs/{id}", _Envs_GetEnvs0_HTTP_Handler(srv))
	r.POST("/api/v1/envs/list", _Envs_ListEnvs0_HTTP_Handler(srv))
}

func _Envs_CreateEnvs0_HTTP_Handler(srv EnvsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateEnvsRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationEnvsCreateEnvs)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateEnvs(ctx, req.(*CreateEnvsRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateEnvsReply)
		return ctx.Result(200, reply)
	}
}

func _Envs_UpdateEnvs0_HTTP_Handler(srv EnvsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateEnvsRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationEnvsUpdateEnvs)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateEnvs(ctx, req.(*UpdateEnvsRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateEnvsReply)
		return ctx.Result(200, reply)
	}
}

func _Envs_DeleteEnvs0_HTTP_Handler(srv EnvsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteEnvsRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationEnvsDeleteEnvs)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteEnvs(ctx, req.(*DeleteEnvsRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteEnvsReply)
		return ctx.Result(200, reply)
	}
}

func _Envs_GetEnvs0_HTTP_Handler(srv EnvsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetEnvsRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationEnvsGetEnvs)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetEnvs(ctx, req.(*GetEnvsRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetEnvsReply)
		return ctx.Result(200, reply)
	}
}

func _Envs_ListEnvs0_HTTP_Handler(srv EnvsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListEnvsRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationEnvsListEnvs)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListEnvs(ctx, req.(*ListEnvsRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListEnvsReply)
		return ctx.Result(200, reply)
	}
}

type EnvsHTTPClient interface {
	CreateEnvs(ctx context.Context, req *CreateEnvsRequest, opts ...http.CallOption) (rsp *CreateEnvsReply, err error)
	DeleteEnvs(ctx context.Context, req *DeleteEnvsRequest, opts ...http.CallOption) (rsp *DeleteEnvsReply, err error)
	GetEnvs(ctx context.Context, req *GetEnvsRequest, opts ...http.CallOption) (rsp *GetEnvsReply, err error)
	ListEnvs(ctx context.Context, req *ListEnvsRequest, opts ...http.CallOption) (rsp *ListEnvsReply, err error)
	UpdateEnvs(ctx context.Context, req *UpdateEnvsRequest, opts ...http.CallOption) (rsp *UpdateEnvsReply, err error)
}

type EnvsHTTPClientImpl struct {
	cc *http.Client
}

func NewEnvsHTTPClient(client *http.Client) EnvsHTTPClient {
	return &EnvsHTTPClientImpl{client}
}

func (c *EnvsHTTPClientImpl) CreateEnvs(ctx context.Context, in *CreateEnvsRequest, opts ...http.CallOption) (*CreateEnvsReply, error) {
	var out CreateEnvsReply
	pattern := "/api/v1/envs/create"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationEnvsCreateEnvs))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *EnvsHTTPClientImpl) DeleteEnvs(ctx context.Context, in *DeleteEnvsRequest, opts ...http.CallOption) (*DeleteEnvsReply, error) {
	var out DeleteEnvsReply
	pattern := "/api/v1/envs/delete"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationEnvsDeleteEnvs))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *EnvsHTTPClientImpl) GetEnvs(ctx context.Context, in *GetEnvsRequest, opts ...http.CallOption) (*GetEnvsReply, error) {
	var out GetEnvsReply
	pattern := "/api/v1/envs/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationEnvsGetEnvs))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *EnvsHTTPClientImpl) ListEnvs(ctx context.Context, in *ListEnvsRequest, opts ...http.CallOption) (*ListEnvsReply, error) {
	var out ListEnvsReply
	pattern := "/api/v1/envs/list"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationEnvsListEnvs))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *EnvsHTTPClientImpl) UpdateEnvs(ctx context.Context, in *UpdateEnvsRequest, opts ...http.CallOption) (*UpdateEnvsReply, error) {
	var out UpdateEnvsReply
	pattern := "/api/v1/envs/update"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationEnvsUpdateEnvs))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
