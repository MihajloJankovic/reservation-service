// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.0
// source: appres.proto

package genfiles

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
	Reservation_GetReservation_FullMethodName                = "/reservation/GetReservation"
	Reservation_GetAllReservations_FullMethodName            = "/reservation/GetAllReservations"
	Reservation_SetReservation_FullMethodName                = "/reservation/SetReservation"
	Reservation_UpdateReservation_FullMethodName             = "/reservation/UpdateReservation"
	Reservation_DeleteByAccomnendation_FullMethodName        = "/reservation/DeleteByAccomnendation"
	Reservation_CheckActiveReservation_FullMethodName        = "/reservation/CheckActiveReservation"
	Reservation_CheckUsersActiveReservation_FullMethodName   = "/reservation/CheckUsersActiveReservation"
	Reservation_CheckActiveReservationByEmail_FullMethodName = "/reservation/CheckActiveReservationByEmail"
)

// ReservationClient is the client API for Reservation service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReservationClient interface {
	GetReservation(ctx context.Context, in *ReservationRequest, opts ...grpc.CallOption) (*DummyLista, error)
	GetAllReservations(ctx context.Context, in *Emptyaa, opts ...grpc.CallOption) (*DummyLista, error)
	SetReservation(ctx context.Context, in *ReservationResponse, opts ...grpc.CallOption) (*Emptyaa, error)
	UpdateReservation(ctx context.Context, in *ReservationResponse, opts ...grpc.CallOption) (*Emptyaa, error)
	DeleteByAccomnendation(ctx context.Context, in *DeleteRequestaa, opts ...grpc.CallOption) (*Emptyaa, error)
	CheckActiveReservation(ctx context.Context, in *DateFromDateTo, opts ...grpc.CallOption) (*Emptyaa, error)
	CheckUsersActiveReservation(ctx context.Context, in *EmailCheck, opts ...grpc.CallOption) (*Emptyaa, error)
	CheckActiveReservationByEmail(ctx context.Context, in *Emaill, opts ...grpc.CallOption) (*Emptyaa, error)
}

type reservationClient struct {
	cc grpc.ClientConnInterface
}

func NewReservationClient(cc grpc.ClientConnInterface) ReservationClient {
	return &reservationClient{cc}
}

func (c *reservationClient) GetReservation(ctx context.Context, in *ReservationRequest, opts ...grpc.CallOption) (*DummyLista, error) {
	out := new(DummyLista)
	err := c.cc.Invoke(ctx, Reservation_GetReservation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationClient) GetAllReservations(ctx context.Context, in *Emptyaa, opts ...grpc.CallOption) (*DummyLista, error) {
	out := new(DummyLista)
	err := c.cc.Invoke(ctx, Reservation_GetAllReservations_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationClient) SetReservation(ctx context.Context, in *ReservationResponse, opts ...grpc.CallOption) (*Emptyaa, error) {
	out := new(Emptyaa)
	err := c.cc.Invoke(ctx, Reservation_SetReservation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationClient) UpdateReservation(ctx context.Context, in *ReservationResponse, opts ...grpc.CallOption) (*Emptyaa, error) {
	out := new(Emptyaa)
	err := c.cc.Invoke(ctx, Reservation_UpdateReservation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationClient) DeleteByAccomnendation(ctx context.Context, in *DeleteRequestaa, opts ...grpc.CallOption) (*Emptyaa, error) {
	out := new(Emptyaa)
	err := c.cc.Invoke(ctx, Reservation_DeleteByAccomnendation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationClient) CheckActiveReservation(ctx context.Context, in *DateFromDateTo, opts ...grpc.CallOption) (*Emptyaa, error) {
	out := new(Emptyaa)
	err := c.cc.Invoke(ctx, Reservation_CheckActiveReservation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationClient) CheckUsersActiveReservation(ctx context.Context, in *EmailCheck, opts ...grpc.CallOption) (*Emptyaa, error) {
	out := new(Emptyaa)
	err := c.cc.Invoke(ctx, Reservation_CheckUsersActiveReservation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationClient) CheckActiveReservationByEmail(ctx context.Context, in *Emaill, opts ...grpc.CallOption) (*Emptyaa, error) {
	out := new(Emptyaa)
	err := c.cc.Invoke(ctx, Reservation_CheckActiveReservationByEmail_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReservationServer is the server API for Reservation service.
// All implementations must embed UnimplementedReservationServer
// for forward compatibility
type ReservationServer interface {
	GetReservation(context.Context, *ReservationRequest) (*DummyLista, error)
	GetAllReservations(context.Context, *Emptyaa) (*DummyLista, error)
	SetReservation(context.Context, *ReservationResponse) (*Emptyaa, error)
	UpdateReservation(context.Context, *ReservationResponse) (*Emptyaa, error)
	DeleteByAccomnendation(context.Context, *DeleteRequestaa) (*Emptyaa, error)
	CheckActiveReservation(context.Context, *DateFromDateTo) (*Emptyaa, error)
	CheckUsersActiveReservation(context.Context, *EmailCheck) (*Emptyaa, error)
	CheckActiveReservationByEmail(context.Context, *Emaill) (*Emptyaa, error)
	mustEmbedUnimplementedReservationServer()
}

// UnimplementedReservationServer must be embedded to have forward compatible implementations.
type UnimplementedReservationServer struct {
}

func (UnimplementedReservationServer) GetReservation(context.Context, *ReservationRequest) (*DummyLista, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReservation not implemented")
}
func (UnimplementedReservationServer) GetAllReservations(context.Context, *Emptyaa) (*DummyLista, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllReservations not implemented")
}
func (UnimplementedReservationServer) SetReservation(context.Context, *ReservationResponse) (*Emptyaa, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetReservation not implemented")
}
func (UnimplementedReservationServer) UpdateReservation(context.Context, *ReservationResponse) (*Emptyaa, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateReservation not implemented")
}
func (UnimplementedReservationServer) DeleteByAccomnendation(context.Context, *DeleteRequestaa) (*Emptyaa, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteByAccomnendation not implemented")
}
func (UnimplementedReservationServer) CheckActiveReservation(context.Context, *DateFromDateTo) (*Emptyaa, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckActiveReservation not implemented")
}
func (UnimplementedReservationServer) CheckUsersActiveReservation(context.Context, *EmailCheck) (*Emptyaa, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckUsersActiveReservation not implemented")
}
func (UnimplementedReservationServer) CheckActiveReservationByEmail(context.Context, *Emaill) (*Emptyaa, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckActiveReservationByEmail not implemented")
}
func (UnimplementedReservationServer) mustEmbedUnimplementedReservationServer() {}

// UnsafeReservationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReservationServer will
// result in compilation errors.
type UnsafeReservationServer interface {
	mustEmbedUnimplementedReservationServer()
}

func RegisterReservationServer(s grpc.ServiceRegistrar, srv ReservationServer) {
	s.RegisterService(&Reservation_ServiceDesc, srv)
}

func _Reservation_GetReservation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReservationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServer).GetReservation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Reservation_GetReservation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServer).GetReservation(ctx, req.(*ReservationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Reservation_GetAllReservations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Emptyaa)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServer).GetAllReservations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Reservation_GetAllReservations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServer).GetAllReservations(ctx, req.(*Emptyaa))
	}
	return interceptor(ctx, in, info, handler)
}

func _Reservation_SetReservation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReservationResponse)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServer).SetReservation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Reservation_SetReservation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServer).SetReservation(ctx, req.(*ReservationResponse))
	}
	return interceptor(ctx, in, info, handler)
}

