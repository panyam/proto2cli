// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.6
// source: labels.proto

package protos

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

const (
	LabelService_CreateLabel_FullMethodName = "/musicservice.LabelService/CreateLabel"
	LabelService_GetLabels_FullMethodName   = "/musicservice.LabelService/GetLabels"
	LabelService_UpdateLabel_FullMethodName = "/musicservice.LabelService/UpdateLabel"
	LabelService_DeleteLabel_FullMethodName = "/musicservice.LabelService/DeleteLabel"
	LabelService_ListLabels_FullMethodName  = "/musicservice.LabelService/ListLabels"
	LabelService_AddAlbum_FullMethodName    = "/musicservice.LabelService/AddAlbum"
	LabelService_RemoveAlbum_FullMethodName = "/musicservice.LabelService/RemoveAlbum"
	LabelService_ListAlbums_FullMethodName  = "/musicservice.LabelService/ListAlbums"
)

// LabelServiceClient is the client API for LabelService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LabelServiceClient interface {
	// *
	// Create a new Label
	CreateLabel(ctx context.Context, in *CreateLabelRequest, opts ...grpc.CallOption) (*CreateLabelResponse, error)
	// *
	// Batch gets multiple labels.
	GetLabels(ctx context.Context, in *GetLabelsRequest, opts ...grpc.CallOption) (*GetLabelsResponse, error)
	// *
	// Updates specific fields of an Label
	UpdateLabel(ctx context.Context, in *UpdateLabelRequest, opts ...grpc.CallOption) (*UpdateLabelResponse, error)
	// *
	// Deletes an label from our system.
	DeleteLabel(ctx context.Context, in *DeleteLabelRequest, opts ...grpc.CallOption) (*DeleteLabelResponse, error)
	// *
	// Finds and retrieves labels matching the particular criteria.
	ListLabels(ctx context.Context, in *ListLabelsRequest, opts ...grpc.CallOption) (*ListLabelsResponse, error)
	// *
	// Add an album to this label.
	AddAlbum(ctx context.Context, in *AddAlbumRequest, opts ...grpc.CallOption) (*AddAlbumResponse, error)
	// *
	// Remove an album from this label.
	RemoveAlbum(ctx context.Context, in *RemoveAlbumRequest, opts ...grpc.CallOption) (*RemoveAlbumResponse, error)
	// *
	// List all albums owned by this label.
	ListAlbums(ctx context.Context, in *ListAlbumsInLabelRequest, opts ...grpc.CallOption) (*ListAlbumsInLabelResponse, error)
}

type labelServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLabelServiceClient(cc grpc.ClientConnInterface) LabelServiceClient {
	return &labelServiceClient{cc}
}

