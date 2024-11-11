// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.12.4
// source: contracts/proto/profiles/profile.proto

package proto

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
	Profile_AddProfile_FullMethodName            = "/protobuf.Profile/AddProfile"
	Profile_UpdateProfile_FullMethodName         = "/protobuf.Profile/UpdateProfile"
	Profile_DeleteProfile_FullMethodName         = "/protobuf.Profile/DeleteProfile"
	Profile_GetProfileBySessionId_FullMethodName = "/protobuf.Profile/GetProfileBySessionId"
	Profile_GetProfileDetail_FullMethodName      = "/protobuf.Profile/GetProfileDetail"
	Profile_GetProfileShortInfo_FullMethodName   = "/protobuf.Profile/GetProfileShortInfo"
	Profile_GetProfileList_FullMethodName        = "/protobuf.Profile/GetProfileList"
	Profile_GetImageBySessionId_FullMethodName   = "/protobuf.Profile/GetImageBySessionId"
	Profile_GetImageById_FullMethodName          = "/protobuf.Profile/GetImageById"
	Profile_DeleteImage_FullMethodName           = "/protobuf.Profile/DeleteImage"
	Profile_GetFilterBySessionId_FullMethodName  = "/protobuf.Profile/GetFilterBySessionId"
	Profile_UpdateFilter_FullMethodName          = "/protobuf.Profile/UpdateFilter"
	Profile_AddBlock_FullMethodName              = "/protobuf.Profile/AddBlock"
	Profile_AddLike_FullMethodName               = "/protobuf.Profile/AddLike"
	Profile_UpdateLike_FullMethodName            = "/protobuf.Profile/UpdateLike"
	Profile_AddComplaint_FullMethodName          = "/protobuf.Profile/AddComplaint"
	Profile_UpdateCoordinates_FullMethodName     = "/protobuf.Profile/UpdateCoordinates"
)

// ProfileClient is the client API for Profile service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Описание сервиса Profile
type ProfileClient interface {
	AddProfile(ctx context.Context, in *ProfileAddRequest, opts ...grpc.CallOption) (*ProfileAddResponse, error)
	UpdateProfile(ctx context.Context, in *ProfileUpdateRequest, opts ...grpc.CallOption) (*ProfileBySessionIdResponse, error)
	DeleteProfile(ctx context.Context, in *ProfileDeleteRequest, opts ...grpc.CallOption) (*ProfileDeleteResponse, error)
	GetProfileBySessionId(ctx context.Context, in *ProfileGetBySessionIdRequest, opts ...grpc.CallOption) (*ProfileBySessionIdResponse, error)
	GetProfileDetail(ctx context.Context, in *ProfileGetDetailRequest, opts ...grpc.CallOption) (*ProfileDetailResponse, error)
	GetProfileShortInfo(ctx context.Context, in *ProfileGetShortInfoRequest, opts ...grpc.CallOption) (*ProfileShortInfoResponse, error)
	GetProfileList(ctx context.Context, in *ProfileGetListRequest, opts ...grpc.CallOption) (*ProfileListResponse, error)
	GetImageBySessionId(ctx context.Context, in *GetImageBySessionIdRequest, opts ...grpc.CallOption) (*ImageBySessionIdResponse, error)
	GetImageById(ctx context.Context, in *GetImageByIdRequest, opts ...grpc.CallOption) (*Image, error)
	DeleteImage(ctx context.Context, in *ImageDeleteRequest, opts ...grpc.CallOption) (*ImageDeleteResponse, error)
	GetFilterBySessionId(ctx context.Context, in *FilterGetRequest, opts ...grpc.CallOption) (*FilterGetResponse, error)
	UpdateFilter(ctx context.Context, in *FilterUpdateRequest, opts ...grpc.CallOption) (*FilterUpdateResponse, error)
	AddBlock(ctx context.Context, in *BlockAddRequest, opts ...grpc.CallOption) (*BlockAddResponse, error)
	AddLike(ctx context.Context, in *LikeAddRequest, opts ...grpc.CallOption) (*LikeAddResponse, error)
	UpdateLike(ctx context.Context, in *LikeUpdateRequest, opts ...grpc.CallOption) (*LikeUpdateResponse, error)
	AddComplaint(ctx context.Context, in *ComplaintAddRequest, opts ...grpc.CallOption) (*ComplaintAddResponse, error)
	UpdateCoordinates(ctx context.Context, in *NavigatorUpdateRequest, opts ...grpc.CallOption) (*NavigatorUpdateResponse, error)
}

type profileClient struct {
	cc grpc.ClientConnInterface
}

func NewProfileClient(cc grpc.ClientConnInterface) ProfileClient {
	return &profileClient{cc}
}

