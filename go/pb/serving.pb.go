// Code generated by protoc-gen-go.
// source: serving.proto
// DO NOT EDIT!

/*
Package serving is a generated protocol buffer package.

It is generated from these files:
	serving.proto

It has these top-level messages:
	DetectionResponse
	DetectionRequest
	Point
	Rectangle
	Region
	Shape
	Object
*/
package serving

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type DetectionResponse struct {
	Type       string       `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	Points     []*Point     `protobuf:"bytes,2,rep,name=points" json:"points,omitempty"`
	Objects    []*Object    `protobuf:"bytes,3,rep,name=objects" json:"objects,omitempty"`
	Rectangles []*Rectangle `protobuf:"bytes,4,rep,name=rectangles" json:"rectangles,omitempty"`
}

func (m *DetectionResponse) Reset()                    { *m = DetectionResponse{} }
func (m *DetectionResponse) String() string            { return proto.CompactTextString(m) }
func (*DetectionResponse) ProtoMessage()               {}
func (*DetectionResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *DetectionResponse) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *DetectionResponse) GetPoints() []*Point {
	if m != nil {
		return m.Points
	}
	return nil
}

func (m *DetectionResponse) GetObjects() []*Object {
	if m != nil {
		return m.Objects
	}
	return nil
}

func (m *DetectionResponse) GetRectangles() []*Rectangle {
	if m != nil {
		return m.Rectangles
	}
	return nil
}

type DetectionRequest struct {
	Image  []byte     `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	Region *Rectangle `protobuf:"bytes,2,opt,name=region" json:"region,omitempty"`
}

func (m *DetectionRequest) Reset()                    { *m = DetectionRequest{} }
func (m *DetectionRequest) String() string            { return proto.CompactTextString(m) }
func (*DetectionRequest) ProtoMessage()               {}
func (*DetectionRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *DetectionRequest) GetImage() []byte {
	if m != nil {
		return m.Image
	}
	return nil
}

func (m *DetectionRequest) GetRegion() *Rectangle {
	if m != nil {
		return m.Region
	}
	return nil
}

type Point struct {
	X int32 `protobuf:"varint,1,opt,name=x" json:"x,omitempty"`
	Y int32 `protobuf:"varint,2,opt,name=y" json:"y,omitempty"`
}

func (m *Point) Reset()                    { *m = Point{} }
func (m *Point) String() string            { return proto.CompactTextString(m) }
func (*Point) ProtoMessage()               {}
func (*Point) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Point) GetX() int32 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *Point) GetY() int32 {
	if m != nil {
		return m.Y
	}
	return 0
}

type Rectangle struct {
	X      int32 `protobuf:"varint,1,opt,name=x" json:"x,omitempty"`
	Y      int32 `protobuf:"varint,2,opt,name=y" json:"y,omitempty"`
	Width  int32 `protobuf:"varint,3,opt,name=width" json:"width,omitempty"`
	Height int32 `protobuf:"varint,4,opt,name=height" json:"height,omitempty"`
}

func (m *Rectangle) Reset()                    { *m = Rectangle{} }
func (m *Rectangle) String() string            { return proto.CompactTextString(m) }
func (*Rectangle) ProtoMessage()               {}
func (*Rectangle) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Rectangle) GetX() int32 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *Rectangle) GetY() int32 {
	if m != nil {
		return m.Y
	}
	return 0
}

func (m *Rectangle) GetWidth() int32 {
	if m != nil {
		return m.Width
	}
	return 0
}

func (m *Rectangle) GetHeight() int32 {
	if m != nil {
		return m.Height
	}
	return 0
}

type Region struct {
	X      int32 `protobuf:"varint,1,opt,name=x" json:"x,omitempty"`
	Y      int32 `protobuf:"varint,2,opt,name=y" json:"y,omitempty"`
	Width  int32 `protobuf:"varint,3,opt,name=width" json:"width,omitempty"`
	Height int32 `protobuf:"varint,4,opt,name=height" json:"height,omitempty"`
}

func (m *Region) Reset()                    { *m = Region{} }
func (m *Region) String() string            { return proto.CompactTextString(m) }
func (*Region) ProtoMessage()               {}
func (*Region) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Region) GetX() int32 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *Region) GetY() int32 {
	if m != nil {
		return m.Y
	}
	return 0
}

