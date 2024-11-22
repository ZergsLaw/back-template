// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.20.3
// source: api/twitter/v1/twitter.proto

package pb

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
	Twitter_CreatePost_FullMethodName      = "/api.twitter.v1.Twitter/CreatePost"
	Twitter_DeletePost_FullMethodName      = "/api.twitter.v1.Twitter/DeletePost"
	Twitter_GetPost_FullMethodName         = "/api.twitter.v1.Twitter/GetPost"
	Twitter_UpdatePost_FullMethodName      = "/api.twitter.v1.Twitter/UpdatePost"
	Twitter_SearchPostsUser_FullMethodName = "/api.twitter.v1.Twitter/SearchPostsUser"
)

// TwitterClient is the client API for Twitter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TwitterClient interface {
	CreatePost(ctx context.Context, in *CreatePostRequest, opts ...grpc.CallOption) (*CreatePostResponse, error)
	DeletePost(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	GetPost(ctx context.Context, in *GetPostRequest, opts ...grpc.CallOption) (*GetPostResponse, error)
	UpdatePost(ctx context.Context, in *UpdatePostRequest, opts ...grpc.CallOption) (*UpdatePostResponse, error)
	SearchPostsUser(ctx context.Context, in *SearchPostsUserRequest, opts ...grpc.CallOption) (*SearchPostsUserResponse, error)
}

type twitterClient struct {
	cc grpc.ClientConnInterface
}

func NewTwitterClient(cc grpc.ClientConnInterface) TwitterClient {
	return &twitterClient{cc}
}

func (c *twitterClient) CreatePost(ctx context.Context, in *CreatePostRequest, opts ...grpc.CallOption) (*CreatePostResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreatePostResponse)
	err := c.cc.Invoke(ctx, Twitter_CreatePost_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *twitterClient) DeletePost(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, Twitter_DeletePost_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *twitterClient) GetPost(ctx context.Context, in *GetPostRequest, opts ...grpc.CallOption) (*GetPostResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetPostResponse)
	err := c.cc.Invoke(ctx, Twitter_GetPost_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *twitterClient) UpdatePost(ctx context.Context, in *UpdatePostRequest, opts ...grpc.CallOption) (*UpdatePostResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdatePostResponse)
	err := c.cc.Invoke(ctx, Twitter_UpdatePost_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *twitterClient) SearchPostsUser(ctx context.Context, in *SearchPostsUserRequest, opts ...grpc.CallOption) (*SearchPostsUserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SearchPostsUserResponse)
	err := c.cc.Invoke(ctx, Twitter_SearchPostsUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TwitterServer is the server API for Twitter service.
// All implementations should embed UnimplementedTwitterServer
// for forward compatibility.
type TwitterServer interface {
	CreatePost(context.Context, *CreatePostRequest) (*CreatePostResponse, error)
	DeletePost(context.Context, *DeleteRequest) (*DeleteResponse, error)
	GetPost(context.Context, *GetPostRequest) (*GetPostResponse, error)
	UpdatePost(context.Context, *UpdatePostRequest) (*UpdatePostResponse, error)
	SearchPostsUser(context.Context, *SearchPostsUserRequest) (*SearchPostsUserResponse, error)
}

// UnimplementedTwitterServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTwitterServer struct{}

func (UnimplementedTwitterServer) CreatePost(context.Context, *CreatePostRequest) (*CreatePostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePost not implemented")
}
func (UnimplementedTwitterServer) DeletePost(context.Context, *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePost not implemented")
}
func (UnimplementedTwitterServer) GetPost(context.Context, *GetPostRequest) (*GetPostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPost not implemented")
}
func (UnimplementedTwitterServer) UpdatePost(context.Context, *UpdatePostRequest) (*UpdatePostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePost not implemented")
}
func (UnimplementedTwitterServer) SearchPostsUser(context.Context, *SearchPostsUserRequest) (*SearchPostsUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchPostsUser not implemented")
}
func (UnimplementedTwitterServer) testEmbeddedByValue() {}

// UnsafeTwitterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TwitterServer will
// result in compilation errors.
type UnsafeTwitterServer interface {
	mustEmbedUnimplementedTwitterServer()
}

func RegisterTwitterServer(s grpc.ServiceRegistrar, srv TwitterServer) {
	// If the following call pancis, it indicates UnimplementedTwitterServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Twitter_ServiceDesc, srv)
}

func _Twitter_CreatePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TwitterServer).CreatePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Twitter_CreatePost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TwitterServer).CreatePost(ctx, req.(*CreatePostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Twitter_DeletePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TwitterServer).DeletePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Twitter_DeletePost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TwitterServer).DeletePost(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Twitter_GetPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TwitterServer).GetPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Twitter_GetPost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TwitterServer).GetPost(ctx, req.(*GetPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Twitter_UpdatePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TwitterServer).UpdatePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Twitter_UpdatePost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TwitterServer).UpdatePost(ctx, req.(*UpdatePostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Twitter_SearchPostsUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchPostsUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TwitterServer).SearchPostsUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Twitter_SearchPostsUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TwitterServer).SearchPostsUser(ctx, req.(*SearchPostsUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Twitter_ServiceDesc is the grpc.ServiceDesc for Twitter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Twitter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.twitter.v1.Twitter",
	HandlerType: (*TwitterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePost",
			Handler:    _Twitter_CreatePost_Handler,
		},
		{
			MethodName: "DeletePost",
			Handler:    _Twitter_DeletePost_Handler,
		},
		{
			MethodName: "GetPost",
			Handler:    _Twitter_GetPost_Handler,
		},
		{
			MethodName: "UpdatePost",
			Handler:    _Twitter_UpdatePost_Handler,
		},
		{
			MethodName: "SearchPostsUser",
			Handler:    _Twitter_SearchPostsUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/twitter/v1/twitter.proto",
}