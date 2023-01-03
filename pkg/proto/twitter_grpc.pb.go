// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.8
// source: pkg/proto/twitter.proto

package proto

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// TwitterClient is the client API for Twitter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TwitterClient interface {
	AddTweet(ctx context.Context, in *AddTweetRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetNewsFeed(ctx context.Context, in *GetNewsFeedRequest, opts ...grpc.CallOption) (*GetNewsFeedResponse, error)
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	Follow(ctx context.Context, in *FollowRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type twitterClient struct {
	cc grpc.ClientConnInterface
}

func NewTwitterClient(cc grpc.ClientConnInterface) TwitterClient {
	return &twitterClient{cc}
}

func (c *twitterClient) AddTweet(ctx context.Context, in *AddTweetRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/Twitter/AddTweet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *twitterClient) GetNewsFeed(ctx context.Context, in *GetNewsFeedRequest, opts ...grpc.CallOption) (*GetNewsFeedResponse, error) {
	out := new(GetNewsFeedResponse)
	err := c.cc.Invoke(ctx, "/Twitter/GetNewsFeed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *twitterClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, "/Twitter/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *twitterClient) Follow(ctx context.Context, in *FollowRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/Twitter/Follow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TwitterServer is the server API for Twitter service.
// All implementations must embed UnimplementedTwitterServer
// for forward compatibility
type TwitterServer interface {
	AddTweet(context.Context, *AddTweetRequest) (*emptypb.Empty, error)
	GetNewsFeed(context.Context, *GetNewsFeedRequest) (*GetNewsFeedResponse, error)
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	Follow(context.Context, *FollowRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedTwitterServer()
}

// UnimplementedTwitterServer must be embedded to have forward compatible implementations.
type UnimplementedTwitterServer struct {
}

func (UnimplementedTwitterServer) AddTweet(context.Context, *AddTweetRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTweet not implemented")
}
func (UnimplementedTwitterServer) GetNewsFeed(context.Context, *GetNewsFeedRequest) (*GetNewsFeedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNewsFeed not implemented")
}
func (UnimplementedTwitterServer) Register(context.Context, *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedTwitterServer) Follow(context.Context, *FollowRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Follow not implemented")
}
func (UnimplementedTwitterServer) mustEmbedUnimplementedTwitterServer() {}

// UnsafeTwitterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TwitterServer will
// result in compilation errors.
type UnsafeTwitterServer interface {
	mustEmbedUnimplementedTwitterServer()
}

func RegisterTwitterServer(s grpc.ServiceRegistrar, srv TwitterServer) {
	s.RegisterService(&Twitter_ServiceDesc, srv)
}

func _Twitter_AddTweet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddTweetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TwitterServer).AddTweet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Twitter/AddTweet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TwitterServer).AddTweet(ctx, req.(*AddTweetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Twitter_GetNewsFeed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNewsFeedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TwitterServer).GetNewsFeed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Twitter/GetNewsFeed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TwitterServer).GetNewsFeed(ctx, req.(*GetNewsFeedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Twitter_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TwitterServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Twitter/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TwitterServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Twitter_Follow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FollowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TwitterServer).Follow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Twitter/Follow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TwitterServer).Follow(ctx, req.(*FollowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Twitter_ServiceDesc is the grpc.ServiceDesc for Twitter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Twitter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Twitter",
	HandlerType: (*TwitterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddTweet",
			Handler:    _Twitter_AddTweet_Handler,
		},
		{
			MethodName: "GetNewsFeed",
			Handler:    _Twitter_GetNewsFeed_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _Twitter_Register_Handler,
		},
		{
			MethodName: "Follow",
			Handler:    _Twitter_Follow_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/proto/twitter.proto",
}
