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
	CompleteParagraph(ctx context.Context, in *Paragraph, opts ...grpc.CallOption) (*Paragraph, error)
	ListLookup(ctx context.Context, in *LookupRequest, opts ...grpc.CallOption) (*Lookups, error)
	ListSentence(ctx context.Context, in *SentenceRequest, opts ...grpc.CallOption) (*Sentences, error)
	GetParagraphMax(ctx context.Context, in *ParagraphRequest, opts ...grpc.CallOption) (*Paragraph, error)
	ListParagraph(ctx context.Context, in *ParagraphRequest, opts ...grpc.CallOption) (*Paragraphs, error)
	ListListenDisplay(ctx context.Context, in *ListenDisplayRequest, opts ...grpc.CallOption) (*ListenDisplays, error)
	ListSpeakDisplay(ctx context.Context, in *SpeakDisplayRequest, opts ...grpc.CallOption) (*SpeakDisplays, error)
	ListFavourite(ctx context.Context, in *FavouriteRequest, opts ...grpc.CallOption) (*Favourites, error)
	CreateFavourite(ctx context.Context, in *Favourite, opts ...grpc.CallOption) (*Favourite, error)
	DeleteFavourite(ctx context.Context, in *Favourite, opts ...grpc.CallOption) (*Favourite, error)
	StartParagraph(ctx context.Context, in *Paragraph, opts ...grpc.CallOption) (*Paragraph, error)
	ListNewDetail(ctx context.Context, in *NewDetailRequest, opts ...grpc.CallOption) (*NewDetails, error)
	ListNew(ctx context.Context, in *NewRequest, opts ...grpc.CallOption) (*News, error)
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

