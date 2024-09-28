// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: library.proto

package library

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

// LibraryServiceClient is the client API for LibraryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LibraryServiceClient interface {
	// ---------------- Course -----------------
	ListCourses(ctx context.Context, in *CourseRequest, opts ...grpc.CallOption) (*Courses, error)
	ListDocument(ctx context.Context, in *DocumentRequest, opts ...grpc.CallOption) (*Documents, error)
	ListParagraph(ctx context.Context, in *ParagraphRequest, opts ...grpc.CallOption) (*Paragraphs, error)
	FindParagraph(ctx context.Context, in *ParagraphRequest, opts ...grpc.CallOption) (*Paragraph, error)
	GetDataSearchEngine(ctx context.Context, in *SearchEngineRequest, opts ...grpc.CallOption) (*SearchEngines, error)
	ListSentence(ctx context.Context, in *SentenceRequest, opts ...grpc.CallOption) (*Sentences, error)
	StartLearnListen(ctx context.Context, in *SentenceRequest, opts ...grpc.CallOption) (*Sentences, error)
	EndLearnListen(ctx context.Context, in *Listen, opts ...grpc.CallOption) (*Listen, error)
	GetPos(ctx context.Context, in *PosRequest, opts ...grpc.CallOption) (*Pos, error)
	ApproveParagraph(ctx context.Context, in *Paragraph, opts ...grpc.CallOption) (*Paragraph, error)
	ListSentencePos(ctx context.Context, in *SentencePosRequest, opts ...grpc.CallOption) (*SentencePoses, error)
	GetSentenceDetail(ctx context.Context, in *SentenceRequest, opts ...grpc.CallOption) (*Sentence, error)
	ListPronounce(ctx context.Context, in *PronounceRequest, opts ...grpc.CallOption) (*Pronounces, error)
	GetPronounce(ctx context.Context, in *PronounceRequest, opts ...grpc.CallOption) (*Pronounce, error)
	ListSentencePoses(ctx context.Context, in *SentencePosRequest, opts ...grpc.CallOption) (*SentencePoses, error)
	ListKnowledge(ctx context.Context, in *KnowledgeRequest, opts ...grpc.CallOption) (*Knowledges, error)
	GetKnowledge(ctx context.Context, in *Knowledge, opts ...grpc.CallOption) (*Knowledge, error)
	ListTagDetail(ctx context.Context, in *TagDetailRequest, opts ...grpc.CallOption) (*TagDetails, error)
	ListQuestion(ctx context.Context, in *QuestionRequest, opts ...grpc.CallOption) (*Questions, error)
	ListAnswer(ctx context.Context, in *AnswerRequest, opts ...grpc.CallOption) (*Answers, error)
	ListQuiz(ctx context.Context, in *QuizRequest, opts ...grpc.CallOption) (*Quizzes, error)
	ListSentenceGroup(ctx context.Context, in *SentenceGroupRequest, opts ...grpc.CallOption) (*SentenceGroups, error)
	ListTag(ctx context.Context, in *TagRequest, opts ...grpc.CallOption) (*Tags, error)
	FindTag(ctx context.Context, in *TagRequest, opts ...grpc.CallOption) (*Tag, error)
	UpdateSentences(ctx context.Context, in *Sentences, opts ...grpc.CallOption) (*Empty, error)
}

type libraryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLibraryServiceClient(cc grpc.ClientConnInterface) LibraryServiceClient {
	return &libraryServiceClient{cc}
}

