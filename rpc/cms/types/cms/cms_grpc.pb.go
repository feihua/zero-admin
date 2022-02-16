// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package cms

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

// CmsClient is the client API for Cms service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CmsClient interface {
	Greet(ctx context.Context, in *StreamReq, opts ...grpc.CallOption) (*StreamResp, error)
}

type cmsClient struct {
	cc grpc.ClientConnInterface
}

func NewCmsClient(cc grpc.ClientConnInterface) CmsClient {
	return &cmsClient{cc}
}

func (c *cmsClient) Greet(ctx context.Context, in *StreamReq, opts ...grpc.CallOption) (*StreamResp, error) {
	out := new(StreamResp)
	err := c.cc.Invoke(ctx, "/cms.Cms/greet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CmsServer is the server API for Cms service.
// All implementations must embed UnimplementedCmsServer
// for forward compatibility
type CmsServer interface {
	Greet(context.Context, *StreamReq) (*StreamResp, error)
	mustEmbedUnimplementedCmsServer()
}

// UnimplementedCmsServer must be embedded to have forward compatible implementations.
type UnimplementedCmsServer struct {
}

func (UnimplementedCmsServer) Greet(context.Context, *StreamReq) (*StreamResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Greet not implemented")
}
func (UnimplementedCmsServer) mustEmbedUnimplementedCmsServer() {}

// UnsafeCmsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CmsServer will
// result in compilation errors.
type UnsafeCmsServer interface {
	mustEmbedUnimplementedCmsServer()
}

func RegisterCmsServer(s grpc.ServiceRegistrar, srv CmsServer) {
	s.RegisterService(&Cms_ServiceDesc, srv)
}

func _Cms_Greet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StreamReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmsServer).Greet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cms.Cms/greet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmsServer).Greet(ctx, req.(*StreamReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Cms_ServiceDesc is the grpc.ServiceDesc for Cms service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Cms_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cms.Cms",
	HandlerType: (*CmsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "greet",
			Handler:    _Cms_Greet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cms.proto",
}
