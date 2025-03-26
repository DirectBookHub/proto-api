// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v6.30.1
// source: v1/flight_search.proto

package flightsearch

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
	FlightSearchService_SearchFlights_FullMethodName = "/flightsearch.FlightSearchService/SearchFlights"
)

// FlightSearchServiceClient is the client API for FlightSearchService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FlightSearchServiceClient interface {
	// Searches for flights based on the given query.
	SearchFlights(ctx context.Context, in *SearchFlightsRequest, opts ...grpc.CallOption) (*SearchFlightsResponse, error)
}

type flightSearchServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFlightSearchServiceClient(cc grpc.ClientConnInterface) FlightSearchServiceClient {
	return &flightSearchServiceClient{cc}
}

func (c *flightSearchServiceClient) SearchFlights(ctx context.Context, in *SearchFlightsRequest, opts ...grpc.CallOption) (*SearchFlightsResponse, error) {
	out := new(SearchFlightsResponse)
	err := c.cc.Invoke(ctx, FlightSearchService_SearchFlights_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FlightSearchServiceServer is the server API for FlightSearchService service.
// All implementations must embed UnimplementedFlightSearchServiceServer
// for forward compatibility
type FlightSearchServiceServer interface {
	// Searches for flights based on the given query.
	SearchFlights(context.Context, *SearchFlightsRequest) (*SearchFlightsResponse, error)
	mustEmbedUnimplementedFlightSearchServiceServer()
}

// UnimplementedFlightSearchServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFlightSearchServiceServer struct {
}

func (UnimplementedFlightSearchServiceServer) SearchFlights(context.Context, *SearchFlightsRequest) (*SearchFlightsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchFlights not implemented")
}
func (UnimplementedFlightSearchServiceServer) mustEmbedUnimplementedFlightSearchServiceServer() {}

// UnsafeFlightSearchServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FlightSearchServiceServer will
// result in compilation errors.
type UnsafeFlightSearchServiceServer interface {
	mustEmbedUnimplementedFlightSearchServiceServer()
}

func RegisterFlightSearchServiceServer(s grpc.ServiceRegistrar, srv FlightSearchServiceServer) {
	s.RegisterService(&FlightSearchService_ServiceDesc, srv)
}

func _FlightSearchService_SearchFlights_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchFlightsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlightSearchServiceServer).SearchFlights(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FlightSearchService_SearchFlights_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlightSearchServiceServer).SearchFlights(ctx, req.(*SearchFlightsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FlightSearchService_ServiceDesc is the grpc.ServiceDesc for FlightSearchService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FlightSearchService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "flightsearch.FlightSearchService",
	HandlerType: (*FlightSearchServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SearchFlights",
			Handler:    _FlightSearchService_SearchFlights_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/flight_search.proto",
}
