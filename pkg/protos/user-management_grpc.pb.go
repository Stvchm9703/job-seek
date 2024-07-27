// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.1
// source: user-management.proto

package protos

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	UserManagementService_SaveUserAccount_FullMethodName                        = "/job_seek.user_management.UserManagementService/SaveUserAccount"
	UserManagementService_GetUserAccount_FullMethodName                         = "/job_seek.user_management.UserManagementService/GetUserAccount"
	UserManagementService_SaveUserSearchPreference_FullMethodName               = "/job_seek.user_management.UserManagementService/SaveUserSearchPreference"
	UserManagementService_UpdateUserSearchPreference_FullMethodName             = "/job_seek.user_management.UserManagementService/UpdateUserSearchPreference"
	UserManagementService_GetUserSearchPreference_FullMethodName                = "/job_seek.user_management.UserManagementService/GetUserSearchPreference"
	UserManagementService_SaveUserJobSearchPredictedPreference_FullMethodName   = "/job_seek.user_management.UserManagementService/SaveUserJobSearchPredictedPreference"
	UserManagementService_GetUserJobSearchPredictedPreference_FullMethodName    = "/job_seek.user_management.UserManagementService/GetUserJobSearchPredictedPreference"
	UserManagementService_CreateUserJobSearchPredictedPreference_FullMethodName = "/job_seek.user_management.UserManagementService/CreateUserJobSearchPredictedPreference"
)

// UserManagementServiceClient is the client API for UserManagementService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserManagementServiceClient interface {
	// from user management
	SaveUserAccount(ctx context.Context, in *UserAccount, opts ...grpc.CallOption) (*UserResponse, error)
	GetUserAccount(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*UserAccount, error)
	// store point of user search Preference
	SaveUserSearchPreference(ctx context.Context, in *UserSearchPreference, opts ...grpc.CallOption) (*UserResponse, error)
	UpdateUserSearchPreference(ctx context.Context, in *UserSearchPreference, opts ...grpc.CallOption) (*UserResponse, error)
	GetUserSearchPreference(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*UserSearchPreference, error)
	SaveUserJobSearchPredictedPreference(ctx context.Context, in *UserJobSearchPredictedPreference, opts ...grpc.CallOption) (*PredictedPreferenceResponse, error)
	GetUserJobSearchPredictedPreference(ctx context.Context, in *GetPredictedPreference, opts ...grpc.CallOption) (*UserJobSearchPredictedPreference, error)
	CreateUserJobSearchPredictedPreference(ctx context.Context, opts ...grpc.CallOption) (UserManagementService_CreateUserJobSearchPredictedPreferenceClient, error)
}

type userManagementServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserManagementServiceClient(cc grpc.ClientConnInterface) UserManagementServiceClient {
	return &userManagementServiceClient{cc}
}

func (c *userManagementServiceClient) SaveUserAccount(ctx context.Context, in *UserAccount, opts ...grpc.CallOption) (*UserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, UserManagementService_SaveUserAccount_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userManagementServiceClient) GetUserAccount(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*UserAccount, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserAccount)
	err := c.cc.Invoke(ctx, UserManagementService_GetUserAccount_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userManagementServiceClient) SaveUserSearchPreference(ctx context.Context, in *UserSearchPreference, opts ...grpc.CallOption) (*UserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, UserManagementService_SaveUserSearchPreference_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userManagementServiceClient) UpdateUserSearchPreference(ctx context.Context, in *UserSearchPreference, opts ...grpc.CallOption) (*UserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, UserManagementService_UpdateUserSearchPreference_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userManagementServiceClient) GetUserSearchPreference(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*UserSearchPreference, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserSearchPreference)
	err := c.cc.Invoke(ctx, UserManagementService_GetUserSearchPreference_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userManagementServiceClient) SaveUserJobSearchPredictedPreference(ctx context.Context, in *UserJobSearchPredictedPreference, opts ...grpc.CallOption) (*PredictedPreferenceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PredictedPreferenceResponse)
	err := c.cc.Invoke(ctx, UserManagementService_SaveUserJobSearchPredictedPreference_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userManagementServiceClient) GetUserJobSearchPredictedPreference(ctx context.Context, in *GetPredictedPreference, opts ...grpc.CallOption) (*UserJobSearchPredictedPreference, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserJobSearchPredictedPreference)
	err := c.cc.Invoke(ctx, UserManagementService_GetUserJobSearchPredictedPreference_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userManagementServiceClient) CreateUserJobSearchPredictedPreference(ctx context.Context, opts ...grpc.CallOption) (UserManagementService_CreateUserJobSearchPredictedPreferenceClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &UserManagementService_ServiceDesc.Streams[0], UserManagementService_CreateUserJobSearchPredictedPreference_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &userManagementServiceCreateUserJobSearchPredictedPreferenceClient{ClientStream: stream}
	return x, nil
}