func (c *libraryServiceClient) ListCourses(ctx context.Context, in *CourseRequest, opts ...grpc.CallOption) (*Courses, error) {
	out := new(Courses)
	err := c.cc.Invoke(ctx, "/library.LibraryService/ListCourses", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libraryServiceClient) ListDocument(ctx context.Context, in *DocumentRequest, opts ...grpc.CallOption) (*Documents, error) {
	out := new(Documents)
	err := c.cc.Invoke(ctx, "/library.LibraryService/ListDocument", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libraryServiceClient) ListParagraph(ctx context.Context, in *ParagraphRequest, opts ...grpc.CallOption) (*Paragraphs, error) {
	out := new(Paragraphs)
	err := c.cc.Invoke(ctx, "/library.LibraryService/ListParagraph", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libraryServiceClient) FindParagraph(ctx context.Context, in *ParagraphRequest, opts ...grpc.CallOption) (*Paragraph, error) {
	out := new(Paragraph)
	err := c.cc.Invoke(ctx, "/library.LibraryService/FindParagraph", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libraryServiceClient) GetDataSearchEngine(ctx context.Context, in *SearchEngineRequest, opts ...grpc.CallOption) (*SearchEngines, error) {
	out := new(SearchEngines)
	err := c.cc.Invoke(ctx, "/library.LibraryService/GetDataSearchEngine", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libraryServiceClient) ListSentence(ctx context.Context, in *SentenceRequest, opts ...grpc.CallOption) (*Sentences, error) {
	out := new(Sentences)
	err := c.cc.Invoke(ctx, "/library.LibraryService/ListSentence", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libraryServiceClient) StartLearnListen(ctx context.Context, in *SentenceRequest, opts ...grpc.CallOption) (*Sentences, error) {
	out := new(Sentences)
	err := c.cc.Invoke(ctx, "/library.LibraryService/StartLearnListen", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libraryServiceClient) EndLearnListen(ctx context.Context, in *Listen, opts ...grpc.CallOption) (*Listen, error) {
	out := new(Listen)
	err := c.cc.Invoke(ctx, "/library.LibraryService/EndLearnListen", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libraryServiceClient) GetPos(ctx context.Context, in *PosRequest, opts ...grpc.CallOption) (*Pos, error) {
	out := new(Pos)
	err := c.cc.Invoke(ctx, "/library.LibraryService/GetPos", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libraryServiceClient) ApproveParagraph(ctx context.Context, in *Paragraph, opts ...grpc.CallOption) (*Paragraph, error) {
	out := new(Paragraph)
	err := c.cc.Invoke(ctx, "/library.LibraryService/ApproveParagraph", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libraryServiceClient) ListSentencePos(ctx context.Context, in *SentencePosRequest, opts ...grpc.CallOption) (*SentencePoses, error) {
	out := new(SentencePoses)
	err := c.cc.Invoke(ctx, "/library.LibraryService/ListSentencePos", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libraryServiceClient) GetSentenceDetail(ctx context.Context, in *SentenceRequest, opts ...grpc.CallOption) (*Sentence, error) {
	out := new(Sentence)
	err := c.cc.Invoke(ctx, "/library.LibraryService/GetSentenceDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libraryServiceClient) ListPronounce(ctx context.Context, in *PronounceRequest, opts ...grpc.CallOption) (*Pronounces, error) {
	out := new(Pronounces)
	err := c.cc.Invoke(ctx, "/library.LibraryService/ListPronounce", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libraryServiceClient) GetPronounce(ctx context.Context, in *PronounceRequest, opts ...grpc.CallOption) (*Pronounce, error) {
	out := new(Pronounce)
	err := c.cc.Invoke(ctx, "/library.LibraryService/GetPronounce", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libraryServiceClient) ListSentencePoses(ctx context.Context, in *SentencePosRequest, opts ...grpc.CallOption) (*SentencePoses, error) {
	out := new(SentencePoses)
	err := c.cc.Invoke(ctx, "/library.LibraryService/ListSentencePoses", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libraryServiceClient) ListKnowledge(ctx context.Context, in *KnowledgeRequest, opts ...grpc.CallOption) (*Knowledges, error) {
	out := new(Knowledges)
	err := c.cc.Invoke(ctx, "/library.LibraryService/ListKnowledge", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libraryServiceClient) GetKnowledge(ctx context.Context, in *Knowledge, opts ...grpc.CallOption) (*Knowledge, error) {
	out := new(Knowledge)
	err := c.cc.Invoke(ctx, "/library.LibraryService/GetKnowledge", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libraryServiceClient) ListTagDetail(ctx context.Context, in *TagDetailRequest, opts ...grpc.CallOption) (*TagDetails, error) {
	out := new(TagDetails)
	err := c.cc.Invoke(ctx, "/library.LibraryService/ListTagDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libraryServiceClient) ListQuestion(ctx context.Context, in *QuestionRequest, opts ...grpc.CallOption) (*Questions, error) {
	out := new(Questions)
	err := c.cc.Invoke(ctx, "/library.LibraryService/ListQuestion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libraryServiceClient) ListAnswer(ctx context.Context, in *AnswerRequest, opts ...grpc.CallOption) (*Answers, error) {
	out := new(Answers)
	err := c.cc.Invoke(ctx, "/library.LibraryService/ListAnswer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libraryServiceClient) ListQuiz(ctx context.Context, in *QuizRequest, opts ...grpc.CallOption) (*Quizzes, error) {
	out := new(Quizzes)
	err := c.cc.Invoke(ctx, "/library.LibraryService/ListQuiz", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libraryServiceClient) ListSentenceGroup(ctx context.Context, in *SentenceGroupRequest, opts ...grpc.CallOption) (*SentenceGroups, error) {
	out := new(SentenceGroups)
	err := c.cc.Invoke(ctx, "/library.LibraryService/ListSentenceGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libraryServiceClient) ListTag(ctx context.Context, in *TagRequest, opts ...grpc.CallOption) (*Tags, error) {
	out := new(Tags)
	err := c.cc.Invoke(ctx, "/library.LibraryService/ListTag", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libraryServiceClient) FindTag(ctx context.Context, in *TagRequest, opts ...grpc.CallOption) (*Tag, error) {
	out := new(Tag)
	err := c.cc.Invoke(ctx, "/library.LibraryService/FindTag", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libraryServiceClient) UpdateSentences(ctx context.Context, in *Sentences, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/library.LibraryService/UpdateSentences", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LibraryServiceServer is the server API for LibraryService service.
// All implementations should embed UnimplementedLibraryServiceServer
// for forward compatibility
type LibraryServiceServer interface {
	// ---------------- Course -----------------
	ListCourses(context.Context, *CourseRequest) (*Courses, error)
	ListDocument(context.Context, *DocumentRequest) (*Documents, error)
	ListParagraph(context.Context, *ParagraphRequest) (*Paragraphs, error)
	FindParagraph(context.Context, *ParagraphRequest) (*Paragraph, error)
	GetDataSearchEngine(context.Context, *SearchEngineRequest) (*SearchEngines, error)
	ListSentence(context.Context, *SentenceRequest) (*Sentences, error)
	StartLearnListen(context.Context, *SentenceRequest) (*Sentences, error)
	EndLearnListen(context.Context, *Listen) (*Listen, error)
	GetPos(context.Context, *PosRequest) (*Pos, error)
	ApproveParagraph(context.Context, *Paragraph) (*Paragraph, error)
	ListSentencePos(context.Context, *SentencePosRequest) (*SentencePoses, error)
	GetSentenceDetail(context.Context, *SentenceRequest) (*Sentence, error)
	ListPronounce(context.Context, *PronounceRequest) (*Pronounces, error)
	GetPronounce(context.Context, *PronounceRequest) (*Pronounce, error)
	ListSentencePoses(context.Context, *SentencePosRequest) (*SentencePoses, error)
	ListKnowledge(context.Context, *KnowledgeRequest) (*Knowledges, error)
	GetKnowledge(context.Context, *Knowledge) (*Knowledge, error)
	ListTagDetail(context.Context, *TagDetailRequest) (*TagDetails, error)
	ListQuestion(context.Context, *QuestionRequest) (*Questions, error)
	ListAnswer(context.Context, *AnswerRequest) (*Answers, error)
	ListQuiz(context.Context, *QuizRequest) (*Quizzes, error)
	ListSentenceGroup(context.Context, *SentenceGroupRequest) (*SentenceGroups, error)
	ListTag(context.Context, *TagRequest) (*Tags, error)
	FindTag(context.Context, *TagRequest) (*Tag, error)
	UpdateSentences(context.Context, *Sentences) (*Empty, error)
}

// UnimplementedLibraryServiceServer should be embedded to have forward compatible implementations.
type UnimplementedLibraryServiceServer struct {
}

func (UnimplementedLibraryServiceServer) ListCourses(context.Context, *CourseRequest) (*Courses, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCourses not implemented")
}
func (UnimplementedLibraryServiceServer) ListDocument(context.Context, *DocumentRequest) (*Documents, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDocument not implemented")
}
func (UnimplementedLibraryServiceServer) ListParagraph(context.Context, *ParagraphRequest) (*Paragraphs, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListParagraph not implemented")
}
func (UnimplementedLibraryServiceServer) FindParagraph(context.Context, *ParagraphRequest) (*Paragraph, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindParagraph not implemented")
}
func (UnimplementedLibraryServiceServer) GetDataSearchEngine(context.Context, *SearchEngineRequest) (*SearchEngines, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDataSearchEngine not implemented")
}
func (UnimplementedLibraryServiceServer) ListSentence(context.Context, *SentenceRequest) (*Sentences, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSentence not implemented")
}
func (UnimplementedLibraryServiceServer) StartLearnListen(context.Context, *SentenceRequest) (*Sentences, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartLearnListen not implemented")
}
func (UnimplementedLibraryServiceServer) EndLearnListen(context.Context, *Listen) (*Listen, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EndLearnListen not implemented")
}
func (UnimplementedLibraryServiceServer) GetPos(context.Context, *PosRequest) (*Pos, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPos not implemented")
}
func (UnimplementedLibraryServiceServer) ApproveParagraph(context.Context, *Paragraph) (*Paragraph, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApproveParagraph not implemented")
}
func (UnimplementedLibraryServiceServer) ListSentencePos(context.Context, *SentencePosRequest) (*SentencePoses, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSentencePos not implemented")
}
func (UnimplementedLibraryServiceServer) GetSentenceDetail(context.Context, *SentenceRequest) (*Sentence, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSentenceDetail not implemented")
}
func (UnimplementedLibraryServiceServer) ListPronounce(context.Context, *PronounceRequest) (*Pronounces, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPronounce not implemented")
}
func (UnimplementedLibraryServiceServer) GetPronounce(context.Context, *PronounceRequest) (*Pronounce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPronounce not implemented")
}
func (UnimplementedLibraryServiceServer) ListSentencePoses(context.Context, *SentencePosRequest) (*SentencePoses, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSentencePoses not implemented")
}
func (UnimplementedLibraryServiceServer) ListKnowledge(context.Context, *KnowledgeRequest) (*Knowledges, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListKnowledge not implemented")
}
func (UnimplementedLibraryServiceServer) GetKnowledge(context.Context, *Knowledge) (*Knowledge, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetKnowledge not implemented")
}
func (UnimplementedLibraryServiceServer) ListTagDetail(context.Context, *TagDetailRequest) (*TagDetails, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTagDetail not implemented")
}
func (UnimplementedLibraryServiceServer) ListQuestion(context.Context, *QuestionRequest) (*Questions, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListQuestion not implemented")
}
func (UnimplementedLibraryServiceServer) ListAnswer(context.Context, *AnswerRequest) (*Answers, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAnswer not implemented")
}
func (UnimplementedLibraryServiceServer) ListQuiz(context.Context, *QuizRequest) (*Quizzes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListQuiz not implemented")
}
func (UnimplementedLibraryServiceServer) ListSentenceGroup(context.Context, *SentenceGroupRequest) (*SentenceGroups, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSentenceGroup not implemented")
}
func (UnimplementedLibraryServiceServer) ListTag(context.Context, *TagRequest) (*Tags, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTag not implemented")
}
func (UnimplementedLibraryServiceServer) FindTag(context.Context, *TagRequest) (*Tag, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindTag not implemented")
}
func (UnimplementedLibraryServiceServer) UpdateSentences(context.Context, *Sentences) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSentences not implemented")
}

// UnsafeLibraryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LibraryServiceServer will
// result in compilation errors.
type UnsafeLibraryServiceServer interface {
	mustEmbedUnimplementedLibraryServiceServer()
}

func RegisterLibraryServiceServer(s grpc.ServiceRegistrar, srv LibraryServiceServer) {
	s.RegisterService(&LibraryService_ServiceDesc, srv)
}

func _LibraryService_ListCourses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CourseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibraryServiceServer).ListCourses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/library.LibraryService/ListCourses",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibraryServiceServer).ListCourses(ctx, req.(*CourseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LibraryService_ListDocument_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DocumentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibraryServiceServer).ListDocument(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/library.LibraryService/ListDocument",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibraryServiceServer).ListDocument(ctx, req.(*DocumentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LibraryService_ListParagraph_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ParagraphRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibraryServiceServer).ListParagraph(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/library.LibraryService/ListParagraph",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibraryServiceServer).ListParagraph(ctx, req.(*ParagraphRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LibraryService_FindParagraph_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ParagraphRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibraryServiceServer).FindParagraph(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/library.LibraryService/FindParagraph",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibraryServiceServer).FindParagraph(ctx, req.(*ParagraphRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LibraryService_GetDataSearchEngine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchEngineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibraryServiceServer).GetDataSearchEngine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/library.LibraryService/GetDataSearchEngine",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibraryServiceServer).GetDataSearchEngine(ctx, req.(*SearchEngineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LibraryService_ListSentence_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SentenceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibraryServiceServer).ListSentence(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/library.LibraryService/ListSentence",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibraryServiceServer).ListSentence(ctx, req.(*SentenceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LibraryService_StartLearnListen_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SentenceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibraryServiceServer).StartLearnListen(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/library.LibraryService/StartLearnListen",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibraryServiceServer).StartLearnListen(ctx, req.(*SentenceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LibraryService_EndLearnListen_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Listen)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibraryServiceServer).EndLearnListen(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/library.LibraryService/EndLearnListen",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibraryServiceServer).EndLearnListen(ctx, req.(*Listen))
	}
	return interceptor(ctx, in, info, handler)
}

func _LibraryService_GetPos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PosRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibraryServiceServer).GetPos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/library.LibraryService/GetPos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibraryServiceServer).GetPos(ctx, req.(*PosRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LibraryService_ApproveParagraph_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Paragraph)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibraryServiceServer).ApproveParagraph(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/library.LibraryService/ApproveParagraph",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibraryServiceServer).ApproveParagraph(ctx, req.(*Paragraph))
	}
	return interceptor(ctx, in, info, handler)
}

func _LibraryService_ListSentencePos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SentencePosRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibraryServiceServer).ListSentencePos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/library.LibraryService/ListSentencePos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibraryServiceServer).ListSentencePos(ctx, req.(*SentencePosRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LibraryService_GetSentenceDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SentenceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibraryServiceServer).GetSentenceDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/library.LibraryService/GetSentenceDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibraryServiceServer).GetSentenceDetail(ctx, req.(*SentenceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LibraryService_ListPronounce_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PronounceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibraryServiceServer).ListPronounce(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/library.LibraryService/ListPronounce",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibraryServiceServer).ListPronounce(ctx, req.(*PronounceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LibraryService_GetPronounce_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PronounceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibraryServiceServer).GetPronounce(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/library.LibraryService/GetPronounce",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibraryServiceServer).GetPronounce(ctx, req.(*PronounceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LibraryService_ListSentencePoses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SentencePosRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibraryServiceServer).ListSentencePoses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/library.LibraryService/ListSentencePoses",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibraryServiceServer).ListSentencePoses(ctx, req.(*SentencePosRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LibraryService_ListKnowledge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KnowledgeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibraryServiceServer).ListKnowledge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/library.LibraryService/ListKnowledge",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibraryServiceServer).ListKnowledge(ctx, req.(*KnowledgeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LibraryService_GetKnowledge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Knowledge)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibraryServiceServer).GetKnowledge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/library.LibraryService/GetKnowledge",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibraryServiceServer).GetKnowledge(ctx, req.(*Knowledge))
	}
	return interceptor(ctx, in, info, handler)
}

func _LibraryService_ListTagDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TagDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibraryServiceServer).ListTagDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/library.LibraryService/ListTagDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibraryServiceServer).ListTagDetail(ctx, req.(*TagDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LibraryService_ListQuestion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuestionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibraryServiceServer).ListQuestion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/library.LibraryService/ListQuestion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibraryServiceServer).ListQuestion(ctx, req.(*QuestionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LibraryService_ListAnswer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AnswerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibraryServiceServer).ListAnswer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/library.LibraryService/ListAnswer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibraryServiceServer).ListAnswer(ctx, req.(*AnswerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LibraryService_ListQuiz_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuizRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibraryServiceServer).ListQuiz(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/library.LibraryService/ListQuiz",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibraryServiceServer).ListQuiz(ctx, req.(*QuizRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LibraryService_ListSentenceGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SentenceGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibraryServiceServer).ListSentenceGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/library.LibraryService/ListSentenceGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibraryServiceServer).ListSentenceGroup(ctx, req.(*SentenceGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LibraryService_ListTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibraryServiceServer).ListTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/library.LibraryService/ListTag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibraryServiceServer).ListTag(ctx, req.(*TagRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LibraryService_FindTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibraryServiceServer).FindTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/library.LibraryService/FindTag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibraryServiceServer).FindTag(ctx, req.(*TagRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LibraryService_UpdateSentences_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Sentences)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibraryServiceServer).UpdateSentences(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/library.LibraryService/UpdateSentences",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibraryServiceServer).UpdateSentences(ctx, req.(*Sentences))
	}
	return interceptor(ctx, in, info, handler)
}

// LibraryService_ServiceDesc is the grpc.ServiceDesc for LibraryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LibraryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "library.LibraryService",
	HandlerType: (*LibraryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListCourses",
			Handler:    _LibraryService_ListCourses_Handler,
		},
		{
			MethodName: "ListDocument",
			Handler:    _LibraryService_ListDocument_Handler,
		},
		{
			MethodName: "ListParagraph",
			Handler:    _LibraryService_ListParagraph_Handler,
		},
		{
			MethodName: "FindParagraph",
			Handler:    _LibraryService_FindParagraph_Handler,
		},
		{
			MethodName: "GetDataSearchEngine",
			Handler:    _LibraryService_GetDataSearchEngine_Handler,
		},
		{
			MethodName: "ListSentence",
			Handler:    _LibraryService_ListSentence_Handler,
		},
		{
			MethodName: "StartLearnListen",
			Handler:    _LibraryService_StartLearnListen_Handler,
		},
		{
			MethodName: "EndLearnListen",
			Handler:    _LibraryService_EndLearnListen_Handler,
		},
		{
			MethodName: "GetPos",
			Handler:    _LibraryService_GetPos_Handler,
		},
		{
			MethodName: "ApproveParagraph",
			Handler:    _LibraryService_ApproveParagraph_Handler,
		},
		{
			MethodName: "ListSentencePos",
			Handler:    _LibraryService_ListSentencePos_Handler,
		},
		{
			MethodName: "GetSentenceDetail",
			Handler:    _LibraryService_GetSentenceDetail_Handler,
		},
		{
			MethodName: "ListPronounce",
			Handler:    _LibraryService_ListPronounce_Handler,
		},
		{
			MethodName: "GetPronounce",
			Handler:    _LibraryService_GetPronounce_Handler,
		},
		{
			MethodName: "ListSentencePoses",
			Handler:    _LibraryService_ListSentencePoses_Handler,
		},
		{
			MethodName: "ListKnowledge",
			Handler:    _LibraryService_ListKnowledge_Handler,
		},
		{
			MethodName: "GetKnowledge",
			Handler:    _LibraryService_GetKnowledge_Handler,
		},
		{
			MethodName: "ListTagDetail",
			Handler:    _LibraryService_ListTagDetail_Handler,
		},
		{
			MethodName: "ListQuestion",
			Handler:    _LibraryService_ListQuestion_Handler,
		},
		{
			MethodName: "ListAnswer",
			Handler:    _LibraryService_ListAnswer_Handler,
		},
		{
			MethodName: "ListQuiz",
			Handler:    _LibraryService_ListQuiz_Handler,
		},
		{
			MethodName: "ListSentenceGroup",
			Handler:    _LibraryService_ListSentenceGroup_Handler,
		},
		{
			MethodName: "ListTag",
			Handler:    _LibraryService_ListTag_Handler,
		},
		{
			MethodName: "FindTag",
			Handler:    _LibraryService_FindTag_Handler,
		},
		{
			MethodName: "UpdateSentences",
			Handler:    _LibraryService_UpdateSentences_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "library.proto",
}