func (c *labelServiceClient) CreateLabel(ctx context.Context, in *CreateLabelRequest, opts ...grpc.CallOption) (*CreateLabelResponse, error) {
	out := new(CreateLabelResponse)
	err := c.cc.Invoke(ctx, LabelService_CreateLabel_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *labelServiceClient) GetLabels(ctx context.Context, in *GetLabelsRequest, opts ...grpc.CallOption) (*GetLabelsResponse, error) {
	out := new(GetLabelsResponse)
	err := c.cc.Invoke(ctx, LabelService_GetLabels_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *labelServiceClient) UpdateLabel(ctx context.Context, in *UpdateLabelRequest, opts ...grpc.CallOption) (*UpdateLabelResponse, error) {
	out := new(UpdateLabelResponse)
	err := c.cc.Invoke(ctx, LabelService_UpdateLabel_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *labelServiceClient) DeleteLabel(ctx context.Context, in *DeleteLabelRequest, opts ...grpc.CallOption) (*DeleteLabelResponse, error) {
	out := new(DeleteLabelResponse)
	err := c.cc.Invoke(ctx, LabelService_DeleteLabel_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *labelServiceClient) ListLabels(ctx context.Context, in *ListLabelsRequest, opts ...grpc.CallOption) (*ListLabelsResponse, error) {
	out := new(ListLabelsResponse)
	err := c.cc.Invoke(ctx, LabelService_ListLabels_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *labelServiceClient) AddAlbum(ctx context.Context, in *AddAlbumRequest, opts ...grpc.CallOption) (*AddAlbumResponse, error) {
	out := new(AddAlbumResponse)
	err := c.cc.Invoke(ctx, LabelService_AddAlbum_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *labelServiceClient) RemoveAlbum(ctx context.Context, in *RemoveAlbumRequest, opts ...grpc.CallOption) (*RemoveAlbumResponse, error) {
	out := new(RemoveAlbumResponse)
	err := c.cc.Invoke(ctx, LabelService_RemoveAlbum_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *labelServiceClient) ListAlbums(ctx context.Context, in *ListAlbumsInLabelRequest, opts ...grpc.CallOption) (*ListAlbumsInLabelResponse, error) {
	out := new(ListAlbumsInLabelResponse)
	err := c.cc.Invoke(ctx, LabelService_ListAlbums_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LabelServiceServer is the server API for LabelService service.
// All implementations must embed UnimplementedLabelServiceServer
// for forward compatibility
type LabelServiceServer interface {
	// *
	// Create a new Label
	CreateLabel(context.Context, *CreateLabelRequest) (*CreateLabelResponse, error)
	// *
	// Batch gets multiple labels.
	GetLabels(context.Context, *GetLabelsRequest) (*GetLabelsResponse, error)
	// *
	// Updates specific fields of an Label
	UpdateLabel(context.Context, *UpdateLabelRequest) (*UpdateLabelResponse, error)
	// *
	// Deletes an label from our system.
	DeleteLabel(context.Context, *DeleteLabelRequest) (*DeleteLabelResponse, error)
	// *
	// Finds and retrieves labels matching the particular criteria.
	ListLabels(context.Context, *ListLabelsRequest) (*ListLabelsResponse, error)
	// *
	// Add an album to this label.
	AddAlbum(context.Context, *AddAlbumRequest) (*AddAlbumResponse, error)
	// *
	// Remove an album from this label.
	RemoveAlbum(context.Context, *RemoveAlbumRequest) (*RemoveAlbumResponse, error)
	// *
	// List all albums owned by this label.
	ListAlbums(context.Context, *ListAlbumsInLabelRequest) (*ListAlbumsInLabelResponse, error)
	mustEmbedUnimplementedLabelServiceServer()
}

// UnimplementedLabelServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLabelServiceServer struct {
}

func (UnimplementedLabelServiceServer) CreateLabel(context.Context, *CreateLabelRequest) (*CreateLabelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLabel not implemented")
}
func (UnimplementedLabelServiceServer) GetLabels(context.Context, *GetLabelsRequest) (*GetLabelsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLabels not implemented")
}
func (UnimplementedLabelServiceServer) UpdateLabel(context.Context, *UpdateLabelRequest) (*UpdateLabelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateLabel not implemented")
}
func (UnimplementedLabelServiceServer) DeleteLabel(context.Context, *DeleteLabelRequest) (*DeleteLabelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteLabel not implemented")
}
func (UnimplementedLabelServiceServer) ListLabels(context.Context, *ListLabelsRequest) (*ListLabelsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListLabels not implemented")
}
func (UnimplementedLabelServiceServer) AddAlbum(context.Context, *AddAlbumRequest) (*AddAlbumResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddAlbum not implemented")
}
func (UnimplementedLabelServiceServer) RemoveAlbum(context.Context, *RemoveAlbumRequest) (*RemoveAlbumResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveAlbum not implemented")
}
func (UnimplementedLabelServiceServer) ListAlbums(context.Context, *ListAlbumsInLabelRequest) (*ListAlbumsInLabelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAlbums not implemented")
}
func (UnimplementedLabelServiceServer) mustEmbedUnimplementedLabelServiceServer() {}

// UnsafeLabelServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LabelServiceServer will
// result in compilation errors.
type UnsafeLabelServiceServer interface {
	mustEmbedUnimplementedLabelServiceServer()
}

func RegisterLabelServiceServer(s grpc.ServiceRegistrar, srv LabelServiceServer) {
	s.RegisterService(&LabelService_ServiceDesc, srv)
}

func _LabelService_CreateLabel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateLabelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LabelServiceServer).CreateLabel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LabelService_CreateLabel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LabelServiceServer).CreateLabel(ctx, req.(*CreateLabelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LabelService_GetLabels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLabelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LabelServiceServer).GetLabels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LabelService_GetLabels_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LabelServiceServer).GetLabels(ctx, req.(*GetLabelsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LabelService_UpdateLabel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateLabelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LabelServiceServer).UpdateLabel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LabelService_UpdateLabel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LabelServiceServer).UpdateLabel(ctx, req.(*UpdateLabelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LabelService_DeleteLabel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteLabelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LabelServiceServer).DeleteLabel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LabelService_DeleteLabel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LabelServiceServer).DeleteLabel(ctx, req.(*DeleteLabelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LabelService_ListLabels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListLabelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LabelServiceServer).ListLabels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LabelService_ListLabels_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LabelServiceServer).ListLabels(ctx, req.(*ListLabelsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LabelService_AddAlbum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddAlbumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LabelServiceServer).AddAlbum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LabelService_AddAlbum_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LabelServiceServer).AddAlbum(ctx, req.(*AddAlbumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LabelService_RemoveAlbum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveAlbumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LabelServiceServer).RemoveAlbum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LabelService_RemoveAlbum_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LabelServiceServer).RemoveAlbum(ctx, req.(*RemoveAlbumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LabelService_ListAlbums_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAlbumsInLabelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LabelServiceServer).ListAlbums(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LabelService_ListAlbums_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LabelServiceServer).ListAlbums(ctx, req.(*ListAlbumsInLabelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LabelService_ServiceDesc is the grpc.ServiceDesc for LabelService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LabelService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "musicservice.LabelService",
	HandlerType: (*LabelServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateLabel",
			Handler:    _LabelService_CreateLabel_Handler,
		},
		{
			MethodName: "GetLabels",
			Handler:    _LabelService_GetLabels_Handler,
		},
		{
			MethodName: "UpdateLabel",
			Handler:    _LabelService_UpdateLabel_Handler,
		},
		{
			MethodName: "DeleteLabel",
			Handler:    _LabelService_DeleteLabel_Handler,
		},
		{
			MethodName: "ListLabels",
			Handler:    _LabelService_ListLabels_Handler,
		},
		{
			MethodName: "AddAlbum",
			Handler:    _LabelService_AddAlbum_Handler,
		},
		{
			MethodName: "RemoveAlbum",
			Handler:    _LabelService_RemoveAlbum_Handler,
		},
		{
			MethodName: "ListAlbums",
			Handler:    _LabelService_ListAlbums_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "labels.proto",
}