type UserManagementService_CreateUserJobSearchPredictedPreferenceClient interface {
	Send(*UserJobSearchPredictedPreference) error
	Recv() (*PredictedPreferenceResponse, error)
	grpc.ClientStream
}

type userManagementServiceCreateUserJobSearchPredictedPreferenceClient struct {
	grpc.ClientStream
}

func (x *userManagementServiceCreateUserJobSearchPredictedPreferenceClient) Send(m *UserJobSearchPredictedPreference) error {
	return x.ClientStream.SendMsg(m)
}

func (x *userManagementServiceCreateUserJobSearchPredictedPreferenceClient) Recv() (*PredictedPreferenceResponse, error) {
	m := new(PredictedPreferenceResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// UserManagementServiceServer is the server API for UserManagementService service.
// All implementations must embed UnimplementedUserManagementServiceServer
// for forward compatibility
type UserManagementServiceServer interface {
	// from user management
	SaveUserAccount(context.Context, *UserAccount) (*UserResponse, error)
	GetUserAccount(context.Context, *GetUserRequest) (*UserAccount, error)
	// store point of user search Preference
	SaveUserSearchPreference(context.Context, *UserSearchPreference) (*UserResponse, error)
	UpdateUserSearchPreference(context.Context, *UserSearchPreference) (*UserResponse, error)
	GetUserSearchPreference(context.Context, *GetUserRequest) (*UserSearchPreference, error)
	SaveUserJobSearchPredictedPreference(context.Context, *UserJobSearchPredictedPreference) (*PredictedPreferenceResponse, error)
	GetUserJobSearchPredictedPreference(context.Context, *GetPredictedPreference) (*UserJobSearchPredictedPreference, error)
	CreateUserJobSearchPredictedPreference(UserManagementService_CreateUserJobSearchPredictedPreferenceServer) error
	mustEmbedUnimplementedUserManagementServiceServer()
}

// UnimplementedUserManagementServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserManagementServiceServer struct {
}

func (UnimplementedUserManagementServiceServer) SaveUserAccount(context.Context, *UserAccount) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveUserAccount not implemented")
}
func (UnimplementedUserManagementServiceServer) GetUserAccount(context.Context, *GetUserRequest) (*UserAccount, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserAccount not implemented")
}
func (UnimplementedUserManagementServiceServer) SaveUserSearchPreference(context.Context, *UserSearchPreference) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveUserSearchPreference not implemented")
}
func (UnimplementedUserManagementServiceServer) UpdateUserSearchPreference(context.Context, *UserSearchPreference) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserSearchPreference not implemented")
}
func (UnimplementedUserManagementServiceServer) GetUserSearchPreference(context.Context, *GetUserRequest) (*UserSearchPreference, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserSearchPreference not implemented")
}
func (UnimplementedUserManagementServiceServer) SaveUserJobSearchPredictedPreference(context.Context, *UserJobSearchPredictedPreference) (*PredictedPreferenceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveUserJobSearchPredictedPreference not implemented")
}
func (UnimplementedUserManagementServiceServer) GetUserJobSearchPredictedPreference(context.Context, *GetPredictedPreference) (*UserJobSearchPredictedPreference, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserJobSearchPredictedPreference not implemented")
}
func (UnimplementedUserManagementServiceServer) CreateUserJobSearchPredictedPreference(UserManagementService_CreateUserJobSearchPredictedPreferenceServer) error {
	return status.Errorf(codes.Unimplemented, "method CreateUserJobSearchPredictedPreference not implemented")
}
func (UnimplementedUserManagementServiceServer) mustEmbedUnimplementedUserManagementServiceServer() {}

// UnsafeUserManagementServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserManagementServiceServer will
// result in compilation errors.
type UnsafeUserManagementServiceServer interface {
	mustEmbedUnimplementedUserManagementServiceServer()
}

