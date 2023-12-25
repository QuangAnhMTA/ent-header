// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: transaction.proto

package transaction

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

// TransactionServiceClient is the client API for TransactionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TransactionServiceClient interface {
	// ---------------- Course -----------------
	ListListen(ctx context.Context, in *ListenRequest, opts ...grpc.CallOption) (*Listen, error)
	EndLearnListen(ctx context.Context, in *Listen, opts ...grpc.CallOption) (*Listen, error)
	ListSpeak(ctx context.Context, in *SpeakRequest, opts ...grpc.CallOption) (*Speaks, error)
	EndLearnSpeak(ctx context.Context, in *Speak, opts ...grpc.CallOption) (*Speak, error)
	CreateSentence(ctx context.Context, in *Sentence, opts ...grpc.CallOption) (*Sentence, error)
	CreateLookup(ctx context.Context, in *Lookup, opts ...grpc.CallOption) (*Lookup, error)
	CreateParagraph(ctx context.Context, in *Paragraph, opts ...grpc.CallOption) (*Paragraph, error)
}

type transactionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTransactionServiceClient(cc grpc.ClientConnInterface) TransactionServiceClient {
	return &transactionServiceClient{cc}
}

func (c *transactionServiceClient) ListListen(ctx context.Context, in *ListenRequest, opts ...grpc.CallOption) (*Listen, error) {
	out := new(Listen)
	err := c.cc.Invoke(ctx, "/transaction.TransactionService/ListListen", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) EndLearnListen(ctx context.Context, in *Listen, opts ...grpc.CallOption) (*Listen, error) {
	out := new(Listen)
	err := c.cc.Invoke(ctx, "/transaction.TransactionService/EndLearnListen", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) ListSpeak(ctx context.Context, in *SpeakRequest, opts ...grpc.CallOption) (*Speaks, error) {
	out := new(Speaks)
	err := c.cc.Invoke(ctx, "/transaction.TransactionService/ListSpeak", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) EndLearnSpeak(ctx context.Context, in *Speak, opts ...grpc.CallOption) (*Speak, error) {
	out := new(Speak)
	err := c.cc.Invoke(ctx, "/transaction.TransactionService/EndLearnSpeak", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) CreateSentence(ctx context.Context, in *Sentence, opts ...grpc.CallOption) (*Sentence, error) {
	out := new(Sentence)
	err := c.cc.Invoke(ctx, "/transaction.TransactionService/CreateSentence", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) CreateLookup(ctx context.Context, in *Lookup, opts ...grpc.CallOption) (*Lookup, error) {
	out := new(Lookup)
	err := c.cc.Invoke(ctx, "/transaction.TransactionService/CreateLookup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) CreateParagraph(ctx context.Context, in *Paragraph, opts ...grpc.CallOption) (*Paragraph, error) {
	out := new(Paragraph)
	err := c.cc.Invoke(ctx, "/transaction.TransactionService/CreateParagraph", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransactionServiceServer is the server API for TransactionService service.
// All implementations should embed UnimplementedTransactionServiceServer
// for forward compatibility
type TransactionServiceServer interface {
	// ---------------- Course -----------------
	ListListen(context.Context, *ListenRequest) (*Listen, error)
	EndLearnListen(context.Context, *Listen) (*Listen, error)
	ListSpeak(context.Context, *SpeakRequest) (*Speaks, error)
	EndLearnSpeak(context.Context, *Speak) (*Speak, error)
	CreateSentence(context.Context, *Sentence) (*Sentence, error)
	CreateLookup(context.Context, *Lookup) (*Lookup, error)
	CreateParagraph(context.Context, *Paragraph) (*Paragraph, error)
}

// UnimplementedTransactionServiceServer should be embedded to have forward compatible implementations.
type UnimplementedTransactionServiceServer struct {
}

func (UnimplementedTransactionServiceServer) ListListen(context.Context, *ListenRequest) (*Listen, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListListen not implemented")
}
func (UnimplementedTransactionServiceServer) EndLearnListen(context.Context, *Listen) (*Listen, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EndLearnListen not implemented")
}
func (UnimplementedTransactionServiceServer) ListSpeak(context.Context, *SpeakRequest) (*Speaks, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSpeak not implemented")
}
func (UnimplementedTransactionServiceServer) EndLearnSpeak(context.Context, *Speak) (*Speak, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EndLearnSpeak not implemented")
}
func (UnimplementedTransactionServiceServer) CreateSentence(context.Context, *Sentence) (*Sentence, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSentence not implemented")
}
func (UnimplementedTransactionServiceServer) CreateLookup(context.Context, *Lookup) (*Lookup, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLookup not implemented")
}
func (UnimplementedTransactionServiceServer) CreateParagraph(context.Context, *Paragraph) (*Paragraph, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateParagraph not implemented")
}

// UnsafeTransactionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TransactionServiceServer will
// result in compilation errors.
type UnsafeTransactionServiceServer interface {
	mustEmbedUnimplementedTransactionServiceServer()
}

func RegisterTransactionServiceServer(s grpc.ServiceRegistrar, srv TransactionServiceServer) {
	s.RegisterService(&TransactionService_ServiceDesc, srv)
}

func _TransactionService_ListListen_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).ListListen(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transaction.TransactionService/ListListen",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).ListListen(ctx, req.(*ListenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_EndLearnListen_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Listen)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).EndLearnListen(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transaction.TransactionService/EndLearnListen",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).EndLearnListen(ctx, req.(*Listen))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_ListSpeak_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SpeakRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).ListSpeak(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transaction.TransactionService/ListSpeak",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).ListSpeak(ctx, req.(*SpeakRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_EndLearnSpeak_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Speak)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).EndLearnSpeak(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transaction.TransactionService/EndLearnSpeak",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).EndLearnSpeak(ctx, req.(*Speak))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_CreateSentence_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Sentence)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).CreateSentence(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transaction.TransactionService/CreateSentence",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).CreateSentence(ctx, req.(*Sentence))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_CreateLookup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Lookup)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).CreateLookup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transaction.TransactionService/CreateLookup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).CreateLookup(ctx, req.(*Lookup))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_CreateParagraph_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Paragraph)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).CreateParagraph(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transaction.TransactionService/CreateParagraph",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).CreateParagraph(ctx, req.(*Paragraph))
	}
	return interceptor(ctx, in, info, handler)
}

// TransactionService_ServiceDesc is the grpc.ServiceDesc for TransactionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TransactionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "transaction.TransactionService",
	HandlerType: (*TransactionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListListen",
			Handler:    _TransactionService_ListListen_Handler,
		},
		{
			MethodName: "EndLearnListen",
			Handler:    _TransactionService_EndLearnListen_Handler,
		},
		{
			MethodName: "ListSpeak",
			Handler:    _TransactionService_ListSpeak_Handler,
		},
		{
			MethodName: "EndLearnSpeak",
			Handler:    _TransactionService_EndLearnSpeak_Handler,
		},
		{
			MethodName: "CreateSentence",
			Handler:    _TransactionService_CreateSentence_Handler,
		},
		{
			MethodName: "CreateLookup",
			Handler:    _TransactionService_CreateLookup_Handler,
		},
		{
			MethodName: "CreateParagraph",
			Handler:    _TransactionService_CreateParagraph_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "transaction.proto",
}