func (m *Region) GetWidth() int32 {
	if m != nil {
		return m.Width
	}
	return 0
}

func (m *Region) GetHeight() int32 {
	if m != nil {
		return m.Height
	}
	return 0
}

type Shape struct {
	ContentType string `protobuf:"bytes,1,opt,name=contentType" json:"contentType,omitempty"`
	// this field will be used when contentType == "polygon/points"
	Points []*Point `protobuf:"bytes,2,rep,name=points" json:"points,omitempty"`
	// this field will be used when contentType == "polygon/series"
	// [x1, y1, x2, y2, x3, y3, .... ]
	Series []int32 `protobuf:"varint,3,rep,packed,name=series" json:"series,omitempty"`
}

func (m *Shape) Reset()                    { *m = Shape{} }
func (m *Shape) String() string            { return proto.CompactTextString(m) }
func (*Shape) ProtoMessage()               {}
func (*Shape) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Shape) GetContentType() string {
	if m != nil {
		return m.ContentType
	}
	return ""
}

func (m *Shape) GetPoints() []*Point {
	if m != nil {
		return m.Points
	}
	return nil
}

func (m *Shape) GetSeries() []int32 {
	if m != nil {
		return m.Series
	}
	return nil
}

type Object struct {
	Label string     `protobuf:"bytes,1,opt,name=label" json:"label,omitempty"`
	Box   *Rectangle `protobuf:"bytes,2,opt,name=box" json:"box,omitempty"`
	Shape *Shape     `protobuf:"bytes,3,opt,name=shape" json:"shape,omitempty"`
}

func (m *Object) Reset()                    { *m = Object{} }
func (m *Object) String() string            { return proto.CompactTextString(m) }
func (*Object) ProtoMessage()               {}
func (*Object) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Object) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *Object) GetBox() *Rectangle {
	if m != nil {
		return m.Box
	}
	return nil
}

func (m *Object) GetShape() *Shape {
	if m != nil {
		return m.Shape
	}
	return nil
}

