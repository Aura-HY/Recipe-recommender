// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: api/user/v1/user.proto

package user

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	UserService_RegisterUser_FullMethodName          = "/user.UserService/RegisterUser"
	UserService_LoginUser_FullMethodName             = "/user.UserService/LoginUser"
	UserService_GetUserProfile_FullMethodName        = "/user.UserService/GetUserProfile"
	UserService_UpdateUserProfile_FullMethodName     = "/user.UserService/UpdateUserProfile"
	UserService_GetHealthRecord_FullMethodName       = "/user.UserService/GetHealthRecord"
	UserService_UpdateHealthRecord_FullMethodName    = "/user.UserService/UpdateHealthRecord"
	UserService_GetUserPreferences_FullMethodName    = "/user.UserService/GetUserPreferences"
	UserService_UpdateUserPreferences_FullMethodName = "/user.UserService/UpdateUserPreferences"
	UserService_LogoutUser_FullMethodName            = "/user.UserService/LogoutUser"
)

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	// 用户注册
	RegisterUser(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	// 用户登录
	LoginUser(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	// 获取用户个人信息
	GetUserProfile(ctx context.Context, in *GetUserProfileRequest, opts ...grpc.CallOption) (*UserProfileResponse, error)
	// 更新用户信息
	UpdateUserProfile(ctx context.Context, in *UpdateUserProfileRequest, opts ...grpc.CallOption) (*UserProfileResponse, error)
	// 获取用户健康档案
	GetHealthRecord(ctx context.Context, in *GetHealthRecordRequest, opts ...grpc.CallOption) (*HealthRecordResponse, error)
	// 更新用户健康档案（体质、过敏史等）
	UpdateHealthRecord(ctx context.Context, in *UpdateHealthRecordRequest, opts ...grpc.CallOption) (*HealthRecordResponse, error)
	// 获取用户饮食偏好
	GetUserPreferences(ctx context.Context, in *GetUserPreferencesRequest, opts ...grpc.CallOption) (*UserPreferencesResponse, error)
	// 更新用户饮食偏好
	UpdateUserPreferences(ctx context.Context, in *UpdateUserPreferencesRequest, opts ...grpc.CallOption) (*UserPreferencesResponse, error)
	// 退出登录
	LogoutUser(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutResponse, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) RegisterUser(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, UserService_RegisterUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) LoginUser(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, UserService_LoginUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserProfile(ctx context.Context, in *GetUserProfileRequest, opts ...grpc.CallOption) (*UserProfileResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserProfileResponse)
	err := c.cc.Invoke(ctx, UserService_GetUserProfile_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateUserProfile(ctx context.Context, in *UpdateUserProfileRequest, opts ...grpc.CallOption) (*UserProfileResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserProfileResponse)
	err := c.cc.Invoke(ctx, UserService_UpdateUserProfile_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetHealthRecord(ctx context.Context, in *GetHealthRecordRequest, opts ...grpc.CallOption) (*HealthRecordResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(HealthRecordResponse)
	err := c.cc.Invoke(ctx, UserService_GetHealthRecord_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateHealthRecord(ctx context.Context, in *UpdateHealthRecordRequest, opts ...grpc.CallOption) (*HealthRecordResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(HealthRecordResponse)
	err := c.cc.Invoke(ctx, UserService_UpdateHealthRecord_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserPreferences(ctx context.Context, in *GetUserPreferencesRequest, opts ...grpc.CallOption) (*UserPreferencesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserPreferencesResponse)
	err := c.cc.Invoke(ctx, UserService_GetUserPreferences_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateUserPreferences(ctx context.Context, in *UpdateUserPreferencesRequest, opts ...grpc.CallOption) (*UserPreferencesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserPreferencesResponse)
	err := c.cc.Invoke(ctx, UserService_UpdateUserPreferences_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) LogoutUser(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LogoutResponse)
	err := c.cc.Invoke(ctx, UserService_LogoutUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility.
type UserServiceServer interface {
	// 用户注册
	RegisterUser(context.Context, *RegisterRequest) (*RegisterResponse, error)
	// 用户登录
	LoginUser(context.Context, *LoginRequest) (*LoginResponse, error)
	// 获取用户个人信息
	GetUserProfile(context.Context, *GetUserProfileRequest) (*UserProfileResponse, error)
	// 更新用户信息
	UpdateUserProfile(context.Context, *UpdateUserProfileRequest) (*UserProfileResponse, error)
	// 获取用户健康档案
	GetHealthRecord(context.Context, *GetHealthRecordRequest) (*HealthRecordResponse, error)
	// 更新用户健康档案（体质、过敏史等）
	UpdateHealthRecord(context.Context, *UpdateHealthRecordRequest) (*HealthRecordResponse, error)
	// 获取用户饮食偏好
	GetUserPreferences(context.Context, *GetUserPreferencesRequest) (*UserPreferencesResponse, error)
	// 更新用户饮食偏好
	UpdateUserPreferences(context.Context, *UpdateUserPreferencesRequest) (*UserPreferencesResponse, error)
	// 退出登录
	LogoutUser(context.Context, *LogoutRequest) (*LogoutResponse, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedUserServiceServer struct{}

func (UnimplementedUserServiceServer) RegisterUser(context.Context, *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterUser not implemented")
}
func (UnimplementedUserServiceServer) LoginUser(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginUser not implemented")
}
func (UnimplementedUserServiceServer) GetUserProfile(context.Context, *GetUserProfileRequest) (*UserProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserProfile not implemented")
}
func (UnimplementedUserServiceServer) UpdateUserProfile(context.Context, *UpdateUserProfileRequest) (*UserProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserProfile not implemented")
}
func (UnimplementedUserServiceServer) GetHealthRecord(context.Context, *GetHealthRecordRequest) (*HealthRecordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHealthRecord not implemented")
}
func (UnimplementedUserServiceServer) UpdateHealthRecord(context.Context, *UpdateHealthRecordRequest) (*HealthRecordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateHealthRecord not implemented")
}
func (UnimplementedUserServiceServer) GetUserPreferences(context.Context, *GetUserPreferencesRequest) (*UserPreferencesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserPreferences not implemented")
}
func (UnimplementedUserServiceServer) UpdateUserPreferences(context.Context, *UpdateUserPreferencesRequest) (*UserPreferencesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserPreferences not implemented")
}
func (UnimplementedUserServiceServer) LogoutUser(context.Context, *LogoutRequest) (*LogoutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LogoutUser not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}
func (UnimplementedUserServiceServer) testEmbeddedByValue()                     {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	// If the following call pancis, it indicates UnimplementedUserServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_RegisterUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).RegisterUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_RegisterUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).RegisterUser(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_LoginUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).LoginUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_LoginUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).LoginUser(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetUserProfile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserProfile(ctx, req.(*GetUserProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateUserProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateUserProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_UpdateUserProfile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateUserProfile(ctx, req.(*UpdateUserProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetHealthRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHealthRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetHealthRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetHealthRecord_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetHealthRecord(ctx, req.(*GetHealthRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateHealthRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateHealthRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateHealthRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_UpdateHealthRecord_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateHealthRecord(ctx, req.(*UpdateHealthRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserPreferences_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserPreferencesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserPreferences(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetUserPreferences_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserPreferences(ctx, req.(*GetUserPreferencesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateUserPreferences_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserPreferencesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateUserPreferences(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_UpdateUserPreferences_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateUserPreferences(ctx, req.(*UpdateUserPreferencesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_LogoutUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).LogoutUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_LogoutUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).LogoutUser(ctx, req.(*LogoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterUser",
			Handler:    _UserService_RegisterUser_Handler,
		},
		{
			MethodName: "LoginUser",
			Handler:    _UserService_LoginUser_Handler,
		},
		{
			MethodName: "GetUserProfile",
			Handler:    _UserService_GetUserProfile_Handler,
		},
		{
			MethodName: "UpdateUserProfile",
			Handler:    _UserService_UpdateUserProfile_Handler,
		},
		{
			MethodName: "GetHealthRecord",
			Handler:    _UserService_GetHealthRecord_Handler,
		},
		{
			MethodName: "UpdateHealthRecord",
			Handler:    _UserService_UpdateHealthRecord_Handler,
		},
		{
			MethodName: "GetUserPreferences",
			Handler:    _UserService_GetUserPreferences_Handler,
		},
		{
			MethodName: "UpdateUserPreferences",
			Handler:    _UserService_UpdateUserPreferences_Handler,
		},
		{
			MethodName: "LogoutUser",
			Handler:    _UserService_LogoutUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/user/v1/user.proto",
}