func RegisterUserManagementServiceServer(s grpc.ServiceRegistrar, srv UserManagementServiceServer) {
	s.RegisterService(&UserManagementService_ServiceDesc, srv)
}

func _UserManagementService_SaveUserAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserAccount)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserManagementServiceServer).SaveUserAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserManagementService_SaveUserAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserManagementServiceServer).SaveUserAccount(ctx, req.(*UserAccount))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserManagementService_GetUserAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserManagementServiceServer).GetUserAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserManagementService_GetUserAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserManagementServiceServer).GetUserAccount(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserManagementService_SaveUserSearchPreference_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserSearchPreference)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserManagementServiceServer).SaveUserSearchPreference(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserManagementService_SaveUserSearchPreference_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserManagementServiceServer).SaveUserSearchPreference(ctx, req.(*UserSearchPreference))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserManagementService_UpdateUserSearchPreference_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserSearchPreference)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserManagementServiceServer).UpdateUserSearchPreference(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserManagementService_UpdateUserSearchPreference_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserManagementServiceServer).UpdateUserSearchPreference(ctx, req.(*UserSearchPreference))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserManagementService_GetUserSearchPreference_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserManagementServiceServer).GetUserSearchPreference(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserManagementService_GetUserSearchPreference_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserManagementServiceServer).GetUserSearchPreference(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserManagementService_SaveUserJobSearchPredictedPreference_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserJobSearchPredictedPreference)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserManagementServiceServer).SaveUserJobSearchPredictedPreference(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserManagementService_SaveUserJobSearchPredictedPreference_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserManagementServiceServer).SaveUserJobSearchPredictedPreference(ctx, req.(*UserJobSearchPredictedPreference))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserManagementService_GetUserJobSearchPredictedPreference_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPredictedPreference)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserManagementServiceServer).GetUserJobSearchPredictedPreference(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserManagementService_GetUserJobSearchPredictedPreference_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserManagementServiceServer).GetUserJobSearchPredictedPreference(ctx, req.(*GetPredictedPreference))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserManagementService_CreateUserJobSearchPredictedPreference_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(UserManagementServiceServer).CreateUserJobSearchPredictedPreference(&userManagementServiceCreateUserJobSearchPredictedPreferenceServer{ServerStream: stream})
}

type UserManagementService_CreateUserJobSearchPredictedPreferenceServer interface {
	Send(*PredictedPreferenceResponse) error
	Recv() (*UserJobSearchPredictedPreference, error)
	grpc.ServerStream
}

type userManagementServiceCreateUserJobSearchPredictedPreferenceServer struct {
	grpc.ServerStream
}

func (x *userManagementServiceCreateUserJobSearchPredictedPreferenceServer) Send(m *PredictedPreferenceResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *userManagementServiceCreateUserJobSearchPredictedPreferenceServer) Recv() (*UserJobSearchPredictedPreference, error) {
	m := new(UserJobSearchPredictedPreference)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// UserManagementService_ServiceDesc is the grpc.ServiceDesc for UserManagementService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserManagementService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "job_seek.user_management.UserManagementService",
	HandlerType: (*UserManagementServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SaveUserAccount",
			Handler:    _UserManagementService_SaveUserAccount_Handler,
		},
		{
			MethodName: "GetUserAccount",
			Handler:    _UserManagementService_GetUserAccount_Handler,
		},
		{
			MethodName: "SaveUserSearchPreference",
			Handler:    _UserManagementService_SaveUserSearchPreference_Handler,
		},
		{
			MethodName: "UpdateUserSearchPreference",
			Handler:    _UserManagementService_UpdateUserSearchPreference_Handler,
		},
		{
			MethodName: "GetUserSearchPreference",
			Handler:    _UserManagementService_GetUserSearchPreference_Handler,
		},
		{
			MethodName: "SaveUserJobSearchPredictedPreference",
			Handler:    _UserManagementService_SaveUserJobSearchPredictedPreference_Handler,
		},
		{
			MethodName: "GetUserJobSearchPredictedPreference",
			Handler:    _UserManagementService_GetUserJobSearchPredictedPreference_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "CreateUserJobSearchPredictedPreference",
			Handler:       _UserManagementService_CreateUserJobSearchPredictedPreference_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "user-management.proto",
}