func init() {
	proto.RegisterType((*DetectionResponse)(nil), "serving.DetectionResponse")
	proto.RegisterType((*DetectionRequest)(nil), "serving.DetectionRequest")
	proto.RegisterType((*Point)(nil), "serving.Point")
	proto.RegisterType((*Rectangle)(nil), "serving.Rectangle")
	proto.RegisterType((*Region)(nil), "serving.Region")
	proto.RegisterType((*Shape)(nil), "serving.Shape")
	proto.RegisterType((*Object)(nil), "serving.Object")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ShapeDetection service

type ShapeDetectionClient interface {
	Detect(ctx context.Context, in *DetectionRequest, opts ...grpc.CallOption) (*DetectionResponse, error)
}

type shapeDetectionClient struct {
	cc *grpc.ClientConn
}

func NewShapeDetectionClient(cc *grpc.ClientConn) ShapeDetectionClient {
	return &shapeDetectionClient{cc}
}

func (c *shapeDetectionClient) Detect(ctx context.Context, in *DetectionRequest, opts ...grpc.CallOption) (*DetectionResponse, error) {
	out := new(DetectionResponse)
	err := grpc.Invoke(ctx, "/serving.ShapeDetection/Detect", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ShapeDetection service

type ShapeDetectionServer interface {
	Detect(context.Context, *DetectionRequest) (*DetectionResponse, error)
}

func RegisterShapeDetectionServer(s *grpc.Server, srv ShapeDetectionServer) {
	s.RegisterService(&_ShapeDetection_serviceDesc, srv)
}

func _ShapeDetection_Detect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DetectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShapeDetectionServer).Detect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/serving.ShapeDetection/Detect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShapeDetectionServer).Detect(ctx, req.(*DetectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ShapeDetection_serviceDesc = grpc.ServiceDesc{
	ServiceName: "serving.ShapeDetection",
	HandlerType: (*ShapeDetectionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Detect",
			Handler:    _ShapeDetection_Detect_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "serving.proto",
}

// Client API for ObjectDetection service

type ObjectDetectionClient interface {
	DetectStream(ctx context.Context, opts ...grpc.CallOption) (ObjectDetection_DetectStreamClient, error)
	Detect(ctx context.Context, in *DetectionRequest, opts ...grpc.CallOption) (*DetectionResponse, error)
}

type objectDetectionClient struct {
	cc *grpc.ClientConn
}

func NewObjectDetectionClient(cc *grpc.ClientConn) ObjectDetectionClient {
	return &objectDetectionClient{cc}
}

func (c *objectDetectionClient) DetectStream(ctx context.Context, opts ...grpc.CallOption) (ObjectDetection_DetectStreamClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_ObjectDetection_serviceDesc.Streams[0], c.cc, "/serving.ObjectDetection/DetectStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &objectDetectionDetectStreamClient{stream}
	return x, nil
}

type ObjectDetection_DetectStreamClient interface {
	Send(*DetectionRequest) error
	Recv() (*DetectionResponse, error)
	grpc.ClientStream
}

type objectDetectionDetectStreamClient struct {
	grpc.ClientStream
}

func (x *objectDetectionDetectStreamClient) Send(m *DetectionRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *objectDetectionDetectStreamClient) Recv() (*DetectionResponse, error) {
	m := new(DetectionResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *objectDetectionClient) Detect(ctx context.Context, in *DetectionRequest, opts ...grpc.CallOption) (*DetectionResponse, error) {
	out := new(DetectionResponse)
	err := grpc.Invoke(ctx, "/serving.ObjectDetection/Detect", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ObjectDetection service

type ObjectDetectionServer interface {
	DetectStream(ObjectDetection_DetectStreamServer) error
	Detect(context.Context, *DetectionRequest) (*DetectionResponse, error)
}

func RegisterObjectDetectionServer(s *grpc.Server, srv ObjectDetectionServer) {
	s.RegisterService(&_ObjectDetection_serviceDesc, srv)
}

func _ObjectDetection_DetectStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ObjectDetectionServer).DetectStream(&objectDetectionDetectStreamServer{stream})
}

type ObjectDetection_DetectStreamServer interface {
	Send(*DetectionResponse) error
	Recv() (*DetectionRequest, error)
	grpc.ServerStream
}

type objectDetectionDetectStreamServer struct {
	grpc.ServerStream
}

func (x *objectDetectionDetectStreamServer) Send(m *DetectionResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *objectDetectionDetectStreamServer) Recv() (*DetectionRequest, error) {
	m := new(DetectionRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ObjectDetection_Detect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DetectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectDetectionServer).Detect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/serving.ObjectDetection/Detect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectDetectionServer).Detect(ctx, req.(*DetectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ObjectDetection_serviceDesc = grpc.ServiceDesc{
	ServiceName: "serving.ObjectDetection",
	HandlerType: (*ObjectDetectionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Detect",
			Handler:    _ObjectDetection_Detect_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "DetectStream",
			Handler:       _ObjectDetection_DetectStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "serving.proto",
}

func init() { proto.RegisterFile("serving.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 393 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0xd1, 0xaa, 0xd3, 0x40,
	0x10, 0x75, 0x6f, 0x9a, 0xbd, 0xdc, 0xb9, 0xb5, 0xd5, 0x41, 0x4a, 0xec, 0x53, 0x88, 0x45, 0xaa,
	0x0f, 0x45, 0xe2, 0x17, 0x08, 0xbe, 0xf9, 0xa0, 0x6c, 0xeb, 0x07, 0x24, 0x71, 0x48, 0xb6, 0xb4,
	0xbb, 0x31, 0xbb, 0x6a, 0xfb, 0x39, 0xfe, 0x80, 0xdf, 0x28, 0xd9, 0x4d, 0xd3, 0x20, 0x5a, 0xe4,
	0xd2, 0xb7, 0x9c, 0x39, 0x67, 0x4e, 0xce, 0x30, 0xb3, 0xf0, 0xd8, 0x50, 0xf3, 0x5d, 0xaa, 0x72,
	0x55, 0x37, 0xda, 0x6a, 0xbc, 0xed, 0x60, 0xf2, 0x8b, 0xc1, 0xd3, 0xf7, 0x64, 0xa9, 0xb0, 0x52,
	0x2b, 0x41, 0xa6, 0xd6, 0xca, 0x10, 0x22, 0x8c, 0xec, 0xb1, 0xa6, 0x88, 0xc5, 0x6c, 0x79, 0x27,
	0xdc, 0x37, 0xbe, 0x04, 0x5e, 0x6b, 0xa9, 0xac, 0x89, 0x6e, 0xe2, 0x60, 0x79, 0x9f, 0x4e, 0x56,
	0x27, 0xcb, 0x4f, 0x6d, 0x59, 0x74, 0x2c, 0xbe, 0x82, 0x5b, 0x9d, 0x6f, 0xa9, 0xb0, 0x26, 0x0a,
	0x9c, 0x70, 0xda, 0x0b, 0x3f, 0xba, 0xba, 0x38, 0xf1, 0x98, 0x02, 0x34, 0x54, 0xd8, 0x4c, 0x95,
	0x3b, 0x32, 0xd1, 0xc8, 0xa9, 0xb1, 0x57, 0x8b, 0x13, 0x25, 0x06, 0xaa, 0x64, 0x03, 0x4f, 0x06,
	0x79, 0xbf, 0x7e, 0x23, 0x63, 0xf1, 0x19, 0x84, 0x72, 0x9f, 0x95, 0x3e, 0xef, 0x58, 0x78, 0x80,
	0xaf, 0x81, 0x37, 0x54, 0x4a, 0xad, 0xa2, 0x9b, 0x98, 0xfd, 0xc3, 0xb9, 0x53, 0x24, 0x2f, 0x20,
	0x74, 0x53, 0xe0, 0x18, 0xd8, 0xc1, 0xd9, 0x84, 0x82, 0x1d, 0x5a, 0x74, 0x74, 0xdd, 0xa1, 0x60,
	0xc7, 0xe4, 0x33, 0xdc, 0xf5, 0x9d, 0x97, 0x84, 0x6d, 0x9e, 0x1f, 0xf2, 0x8b, 0xad, 0xa2, 0xc0,
	0x55, 0x3c, 0xc0, 0x19, 0xf0, 0x8a, 0x64, 0x59, 0xd9, 0x68, 0xe4, 0xca, 0x1d, 0x4a, 0x04, 0x70,
	0xe1, 0x52, 0x5c, 0xd1, 0x53, 0x42, 0xb8, 0xae, 0xb2, 0x9a, 0x30, 0x86, 0xfb, 0x42, 0x2b, 0x4b,
	0xca, 0x6e, 0xce, 0x0b, 0x1d, 0x96, 0xfe, 0x7b, 0xaf, 0x33, 0xe0, 0x86, 0x1a, 0x49, 0x7e, 0xad,
	0xa1, 0xe8, 0x50, 0xb2, 0x05, 0xee, 0xf7, 0xda, 0x46, 0xdc, 0x65, 0x39, 0xed, 0xba, 0xbf, 0x78,
	0x80, 0x0b, 0x08, 0x72, 0x7d, 0xb8, 0xb0, 0x83, 0x96, 0xc6, 0x05, 0x84, 0xa6, 0x0d, 0xec, 0xc6,
	0x1b, 0x86, 0x70, 0x63, 0x08, 0x4f, 0xa6, 0x6b, 0x98, 0x38, 0xdc, 0x5f, 0x00, 0xbe, 0x03, 0xee,
	0x01, 0x3e, 0xef, 0x5b, 0xfe, 0xbc, 0x8f, 0xf9, 0xfc, 0x6f, 0x94, 0x3f, 0xf5, 0xe4, 0x51, 0xfa,
	0x93, 0xc1, 0xd4, 0x4f, 0x70, 0xb6, 0xfd, 0x00, 0x63, 0x0f, 0xd6, 0xb6, 0xa1, 0x6c, 0xff, 0x60,
	0xf3, 0x25, 0x7b, 0xc3, 0xae, 0x90, 0x31, 0xe7, 0xee, 0xd9, 0xbe, 0xfd, 0x1d, 0x00, 0x00, 0xff,
	0xff, 0x13, 0xc4, 0x93, 0x85, 0xc7, 0x03, 0x00, 0x00,
}