func (c *profileClient) AddProfile(ctx context.Context, in *ProfileAddRequest, opts ...grpc.CallOption) (*ProfileAddResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ProfileAddResponse)
	err := c.cc.Invoke(ctx, Profile_AddProfile_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileClient) UpdateProfile(ctx context.Context, in *ProfileUpdateRequest, opts ...grpc.CallOption) (*ProfileBySessionIdResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ProfileBySessionIdResponse)
	err := c.cc.Invoke(ctx, Profile_UpdateProfile_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileClient) DeleteProfile(ctx context.Context, in *ProfileDeleteRequest, opts ...grpc.CallOption) (*ProfileDeleteResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ProfileDeleteResponse)
	err := c.cc.Invoke(ctx, Profile_DeleteProfile_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileClient) GetProfileBySessionId(ctx context.Context, in *ProfileGetBySessionIdRequest, opts ...grpc.CallOption) (*ProfileBySessionIdResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ProfileBySessionIdResponse)
	err := c.cc.Invoke(ctx, Profile_GetProfileBySessionId_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileClient) GetProfileDetail(ctx context.Context, in *ProfileGetDetailRequest, opts ...grpc.CallOption) (*ProfileDetailResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ProfileDetailResponse)
	err := c.cc.Invoke(ctx, Profile_GetProfileDetail_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileClient) GetProfileShortInfo(ctx context.Context, in *ProfileGetShortInfoRequest, opts ...grpc.CallOption) (*ProfileShortInfoResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ProfileShortInfoResponse)
	err := c.cc.Invoke(ctx, Profile_GetProfileShortInfo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileClient) GetProfileList(ctx context.Context, in *ProfileGetListRequest, opts ...grpc.CallOption) (*ProfileListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ProfileListResponse)
	err := c.cc.Invoke(ctx, Profile_GetProfileList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileClient) GetImageBySessionId(ctx context.Context, in *GetImageBySessionIdRequest, opts ...grpc.CallOption) (*ImageBySessionIdResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ImageBySessionIdResponse)
	err := c.cc.Invoke(ctx, Profile_GetImageBySessionId_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileClient) GetImageById(ctx context.Context, in *GetImageByIdRequest, opts ...grpc.CallOption) (*Image, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Image)
	err := c.cc.Invoke(ctx, Profile_GetImageById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileClient) DeleteImage(ctx context.Context, in *ImageDeleteRequest, opts ...grpc.CallOption) (*ImageDeleteResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ImageDeleteResponse)
	err := c.cc.Invoke(ctx, Profile_DeleteImage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileClient) GetFilterBySessionId(ctx context.Context, in *FilterGetRequest, opts ...grpc.CallOption) (*FilterGetResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FilterGetResponse)
	err := c.cc.Invoke(ctx, Profile_GetFilterBySessionId_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileClient) UpdateFilter(ctx context.Context, in *FilterUpdateRequest, opts ...grpc.CallOption) (*FilterUpdateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FilterUpdateResponse)
	err := c.cc.Invoke(ctx, Profile_UpdateFilter_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileClient) AddBlock(ctx context.Context, in *BlockAddRequest, opts ...grpc.CallOption) (*BlockAddResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BlockAddResponse)
	err := c.cc.Invoke(ctx, Profile_AddBlock_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileClient) AddLike(ctx context.Context, in *LikeAddRequest, opts ...grpc.CallOption) (*LikeAddResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LikeAddResponse)
	err := c.cc.Invoke(ctx, Profile_AddLike_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileClient) UpdateLike(ctx context.Context, in *LikeUpdateRequest, opts ...grpc.CallOption) (*LikeUpdateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LikeUpdateResponse)
	err := c.cc.Invoke(ctx, Profile_UpdateLike_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileClient) AddComplaint(ctx context.Context, in *ComplaintAddRequest, opts ...grpc.CallOption) (*ComplaintAddResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ComplaintAddResponse)
	err := c.cc.Invoke(ctx, Profile_AddComplaint_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileClient) UpdateCoordinates(ctx context.Context, in *NavigatorUpdateRequest, opts ...grpc.CallOption) (*NavigatorUpdateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(NavigatorUpdateResponse)
	err := c.cc.Invoke(ctx, Profile_UpdateCoordinates_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProfileServer is the server API for Profile service.
// All implementations must embed UnimplementedProfileServer
// for forward compatibility.
//
// Описание сервиса Profile
type ProfileServer interface {
	AddProfile(context.Context, *ProfileAddRequest) (*ProfileAddResponse, error)
	UpdateProfile(context.Context, *ProfileUpdateRequest) (*ProfileBySessionIdResponse, error)
	DeleteProfile(context.Context, *ProfileDeleteRequest) (*ProfileDeleteResponse, error)
	GetProfileBySessionId(context.Context, *ProfileGetBySessionIdRequest) (*ProfileBySessionIdResponse, error)
	GetProfileDetail(context.Context, *ProfileGetDetailRequest) (*ProfileDetailResponse, error)
	GetProfileShortInfo(context.Context, *ProfileGetShortInfoRequest) (*ProfileShortInfoResponse, error)
	GetProfileList(context.Context, *ProfileGetListRequest) (*ProfileListResponse, error)
	GetImageBySessionId(context.Context, *GetImageBySessionIdRequest) (*ImageBySessionIdResponse, error)
	GetImageById(context.Context, *GetImageByIdRequest) (*Image, error)
	DeleteImage(context.Context, *ImageDeleteRequest) (*ImageDeleteResponse, error)
	GetFilterBySessionId(context.Context, *FilterGetRequest) (*FilterGetResponse, error)
	UpdateFilter(context.Context, *FilterUpdateRequest) (*FilterUpdateResponse, error)
	AddBlock(context.Context, *BlockAddRequest) (*BlockAddResponse, error)
	AddLike(context.Context, *LikeAddRequest) (*LikeAddResponse, error)
	UpdateLike(context.Context, *LikeUpdateRequest) (*LikeUpdateResponse, error)
	AddComplaint(context.Context, *ComplaintAddRequest) (*ComplaintAddResponse, error)
	UpdateCoordinates(context.Context, *NavigatorUpdateRequest) (*NavigatorUpdateResponse, error)
	mustEmbedUnimplementedProfileServer()
}

// UnimplementedProfileServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedProfileServer struct{}

func (UnimplementedProfileServer) AddProfile(context.Context, *ProfileAddRequest) (*ProfileAddResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddProfile not implemented")
}
func (UnimplementedProfileServer) UpdateProfile(context.Context, *ProfileUpdateRequest) (*ProfileBySessionIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProfile not implemented")
}
func (UnimplementedProfileServer) DeleteProfile(context.Context, *ProfileDeleteRequest) (*ProfileDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProfile not implemented")
}
func (UnimplementedProfileServer) GetProfileBySessionId(context.Context, *ProfileGetBySessionIdRequest) (*ProfileBySessionIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProfileBySessionId not implemented")
}
func (UnimplementedProfileServer) GetProfileDetail(context.Context, *ProfileGetDetailRequest) (*ProfileDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProfileDetail not implemented")
}
func (UnimplementedProfileServer) GetProfileShortInfo(context.Context, *ProfileGetShortInfoRequest) (*ProfileShortInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProfileShortInfo not implemented")
}
func (UnimplementedProfileServer) GetProfileList(context.Context, *ProfileGetListRequest) (*ProfileListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProfileList not implemented")
}
func (UnimplementedProfileServer) GetImageBySessionId(context.Context, *GetImageBySessionIdRequest) (*ImageBySessionIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetImageBySessionId not implemented")
}
func (UnimplementedProfileServer) GetImageById(context.Context, *GetImageByIdRequest) (*Image, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetImageById not implemented")
}
func (UnimplementedProfileServer) DeleteImage(context.Context, *ImageDeleteRequest) (*ImageDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteImage not implemented")
}
func (UnimplementedProfileServer) GetFilterBySessionId(context.Context, *FilterGetRequest) (*FilterGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFilterBySessionId not implemented")
}
func (UnimplementedProfileServer) UpdateFilter(context.Context, *FilterUpdateRequest) (*FilterUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFilter not implemented")
}
func (UnimplementedProfileServer) AddBlock(context.Context, *BlockAddRequest) (*BlockAddResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddBlock not implemented")
}
func (UnimplementedProfileServer) AddLike(context.Context, *LikeAddRequest) (*LikeAddResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddLike not implemented")
}
func (UnimplementedProfileServer) UpdateLike(context.Context, *LikeUpdateRequest) (*LikeUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateLike not implemented")
}
func (UnimplementedProfileServer) AddComplaint(context.Context, *ComplaintAddRequest) (*ComplaintAddResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddComplaint not implemented")
}
func (UnimplementedProfileServer) UpdateCoordinates(context.Context, *NavigatorUpdateRequest) (*NavigatorUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCoordinates not implemented")
}
func (UnimplementedProfileServer) mustEmbedUnimplementedProfileServer() {}
func (UnimplementedProfileServer) testEmbeddedByValue()                 {}

// UnsafeProfileServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProfileServer will
// result in compilation errors.
type UnsafeProfileServer interface {
	mustEmbedUnimplementedProfileServer()
}

func RegisterProfileServer(s grpc.ServiceRegistrar, srv ProfileServer) {
	// If the following call pancis, it indicates UnimplementedProfileServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Profile_ServiceDesc, srv)
}

func _Profile_AddProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProfileAddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).AddProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Profile_AddProfile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).AddProfile(ctx, req.(*ProfileAddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profile_UpdateProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProfileUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).UpdateProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Profile_UpdateProfile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).UpdateProfile(ctx, req.(*ProfileUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profile_DeleteProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProfileDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).DeleteProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Profile_DeleteProfile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).DeleteProfile(ctx, req.(*ProfileDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profile_GetProfileBySessionId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProfileGetBySessionIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).GetProfileBySessionId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Profile_GetProfileBySessionId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).GetProfileBySessionId(ctx, req.(*ProfileGetBySessionIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profile_GetProfileDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProfileGetDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).GetProfileDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Profile_GetProfileDetail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).GetProfileDetail(ctx, req.(*ProfileGetDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profile_GetProfileShortInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProfileGetShortInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).GetProfileShortInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Profile_GetProfileShortInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).GetProfileShortInfo(ctx, req.(*ProfileGetShortInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profile_GetProfileList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProfileGetListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).GetProfileList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Profile_GetProfileList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).GetProfileList(ctx, req.(*ProfileGetListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profile_GetImageBySessionId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetImageBySessionIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).GetImageBySessionId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Profile_GetImageBySessionId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).GetImageBySessionId(ctx, req.(*GetImageBySessionIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profile_GetImageById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetImageByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).GetImageById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Profile_GetImageById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).GetImageById(ctx, req.(*GetImageByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profile_DeleteImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImageDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).DeleteImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Profile_DeleteImage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).DeleteImage(ctx, req.(*ImageDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profile_GetFilterBySessionId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FilterGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).GetFilterBySessionId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Profile_GetFilterBySessionId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).GetFilterBySessionId(ctx, req.(*FilterGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profile_UpdateFilter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FilterUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).UpdateFilter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Profile_UpdateFilter_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).UpdateFilter(ctx, req.(*FilterUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profile_AddBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlockAddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).AddBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Profile_AddBlock_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).AddBlock(ctx, req.(*BlockAddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profile_AddLike_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LikeAddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).AddLike(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Profile_AddLike_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).AddLike(ctx, req.(*LikeAddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profile_UpdateLike_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LikeUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).UpdateLike(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Profile_UpdateLike_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).UpdateLike(ctx, req.(*LikeUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profile_AddComplaint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ComplaintAddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).AddComplaint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Profile_AddComplaint_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).AddComplaint(ctx, req.(*ComplaintAddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profile_UpdateCoordinates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NavigatorUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).UpdateCoordinates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Profile_UpdateCoordinates_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).UpdateCoordinates(ctx, req.(*NavigatorUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Profile_ServiceDesc is the grpc.ServiceDesc for Profile service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Profile_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protobuf.Profile",
	HandlerType: (*ProfileServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddProfile",
			Handler:    _Profile_AddProfile_Handler,
		},
		{
			MethodName: "UpdateProfile",
			Handler:    _Profile_UpdateProfile_Handler,
		},
		{
			MethodName: "DeleteProfile",
			Handler:    _Profile_DeleteProfile_Handler,
		},
		{
			MethodName: "GetProfileBySessionId",
			Handler:    _Profile_GetProfileBySessionId_Handler,
		},
		{
			MethodName: "GetProfileDetail",
			Handler:    _Profile_GetProfileDetail_Handler,
		},
		{
			MethodName: "GetProfileShortInfo",
			Handler:    _Profile_GetProfileShortInfo_Handler,
		},
		{
			MethodName: "GetProfileList",
			Handler:    _Profile_GetProfileList_Handler,
		},
		{
			MethodName: "GetImageBySessionId",
			Handler:    _Profile_GetImageBySessionId_Handler,
		},
		{
			MethodName: "GetImageById",
			Handler:    _Profile_GetImageById_Handler,
		},
		{
			MethodName: "DeleteImage",
			Handler:    _Profile_DeleteImage_Handler,
		},
		{
			MethodName: "GetFilterBySessionId",
			Handler:    _Profile_GetFilterBySessionId_Handler,
		},
		{
			MethodName: "UpdateFilter",
			Handler:    _Profile_UpdateFilter_Handler,
		},
		{
			MethodName: "AddBlock",
			Handler:    _Profile_AddBlock_Handler,
		},
		{
			MethodName: "AddLike",
			Handler:    _Profile_AddLike_Handler,
		},
		{
			MethodName: "UpdateLike",
			Handler:    _Profile_UpdateLike_Handler,
		},
		{
			MethodName: "AddComplaint",
			Handler:    _Profile_AddComplaint_Handler,
		},
		{
			MethodName: "UpdateCoordinates",
			Handler:    _Profile_UpdateCoordinates_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "contracts/proto/profiles/profile.proto",
}