func _Reservation_UpdateReservation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReservationResponse)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServer).UpdateReservation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Reservation_UpdateReservation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServer).UpdateReservation(ctx, req.(*ReservationResponse))
	}
	return interceptor(ctx, in, info, handler)
}

func _Reservation_DeleteByAccomnendation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequestaa)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServer).DeleteByAccomnendation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Reservation_DeleteByAccomnendation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServer).DeleteByAccomnendation(ctx, req.(*DeleteRequestaa))
	}
	return interceptor(ctx, in, info, handler)
}

func _Reservation_CheckActiveReservation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DateFromDateTo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServer).CheckActiveReservation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Reservation_CheckActiveReservation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServer).CheckActiveReservation(ctx, req.(*DateFromDateTo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Reservation_CheckUsersActiveReservation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmailCheck)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServer).CheckUsersActiveReservation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Reservation_CheckUsersActiveReservation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServer).CheckUsersActiveReservation(ctx, req.(*EmailCheck))
	}
	return interceptor(ctx, in, info, handler)
}

func _Reservation_CheckActiveReservationByEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Emaill)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServer).CheckActiveReservationByEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Reservation_CheckActiveReservationByEmail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServer).CheckActiveReservationByEmail(ctx, req.(*Emaill))
	}
	return interceptor(ctx, in, info, handler)
}

// Reservation_ServiceDesc is the grpc.ServiceDesc for Reservation service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Reservation_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "reservation",
	HandlerType: (*ReservationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetReservation",
			Handler:    _Reservation_GetReservation_Handler,
		},
		{
			MethodName: "GetAllReservations",
			Handler:    _Reservation_GetAllReservations_Handler,
		},
		{
			MethodName: "SetReservation",
			Handler:    _Reservation_SetReservation_Handler,
		},
		{
			MethodName: "UpdateReservation",
			Handler:    _Reservation_UpdateReservation_Handler,
		},
		{
			MethodName: "DeleteByAccomnendation",
			Handler:    _Reservation_DeleteByAccomnendation_Handler,
		},
		{
			MethodName: "CheckActiveReservation",
			Handler:    _Reservation_CheckActiveReservation_Handler,
		},
		{
			MethodName: "CheckUsersActiveReservation",
			Handler:    _Reservation_CheckUsersActiveReservation_Handler,
		},
		{
			MethodName: "CheckActiveReservationByEmail",
			Handler:    _Reservation_CheckActiveReservationByEmail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "appres.proto",
}