func (c *transactionServiceClient) CompleteParagraph(ctx context.Context, in *Paragraph, opts ...grpc.CallOption) (*Paragraph, error) {
	out := new(Paragraph)
	err := c.cc.Invoke(ctx, "/transaction.TransactionService/CompleteParagraph", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) ListLookup(ctx context.Context, in *LookupRequest, opts ...grpc.CallOption) (*Lookups, error) {
	out := new(Lookups)
	err := c.cc.Invoke(ctx, "/transaction.TransactionService/ListLookup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) ListSentence(ctx context.Context, in *SentenceRequest, opts ...grpc.CallOption) (*Sentences, error) {
	out := new(Sentences)
	err := c.cc.Invoke(ctx, "/transaction.TransactionService/ListSentence", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) GetParagraphMax(ctx context.Context, in *ParagraphRequest, opts ...grpc.CallOption) (*Paragraph, error) {
	out := new(Paragraph)
	err := c.cc.Invoke(ctx, "/transaction.TransactionService/GetParagraphMax", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) ListParagraph(ctx context.Context, in *ParagraphRequest, opts ...grpc.CallOption) (*Paragraphs, error) {
	out := new(Paragraphs)
	err := c.cc.Invoke(ctx, "/transaction.TransactionService/ListParagraph", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) ListListenDisplay(ctx context.Context, in *ListenDisplayRequest, opts ...grpc.CallOption) (*ListenDisplays, error) {
	out := new(ListenDisplays)
	err := c.cc.Invoke(ctx, "/transaction.TransactionService/ListListenDisplay", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) ListSpeakDisplay(ctx context.Context, in *SpeakDisplayRequest, opts ...grpc.CallOption) (*SpeakDisplays, error) {
	out := new(SpeakDisplays)
	err := c.cc.Invoke(ctx, "/transaction.TransactionService/ListSpeakDisplay", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) ListFavourite(ctx context.Context, in *FavouriteRequest, opts ...grpc.CallOption) (*Favourites, error) {
	out := new(Favourites)
	err := c.cc.Invoke(ctx, "/transaction.TransactionService/ListFavourite", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) CreateFavourite(ctx context.Context, in *Favourite, opts ...grpc.CallOption) (*Favourite, error) {
	out := new(Favourite)
	err := c.cc.Invoke(ctx, "/transaction.TransactionService/CreateFavourite", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) DeleteFavourite(ctx context.Context, in *Favourite, opts ...grpc.CallOption) (*Favourite, error) {
	out := new(Favourite)
	err := c.cc.Invoke(ctx, "/transaction.TransactionService/DeleteFavourite", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) StartParagraph(ctx context.Context, in *Paragraph, opts ...grpc.CallOption) (*Paragraph, error) {
	out := new(Paragraph)
	err := c.cc.Invoke(ctx, "/transaction.TransactionService/StartParagraph", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) ListNewDetail(ctx context.Context, in *NewDetailRequest, opts ...grpc.CallOption) (*NewDetails, error) {
	out := new(NewDetails)
	err := c.cc.Invoke(ctx, "/transaction.TransactionService/ListNewDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) ListNew(ctx context.Context, in *NewRequest, opts ...grpc.CallOption) (*News, error) {
	out := new(News)
	err := c.cc.Invoke(ctx, "/transaction.TransactionService/ListNew", in, out, opts...)
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
	CompleteParagraph(context.Context, *Paragraph) (*Paragraph, error)
	ListLookup(context.Context, *LookupRequest) (*Lookups, error)
	ListSentence(context.Context, *SentenceRequest) (*Sentences, error)
	GetParagraphMax(context.Context, *ParagraphRequest) (*Paragraph, error)
	ListParagraph(context.Context, *ParagraphRequest) (*Paragraphs, error)
	ListListenDisplay(context.Context, *ListenDisplayRequest) (*ListenDisplays, error)
	ListSpeakDisplay(context.Context, *SpeakDisplayRequest) (*SpeakDisplays, error)
	ListFavourite(context.Context, *FavouriteRequest) (*Favourites, error)
	CreateFavourite(context.Context, *Favourite) (*Favourite, error)
	DeleteFavourite(context.Context, *Favourite) (*Favourite, error)
	StartParagraph(context.Context, *Paragraph) (*Paragraph, error)
	ListNewDetail(context.Context, *NewDetailRequest) (*NewDetails, error)
	ListNew(context.Context, *NewRequest) (*News, error)
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
func (UnimplementedTransactionServiceServer) CompleteParagraph(context.Context, *Paragraph) (*Paragraph, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CompleteParagraph not implemented")
}
func (UnimplementedTransactionServiceServer) ListLookup(context.Context, *LookupRequest) (*Lookups, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListLookup not implemented")
}
func (UnimplementedTransactionServiceServer) ListSentence(context.Context, *SentenceRequest) (*Sentences, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSentence not implemented")
}
func (UnimplementedTransactionServiceServer) GetParagraphMax(context.Context, *ParagraphRequest) (*Paragraph, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetParagraphMax not implemented")
}
func (UnimplementedTransactionServiceServer) ListParagraph(context.Context, *ParagraphRequest) (*Paragraphs, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListParagraph not implemented")
}
func (UnimplementedTransactionServiceServer) ListListenDisplay(context.Context, *ListenDisplayRequest) (*ListenDisplays, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListListenDisplay not implemented")
}
func (UnimplementedTransactionServiceServer) ListSpeakDisplay(context.Context, *SpeakDisplayRequest) (*SpeakDisplays, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSpeakDisplay not implemented")
}
func (UnimplementedTransactionServiceServer) ListFavourite(context.Context, *FavouriteRequest) (*Favourites, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFavourite not implemented")
}
func (UnimplementedTransactionServiceServer) CreateFavourite(context.Context, *Favourite) (*Favourite, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFavourite not implemented")
}
func (UnimplementedTransactionServiceServer) DeleteFavourite(context.Context, *Favourite) (*Favourite, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFavourite not implemented")
}
func (UnimplementedTransactionServiceServer) StartParagraph(context.Context, *Paragraph) (*Paragraph, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartParagraph not implemented")
}
func (UnimplementedTransactionServiceServer) ListNewDetail(context.Context, *NewDetailRequest) (*NewDetails, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListNewDetail not implemented")
}
func (UnimplementedTransactionServiceServer) ListNew(context.Context, *NewRequest) (*News, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListNew not implemented")
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

func _TransactionService_CompleteParagraph_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Paragraph)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).CompleteParagraph(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transaction.TransactionService/CompleteParagraph",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).CompleteParagraph(ctx, req.(*Paragraph))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_ListLookup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LookupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).ListLookup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transaction.TransactionService/ListLookup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).ListLookup(ctx, req.(*LookupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_ListSentence_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SentenceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).ListSentence(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transaction.TransactionService/ListSentence",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).ListSentence(ctx, req.(*SentenceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_GetParagraphMax_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ParagraphRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).GetParagraphMax(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transaction.TransactionService/GetParagraphMax",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).GetParagraphMax(ctx, req.(*ParagraphRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_ListParagraph_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ParagraphRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).ListParagraph(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transaction.TransactionService/ListParagraph",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).ListParagraph(ctx, req.(*ParagraphRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_ListListenDisplay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListenDisplayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).ListListenDisplay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transaction.TransactionService/ListListenDisplay",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).ListListenDisplay(ctx, req.(*ListenDisplayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_ListSpeakDisplay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SpeakDisplayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).ListSpeakDisplay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transaction.TransactionService/ListSpeakDisplay",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).ListSpeakDisplay(ctx, req.(*SpeakDisplayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_ListFavourite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FavouriteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).ListFavourite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transaction.TransactionService/ListFavourite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).ListFavourite(ctx, req.(*FavouriteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_CreateFavourite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Favourite)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).CreateFavourite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transaction.TransactionService/CreateFavourite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).CreateFavourite(ctx, req.(*Favourite))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_DeleteFavourite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Favourite)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).DeleteFavourite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transaction.TransactionService/DeleteFavourite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).DeleteFavourite(ctx, req.(*Favourite))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_StartParagraph_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Paragraph)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).StartParagraph(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transaction.TransactionService/StartParagraph",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).StartParagraph(ctx, req.(*Paragraph))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_ListNewDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).ListNewDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transaction.TransactionService/ListNewDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).ListNewDetail(ctx, req.(*NewDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_ListNew_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).ListNew(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transaction.TransactionService/ListNew",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).ListNew(ctx, req.(*NewRequest))
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
		{
			MethodName: "CompleteParagraph",
			Handler:    _TransactionService_CompleteParagraph_Handler,
		},
		{
			MethodName: "ListLookup",
			Handler:    _TransactionService_ListLookup_Handler,
		},
		{
			MethodName: "ListSentence",
			Handler:    _TransactionService_ListSentence_Handler,
		},
		{
			MethodName: "GetParagraphMax",
			Handler:    _TransactionService_GetParagraphMax_Handler,
		},
		{
			MethodName: "ListParagraph",
			Handler:    _TransactionService_ListParagraph_Handler,
		},
		{
			MethodName: "ListListenDisplay",
			Handler:    _TransactionService_ListListenDisplay_Handler,
		},
		{
			MethodName: "ListSpeakDisplay",
			Handler:    _TransactionService_ListSpeakDisplay_Handler,
		},
		{
			MethodName: "ListFavourite",
			Handler:    _TransactionService_ListFavourite_Handler,
		},
		{
			MethodName: "CreateFavourite",
			Handler:    _TransactionService_CreateFavourite_Handler,
		},
		{
			MethodName: "DeleteFavourite",
			Handler:    _TransactionService_DeleteFavourite_Handler,
		},
		{
			MethodName: "StartParagraph",
			Handler:    _TransactionService_StartParagraph_Handler,
		},
		{
			MethodName: "ListNewDetail",
			Handler:    _TransactionService_ListNewDetail_Handler,
		},
		{
			MethodName: "ListNew",
			Handler:    _TransactionService_ListNew_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "transaction.proto",
}
