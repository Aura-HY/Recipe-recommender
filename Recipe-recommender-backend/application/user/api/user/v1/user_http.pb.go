// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.8.4
// - protoc             v5.29.3
// source: api/user/v1/user.proto

package user

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

const OperationUserServiceGetHealthRecord = "/user.UserService/GetHealthRecord"
const OperationUserServiceGetUserPreferences = "/user.UserService/GetUserPreferences"
const OperationUserServiceGetUserProfile = "/user.UserService/GetUserProfile"
const OperationUserServiceLoginUser = "/user.UserService/LoginUser"
const OperationUserServiceLogoutUser = "/user.UserService/LogoutUser"
const OperationUserServiceRegisterUser = "/user.UserService/RegisterUser"
const OperationUserServiceUpdateHealthRecord = "/user.UserService/UpdateHealthRecord"
const OperationUserServiceUpdateUserPreferences = "/user.UserService/UpdateUserPreferences"
const OperationUserServiceUpdateUserProfile = "/user.UserService/UpdateUserProfile"

type UserServiceHTTPServer interface {
	// GetHealthRecord 获取用户健康档案
	GetHealthRecord(context.Context, *GetHealthRecordRequest) (*HealthRecordResponse, error)
	// GetUserPreferences 获取用户饮食偏好
	GetUserPreferences(context.Context, *GetUserPreferencesRequest) (*UserPreferencesResponse, error)
	// GetUserProfile 获取用户个人信息
	GetUserProfile(context.Context, *GetUserProfileRequest) (*UserProfileResponse, error)
	// LoginUser 用户登录
	LoginUser(context.Context, *LoginRequest) (*LoginResponse, error)
	// LogoutUser 退出登录
	LogoutUser(context.Context, *LogoutRequest) (*LogoutResponse, error)
	// RegisterUser 用户注册
	RegisterUser(context.Context, *RegisterRequest) (*RegisterResponse, error)
	// UpdateHealthRecord 更新用户健康档案（体质、过敏史等）
	UpdateHealthRecord(context.Context, *UpdateHealthRecordRequest) (*HealthRecordResponse, error)
	// UpdateUserPreferences 更新用户饮食偏好
	UpdateUserPreferences(context.Context, *UpdateUserPreferencesRequest) (*UserPreferencesResponse, error)
	// UpdateUserProfile 更新用户信息
	UpdateUserProfile(context.Context, *UpdateUserProfileRequest) (*UserProfileResponse, error)
}

func RegisterUserServiceHTTPServer(s *http.Server, srv UserServiceHTTPServer) {
	r := s.Route("/")
	r.POST("/users/register", _UserService_RegisterUser0_HTTP_Handler(srv))
	r.POST("/users/login", _UserService_LoginUser0_HTTP_Handler(srv))
	r.GET("/users/profile", _UserService_GetUserProfile0_HTTP_Handler(srv))
	r.PUT("/users/profile", _UserService_UpdateUserProfile0_HTTP_Handler(srv))
	r.GET("/users/health", _UserService_GetHealthRecord0_HTTP_Handler(srv))
	r.PUT("/users/health", _UserService_UpdateHealthRecord0_HTTP_Handler(srv))
	r.GET("/users/preferences", _UserService_GetUserPreferences0_HTTP_Handler(srv))
	r.PUT("/users/preferences", _UserService_UpdateUserPreferences0_HTTP_Handler(srv))
	r.POST("/users/logout", _UserService_LogoutUser0_HTTP_Handler(srv))
}

func _UserService_RegisterUser0_HTTP_Handler(srv UserServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RegisterRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserServiceRegisterUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.RegisterUser(ctx, req.(*RegisterRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RegisterResponse)
		return ctx.Result(200, reply)
	}
}

func _UserService_LoginUser0_HTTP_Handler(srv UserServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LoginRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserServiceLoginUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.LoginUser(ctx, req.(*LoginRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LoginResponse)
		return ctx.Result(200, reply)
	}
}

func _UserService_GetUserProfile0_HTTP_Handler(srv UserServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetUserProfileRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserServiceGetUserProfile)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetUserProfile(ctx, req.(*GetUserProfileRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserProfileResponse)
		return ctx.Result(200, reply)
	}
}

func _UserService_UpdateUserProfile0_HTTP_Handler(srv UserServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateUserProfileRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserServiceUpdateUserProfile)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateUserProfile(ctx, req.(*UpdateUserProfileRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserProfileResponse)
		return ctx.Result(200, reply)
	}
}

func _UserService_GetHealthRecord0_HTTP_Handler(srv UserServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetHealthRecordRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserServiceGetHealthRecord)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetHealthRecord(ctx, req.(*GetHealthRecordRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*HealthRecordResponse)
		return ctx.Result(200, reply)
	}
}

func _UserService_UpdateHealthRecord0_HTTP_Handler(srv UserServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateHealthRecordRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserServiceUpdateHealthRecord)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateHealthRecord(ctx, req.(*UpdateHealthRecordRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*HealthRecordResponse)
		return ctx.Result(200, reply)
	}
}

