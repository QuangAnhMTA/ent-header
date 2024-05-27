// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: config.proto

package config

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ConfigServiceClient is the client API for ConfigService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConfigServiceClient interface {
	CheckRegisterAccount(ctx context.Context, in *CheckRegisterAccountRequest, opts ...grpc.CallOption) (*CheckRegisterAccountResponse, error)
	VerifyOtpRegister(ctx context.Context, in *VerifyOtpRegisterRequest, opts ...grpc.CallOption) (*VerifyOtpRegisterResponse, error)
	RegisterAccount(ctx context.Context, in *RegisterAccountRequest, opts ...grpc.CallOption) (*Account, error)
	Login(ctx context.Context, in *LoginAccountRequest, opts ...grpc.CallOption) (*LoginAccountResponse, error)
	MemberLogin(ctx context.Context, in *MemberRequest, opts ...grpc.CallOption) (*LoginMemberResponse, error)
	ChangeMemberPassword(ctx context.Context, in *Member, opts ...grpc.CallOption) (*Member, error)
	ListAccounts(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*Accounts, error)
	UpdateAccount(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*Account, error)
	// -----------------member------------------------
	ListMembers(ctx context.Context, in *MemberRequest, opts ...grpc.CallOption) (*Members, error)
	//
	CreateMember(ctx context.Context, in *Member, opts ...grpc.CallOption) (*Member, error)
	GetMember(ctx context.Context, in *Member, opts ...grpc.CallOption) (*Member, error)
	// ------------------category----------------------
	ListCategories(ctx context.Context, in *CategoryRequest, opts ...grpc.CallOption) (*Categories, error)
	CreateMemberCategory(ctx context.Context, in *MemberCategory, opts ...grpc.CallOption) (*MemberCategory, error)
	GetMemberCategory(ctx context.Context, in *MemberCategory, opts ...grpc.CallOption) (*MemberCategory, error)
	ListMemberCategories(ctx context.Context, in *MemberCategoryRequest, opts ...grpc.CallOption) (*MemberCategories, error)
}

type configServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewConfigServiceClient(cc grpc.ClientConnInterface) ConfigServiceClient {
	return &configServiceClient{cc}
}