func _UserService_GetUserPreferences0_HTTP_Handler(srv UserServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetUserPreferencesRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserServiceGetUserPreferences)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetUserPreferences(ctx, req.(*GetUserPreferencesRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserPreferencesResponse)
		return ctx.Result(200, reply)
	}
}

func _UserService_UpdateUserPreferences0_HTTP_Handler(srv UserServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateUserPreferencesRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserServiceUpdateUserPreferences)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateUserPreferences(ctx, req.(*UpdateUserPreferencesRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserPreferencesResponse)
		return ctx.Result(200, reply)
	}
}

func _UserService_LogoutUser0_HTTP_Handler(srv UserServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LogoutRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserServiceLogoutUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.LogoutUser(ctx, req.(*LogoutRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LogoutResponse)
		return ctx.Result(200, reply)
	}
}

type UserServiceHTTPClient interface {
	GetHealthRecord(ctx context.Context, req *GetHealthRecordRequest, opts ...http.CallOption) (rsp *HealthRecordResponse, err error)
	GetUserPreferences(ctx context.Context, req *GetUserPreferencesRequest, opts ...http.CallOption) (rsp *UserPreferencesResponse, err error)
	GetUserProfile(ctx context.Context, req *GetUserProfileRequest, opts ...http.CallOption) (rsp *UserProfileResponse, err error)
	LoginUser(ctx context.Context, req *LoginRequest, opts ...http.CallOption) (rsp *LoginResponse, err error)
	LogoutUser(ctx context.Context, req *LogoutRequest, opts ...http.CallOption) (rsp *LogoutResponse, err error)
	RegisterUser(ctx context.Context, req *RegisterRequest, opts ...http.CallOption) (rsp *RegisterResponse, err error)
	UpdateHealthRecord(ctx context.Context, req *UpdateHealthRecordRequest, opts ...http.CallOption) (rsp *HealthRecordResponse, err error)
	UpdateUserPreferences(ctx context.Context, req *UpdateUserPreferencesRequest, opts ...http.CallOption) (rsp *UserPreferencesResponse, err error)
	UpdateUserProfile(ctx context.Context, req *UpdateUserProfileRequest, opts ...http.CallOption) (rsp *UserProfileResponse, err error)
}

type UserServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewUserServiceHTTPClient(client *http.Client) UserServiceHTTPClient {
	return &UserServiceHTTPClientImpl{client}
}

func (c *UserServiceHTTPClientImpl) GetHealthRecord(ctx context.Context, in *GetHealthRecordRequest, opts ...http.CallOption) (*HealthRecordResponse, error) {
	var out HealthRecordResponse
	pattern := "/users/health"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationUserServiceGetHealthRecord))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *UserServiceHTTPClientImpl) GetUserPreferences(ctx context.Context, in *GetUserPreferencesRequest, opts ...http.CallOption) (*UserPreferencesResponse, error) {
	var out UserPreferencesResponse
	pattern := "/users/preferences"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationUserServiceGetUserPreferences))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *UserServiceHTTPClientImpl) GetUserProfile(ctx context.Context, in *GetUserProfileRequest, opts ...http.CallOption) (*UserProfileResponse, error) {
	var out UserProfileResponse
	pattern := "/users/profile"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationUserServiceGetUserProfile))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *UserServiceHTTPClientImpl) LoginUser(ctx context.Context, in *LoginRequest, opts ...http.CallOption) (*LoginResponse, error) {
	var out LoginResponse
	pattern := "/users/login"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserServiceLoginUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *UserServiceHTTPClientImpl) LogoutUser(ctx context.Context, in *LogoutRequest, opts ...http.CallOption) (*LogoutResponse, error) {
	var out LogoutResponse
	pattern := "/users/logout"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserServiceLogoutUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *UserServiceHTTPClientImpl) RegisterUser(ctx context.Context, in *RegisterRequest, opts ...http.CallOption) (*RegisterResponse, error) {
	var out RegisterResponse
	pattern := "/users/register"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserServiceRegisterUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *UserServiceHTTPClientImpl) UpdateHealthRecord(ctx context.Context, in *UpdateHealthRecordRequest, opts ...http.CallOption) (*HealthRecordResponse, error) {
	var out HealthRecordResponse
	pattern := "/users/health"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserServiceUpdateHealthRecord))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *UserServiceHTTPClientImpl) UpdateUserPreferences(ctx context.Context, in *UpdateUserPreferencesRequest, opts ...http.CallOption) (*UserPreferencesResponse, error) {
	var out UserPreferencesResponse
	pattern := "/users/preferences"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserServiceUpdateUserPreferences))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *UserServiceHTTPClientImpl) UpdateUserProfile(ctx context.Context, in *UpdateUserProfileRequest, opts ...http.CallOption) (*UserProfileResponse, error) {
	var out UserProfileResponse
	pattern := "/users/profile"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserServiceUpdateUserProfile))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