func (c *configServiceClient) CheckRegisterAccount(ctx context.Context, in *CheckRegisterAccountRequest, opts ...grpc.CallOption) (*CheckRegisterAccountResponse, error) {
	out := new(CheckRegisterAccountResponse)
	err := c.cc.Invoke(ctx, "/config.ConfigService/CheckRegisterAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configServiceClient) VerifyOtpRegister(ctx context.Context, in *VerifyOtpRegisterRequest, opts ...grpc.CallOption) (*VerifyOtpRegisterResponse, error) {
	out := new(VerifyOtpRegisterResponse)
	err := c.cc.Invoke(ctx, "/config.ConfigService/VerifyOtpRegister", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configServiceClient) RegisterAccount(ctx context.Context, in *RegisterAccountRequest, opts ...grpc.CallOption) (*Account, error) {
	out := new(Account)
	err := c.cc.Invoke(ctx, "/config.ConfigService/RegisterAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configServiceClient) Login(ctx context.Context, in *LoginAccountRequest, opts ...grpc.CallOption) (*LoginAccountResponse, error) {
	out := new(LoginAccountResponse)
	err := c.cc.Invoke(ctx, "/config.ConfigService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configServiceClient) MemberLogin(ctx context.Context, in *MemberRequest, opts ...grpc.CallOption) (*LoginMemberResponse, error) {
	out := new(LoginMemberResponse)
	err := c.cc.Invoke(ctx, "/config.ConfigService/MemberLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configServiceClient) ChangeMemberPassword(ctx context.Context, in *Member, opts ...grpc.CallOption) (*Member, error) {
	out := new(Member)
	err := c.cc.Invoke(ctx, "/config.ConfigService/ChangeMemberPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configServiceClient) ListAccounts(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*Accounts, error) {
	out := new(Accounts)
	err := c.cc.Invoke(ctx, "/config.ConfigService/ListAccounts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configServiceClient) UpdateAccount(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*Account, error) {
	out := new(Account)
	err := c.cc.Invoke(ctx, "/config.ConfigService/UpdateAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configServiceClient) ListMembers(ctx context.Context, in *MemberRequest, opts ...grpc.CallOption) (*Members, error) {
	out := new(Members)
	err := c.cc.Invoke(ctx, "/config.ConfigService/ListMembers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configServiceClient) CreateMember(ctx context.Context, in *Member, opts ...grpc.CallOption) (*Member, error) {
	out := new(Member)
	err := c.cc.Invoke(ctx, "/config.ConfigService/CreateMember", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configServiceClient) GetMember(ctx context.Context, in *Member, opts ...grpc.CallOption) (*Member, error) {
	out := new(Member)
	err := c.cc.Invoke(ctx, "/config.ConfigService/GetMember", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configServiceClient) ListCategories(ctx context.Context, in *CategoryRequest, opts ...grpc.CallOption) (*Categories, error) {
	out := new(Categories)
	err := c.cc.Invoke(ctx, "/config.ConfigService/ListCategories", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configServiceClient) CreateMemberCategory(ctx context.Context, in *MemberCategory, opts ...grpc.CallOption) (*MemberCategory, error) {
	out := new(MemberCategory)
	err := c.cc.Invoke(ctx, "/config.ConfigService/CreateMemberCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configServiceClient) GetMemberCategory(ctx context.Context, in *MemberCategory, opts ...grpc.CallOption) (*MemberCategory, error) {
	out := new(MemberCategory)
	err := c.cc.Invoke(ctx, "/config.ConfigService/GetMemberCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configServiceClient) ListMemberCategories(ctx context.Context, in *MemberCategoryRequest, opts ...grpc.CallOption) (*MemberCategories, error) {
	out := new(MemberCategories)
	err := c.cc.Invoke(ctx, "/config.ConfigService/ListMemberCategories", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConfigServiceServer is the server API for ConfigService service.
// All implementations should embed UnimplementedConfigServiceServer
// for forward compatibility
type ConfigServiceServer interface {
	CheckRegisterAccount(context.Context, *CheckRegisterAccountRequest) (*CheckRegisterAccountResponse, error)
	VerifyOtpRegister(context.Context, *VerifyOtpRegisterRequest) (*VerifyOtpRegisterResponse, error)
	RegisterAccount(context.Context, *RegisterAccountRequest) (*Account, error)
	Login(context.Context, *LoginAccountRequest) (*LoginAccountResponse, error)
	MemberLogin(context.Context, *MemberRequest) (*LoginMemberResponse, error)
	ChangeMemberPassword(context.Context, *Member) (*Member, error)
	ListAccounts(context.Context, *AccountRequest) (*Accounts, error)
	UpdateAccount(context.Context, *AccountRequest) (*Account, error)
	// -----------------member------------------------
	ListMembers(context.Context, *MemberRequest) (*Members, error)
	//
	CreateMember(context.Context, *Member) (*Member, error)
	GetMember(context.Context, *Member) (*Member, error)
	// ------------------category----------------------
	ListCategories(context.Context, *CategoryRequest) (*Categories, error)
	CreateMemberCategory(context.Context, *MemberCategory) (*MemberCategory, error)
	GetMemberCategory(context.Context, *MemberCategory) (*MemberCategory, error)
	ListMemberCategories(context.Context, *MemberCategoryRequest) (*MemberCategories, error)
}

// UnimplementedConfigServiceServer should be embedded to have forward compatible implementations.
type UnimplementedConfigServiceServer struct {
}

func (UnimplementedConfigServiceServer) CheckRegisterAccount(context.Context, *CheckRegisterAccountRequest) (*CheckRegisterAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckRegisterAccount not implemented")
}
func (UnimplementedConfigServiceServer) VerifyOtpRegister(context.Context, *VerifyOtpRegisterRequest) (*VerifyOtpRegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyOtpRegister not implemented")
}
func (UnimplementedConfigServiceServer) RegisterAccount(context.Context, *RegisterAccountRequest) (*Account, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterAccount not implemented")
}
func (UnimplementedConfigServiceServer) Login(context.Context, *LoginAccountRequest) (*LoginAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedConfigServiceServer) MemberLogin(context.Context, *MemberRequest) (*LoginMemberResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MemberLogin not implemented")
}
func (UnimplementedConfigServiceServer) ChangeMemberPassword(context.Context, *Member) (*Member, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeMemberPassword not implemented")
}
func (UnimplementedConfigServiceServer) ListAccounts(context.Context, *AccountRequest) (*Accounts, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAccounts not implemented")
}
func (UnimplementedConfigServiceServer) UpdateAccount(context.Context, *AccountRequest) (*Account, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAccount not implemented")
}
func (UnimplementedConfigServiceServer) ListMembers(context.Context, *MemberRequest) (*Members, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMembers not implemented")
}
func (UnimplementedConfigServiceServer) CreateMember(context.Context, *Member) (*Member, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMember not implemented")
}
func (UnimplementedConfigServiceServer) GetMember(context.Context, *Member) (*Member, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMember not implemented")
}
func (UnimplementedConfigServiceServer) ListCategories(context.Context, *CategoryRequest) (*Categories, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCategories not implemented")
}
func (UnimplementedConfigServiceServer) CreateMemberCategory(context.Context, *MemberCategory) (*MemberCategory, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMemberCategory not implemented")
}
func (UnimplementedConfigServiceServer) GetMemberCategory(context.Context, *MemberCategory) (*MemberCategory, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMemberCategory not implemented")
}
func (UnimplementedConfigServiceServer) ListMemberCategories(context.Context, *MemberCategoryRequest) (*MemberCategories, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMemberCategories not implemented")
}

// UnsafeConfigServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConfigServiceServer will
// result in compilation errors.
type UnsafeConfigServiceServer interface {
	mustEmbedUnimplementedConfigServiceServer()
}

func RegisterConfigServiceServer(s grpc.ServiceRegistrar, srv ConfigServiceServer) {
	s.RegisterService(&ConfigService_ServiceDesc, srv)
}

func _ConfigService_CheckRegisterAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckRegisterAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServiceServer).CheckRegisterAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/config.ConfigService/CheckRegisterAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServiceServer).CheckRegisterAccount(ctx, req.(*CheckRegisterAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfigService_VerifyOtpRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyOtpRegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServiceServer).VerifyOtpRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/config.ConfigService/VerifyOtpRegister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServiceServer).VerifyOtpRegister(ctx, req.(*VerifyOtpRegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfigService_RegisterAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServiceServer).RegisterAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/config.ConfigService/RegisterAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServiceServer).RegisterAccount(ctx, req.(*RegisterAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfigService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/config.ConfigService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServiceServer).Login(ctx, req.(*LoginAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfigService_MemberLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MemberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServiceServer).MemberLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/config.ConfigService/MemberLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServiceServer).MemberLogin(ctx, req.(*MemberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfigService_ChangeMemberPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Member)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServiceServer).ChangeMemberPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/config.ConfigService/ChangeMemberPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServiceServer).ChangeMemberPassword(ctx, req.(*Member))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfigService_ListAccounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServiceServer).ListAccounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/config.ConfigService/ListAccounts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServiceServer).ListAccounts(ctx, req.(*AccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfigService_UpdateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServiceServer).UpdateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/config.ConfigService/UpdateAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServiceServer).UpdateAccount(ctx, req.(*AccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfigService_ListMembers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MemberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServiceServer).ListMembers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/config.ConfigService/ListMembers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServiceServer).ListMembers(ctx, req.(*MemberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfigService_CreateMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Member)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServiceServer).CreateMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/config.ConfigService/CreateMember",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServiceServer).CreateMember(ctx, req.(*Member))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfigService_GetMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Member)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServiceServer).GetMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/config.ConfigService/GetMember",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServiceServer).GetMember(ctx, req.(*Member))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfigService_ListCategories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServiceServer).ListCategories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/config.ConfigService/ListCategories",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServiceServer).ListCategories(ctx, req.(*CategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfigService_CreateMemberCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MemberCategory)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServiceServer).CreateMemberCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/config.ConfigService/CreateMemberCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServiceServer).CreateMemberCategory(ctx, req.(*MemberCategory))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfigService_GetMemberCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MemberCategory)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServiceServer).GetMemberCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/config.ConfigService/GetMemberCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServiceServer).GetMemberCategory(ctx, req.(*MemberCategory))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfigService_ListMemberCategories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MemberCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServiceServer).ListMemberCategories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/config.ConfigService/ListMemberCategories",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServiceServer).ListMemberCategories(ctx, req.(*MemberCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ConfigService_ServiceDesc is the grpc.ServiceDesc for ConfigService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ConfigService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "config.ConfigService",
	HandlerType: (*ConfigServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckRegisterAccount",
			Handler:    _ConfigService_CheckRegisterAccount_Handler,
		},
		{
			MethodName: "VerifyOtpRegister",
			Handler:    _ConfigService_VerifyOtpRegister_Handler,
		},
		{
			MethodName: "RegisterAccount",
			Handler:    _ConfigService_RegisterAccount_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _ConfigService_Login_Handler,
		},
		{
			MethodName: "MemberLogin",
			Handler:    _ConfigService_MemberLogin_Handler,
		},
		{
			MethodName: "ChangeMemberPassword",
			Handler:    _ConfigService_ChangeMemberPassword_Handler,
		},
		{
			MethodName: "ListAccounts",
			Handler:    _ConfigService_ListAccounts_Handler,
		},
		{
			MethodName: "UpdateAccount",
			Handler:    _ConfigService_UpdateAccount_Handler,
		},
		{
			MethodName: "ListMembers",
			Handler:    _ConfigService_ListMembers_Handler,
		},
		{
			MethodName: "CreateMember",
			Handler:    _ConfigService_CreateMember_Handler,
		},
		{
			MethodName: "GetMember",
			Handler:    _ConfigService_GetMember_Handler,
		},
		{
			MethodName: "ListCategories",
			Handler:    _ConfigService_ListCategories_Handler,
		},
		{
			MethodName: "CreateMemberCategory",
			Handler:    _ConfigService_CreateMemberCategory_Handler,
		},
		{
			MethodName: "GetMemberCategory",
			Handler:    _ConfigService_GetMemberCategory_Handler,
		},
		{
			MethodName: "ListMemberCategories",
			Handler:    _ConfigService_ListMemberCategories_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "config.proto",
}
