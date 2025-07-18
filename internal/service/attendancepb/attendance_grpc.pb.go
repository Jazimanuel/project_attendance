// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v6.31.1
// source: attendance.proto

package attendancepb

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
	StudentService_CreateStudent_FullMethodName  = "/attendance.StudentService/CreateStudent"
	StudentService_GetStudentByID_FullMethodName = "/attendance.StudentService/GetStudentByID"
	StudentService_ListStudents_FullMethodName   = "/attendance.StudentService/ListStudents"
)

// StudentServiceClient is the client API for StudentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StudentServiceClient interface {
	CreateStudent(ctx context.Context, in *CreateStudentRequest, opts ...grpc.CallOption) (*StudentResponse, error)
	GetStudentByID(ctx context.Context, in *GetStudentRequest, opts ...grpc.CallOption) (*StudentResponse, error)
	ListStudents(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*StudentList, error)
}

type studentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStudentServiceClient(cc grpc.ClientConnInterface) StudentServiceClient {
	return &studentServiceClient{cc}
}

func (c *studentServiceClient) CreateStudent(ctx context.Context, in *CreateStudentRequest, opts ...grpc.CallOption) (*StudentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StudentResponse)
	err := c.cc.Invoke(ctx, StudentService_CreateStudent_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentServiceClient) GetStudentByID(ctx context.Context, in *GetStudentRequest, opts ...grpc.CallOption) (*StudentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StudentResponse)
	err := c.cc.Invoke(ctx, StudentService_GetStudentByID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentServiceClient) ListStudents(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*StudentList, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StudentList)
	err := c.cc.Invoke(ctx, StudentService_ListStudents_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StudentServiceServer is the server API for StudentService service.
// All implementations must embed UnimplementedStudentServiceServer
// for forward compatibility.
type StudentServiceServer interface {
	CreateStudent(context.Context, *CreateStudentRequest) (*StudentResponse, error)
	GetStudentByID(context.Context, *GetStudentRequest) (*StudentResponse, error)
	ListStudents(context.Context, *Empty) (*StudentList, error)
	mustEmbedUnimplementedStudentServiceServer()
}

// UnimplementedStudentServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedStudentServiceServer struct{}

func (UnimplementedStudentServiceServer) CreateStudent(context.Context, *CreateStudentRequest) (*StudentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateStudent not implemented")
}
func (UnimplementedStudentServiceServer) GetStudentByID(context.Context, *GetStudentRequest) (*StudentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStudentByID not implemented")
}
func (UnimplementedStudentServiceServer) ListStudents(context.Context, *Empty) (*StudentList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListStudents not implemented")
}
func (UnimplementedStudentServiceServer) mustEmbedUnimplementedStudentServiceServer() {}
func (UnimplementedStudentServiceServer) testEmbeddedByValue()                        {}

// UnsafeStudentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StudentServiceServer will
// result in compilation errors.
type UnsafeStudentServiceServer interface {
	mustEmbedUnimplementedStudentServiceServer()
}

func RegisterStudentServiceServer(s grpc.ServiceRegistrar, srv StudentServiceServer) {
	// If the following call pancis, it indicates UnimplementedStudentServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&StudentService_ServiceDesc, srv)
}

func _StudentService_CreateStudent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateStudentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServiceServer).CreateStudent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StudentService_CreateStudent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServiceServer).CreateStudent(ctx, req.(*CreateStudentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudentService_GetStudentByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStudentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServiceServer).GetStudentByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StudentService_GetStudentByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServiceServer).GetStudentByID(ctx, req.(*GetStudentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudentService_ListStudents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServiceServer).ListStudents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StudentService_ListStudents_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServiceServer).ListStudents(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// StudentService_ServiceDesc is the grpc.ServiceDesc for StudentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StudentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "attendance.StudentService",
	HandlerType: (*StudentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateStudent",
			Handler:    _StudentService_CreateStudent_Handler,
		},
		{
			MethodName: "GetStudentByID",
			Handler:    _StudentService_GetStudentByID_Handler,
		},
		{
			MethodName: "ListStudents",
			Handler:    _StudentService_ListStudents_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "attendance.proto",
}

const (
	CourseService_CreateCourse_FullMethodName = "/attendance.CourseService/CreateCourse"
	CourseService_ListCourses_FullMethodName  = "/attendance.CourseService/ListCourses"
)

// CourseServiceClient is the client API for CourseService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CourseServiceClient interface {
	CreateCourse(ctx context.Context, in *CreateCourseRequest, opts ...grpc.CallOption) (*CourseResponse, error)
	ListCourses(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*CourseList, error)
}

type courseServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCourseServiceClient(cc grpc.ClientConnInterface) CourseServiceClient {
	return &courseServiceClient{cc}
}

func (c *courseServiceClient) CreateCourse(ctx context.Context, in *CreateCourseRequest, opts ...grpc.CallOption) (*CourseResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CourseResponse)
	err := c.cc.Invoke(ctx, CourseService_CreateCourse_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courseServiceClient) ListCourses(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*CourseList, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CourseList)
	err := c.cc.Invoke(ctx, CourseService_ListCourses_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CourseServiceServer is the server API for CourseService service.
// All implementations must embed UnimplementedCourseServiceServer
// for forward compatibility.
type CourseServiceServer interface {
	CreateCourse(context.Context, *CreateCourseRequest) (*CourseResponse, error)
	ListCourses(context.Context, *Empty) (*CourseList, error)
	mustEmbedUnimplementedCourseServiceServer()
}

// UnimplementedCourseServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCourseServiceServer struct{}

func (UnimplementedCourseServiceServer) CreateCourse(context.Context, *CreateCourseRequest) (*CourseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCourse not implemented")
}
func (UnimplementedCourseServiceServer) ListCourses(context.Context, *Empty) (*CourseList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCourses not implemented")
}
func (UnimplementedCourseServiceServer) mustEmbedUnimplementedCourseServiceServer() {}
func (UnimplementedCourseServiceServer) testEmbeddedByValue()                       {}

// UnsafeCourseServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CourseServiceServer will
// result in compilation errors.
type UnsafeCourseServiceServer interface {
	mustEmbedUnimplementedCourseServiceServer()
}

func RegisterCourseServiceServer(s grpc.ServiceRegistrar, srv CourseServiceServer) {
	// If the following call pancis, it indicates UnimplementedCourseServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CourseService_ServiceDesc, srv)
}

func _CourseService_CreateCourse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCourseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourseServiceServer).CreateCourse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CourseService_CreateCourse_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourseServiceServer).CreateCourse(ctx, req.(*CreateCourseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CourseService_ListCourses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourseServiceServer).ListCourses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CourseService_ListCourses_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourseServiceServer).ListCourses(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// CourseService_ServiceDesc is the grpc.ServiceDesc for CourseService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CourseService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "attendance.CourseService",
	HandlerType: (*CourseServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCourse",
			Handler:    _CourseService_CreateCourse_Handler,
		},
		{
			MethodName: "ListCourses",
			Handler:    _CourseService_ListCourses_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "attendance.proto",
}

const (
	AdminService_CreateAdmin_FullMethodName = "/attendance.AdminService/CreateAdmin"
	AdminService_LoginAdmin_FullMethodName  = "/attendance.AdminService/LoginAdmin"
)

// AdminServiceClient is the client API for AdminService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdminServiceClient interface {
	CreateAdmin(ctx context.Context, in *CreateAdminRequest, opts ...grpc.CallOption) (*AdminResponse, error)
	LoginAdmin(ctx context.Context, in *AdminLoginRequest, opts ...grpc.CallOption) (*AdminResponse, error)
}

type adminServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAdminServiceClient(cc grpc.ClientConnInterface) AdminServiceClient {
	return &adminServiceClient{cc}
}

func (c *adminServiceClient) CreateAdmin(ctx context.Context, in *CreateAdminRequest, opts ...grpc.CallOption) (*AdminResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AdminResponse)
	err := c.cc.Invoke(ctx, AdminService_CreateAdmin_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) LoginAdmin(ctx context.Context, in *AdminLoginRequest, opts ...grpc.CallOption) (*AdminResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AdminResponse)
	err := c.cc.Invoke(ctx, AdminService_LoginAdmin_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdminServiceServer is the server API for AdminService service.
// All implementations must embed UnimplementedAdminServiceServer
// for forward compatibility.
type AdminServiceServer interface {
	CreateAdmin(context.Context, *CreateAdminRequest) (*AdminResponse, error)
	LoginAdmin(context.Context, *AdminLoginRequest) (*AdminResponse, error)
	mustEmbedUnimplementedAdminServiceServer()
}

// UnimplementedAdminServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAdminServiceServer struct{}

func (UnimplementedAdminServiceServer) CreateAdmin(context.Context, *CreateAdminRequest) (*AdminResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAdmin not implemented")
}
func (UnimplementedAdminServiceServer) LoginAdmin(context.Context, *AdminLoginRequest) (*AdminResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginAdmin not implemented")
}
func (UnimplementedAdminServiceServer) mustEmbedUnimplementedAdminServiceServer() {}
func (UnimplementedAdminServiceServer) testEmbeddedByValue()                      {}

// UnsafeAdminServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdminServiceServer will
// result in compilation errors.
type UnsafeAdminServiceServer interface {
	mustEmbedUnimplementedAdminServiceServer()
}

func RegisterAdminServiceServer(s grpc.ServiceRegistrar, srv AdminServiceServer) {
	// If the following call pancis, it indicates UnimplementedAdminServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AdminService_ServiceDesc, srv)
}

func _AdminService_CreateAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAdminRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).CreateAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_CreateAdmin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).CreateAdmin(ctx, req.(*CreateAdminRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_LoginAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).LoginAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_LoginAdmin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).LoginAdmin(ctx, req.(*AdminLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AdminService_ServiceDesc is the grpc.ServiceDesc for AdminService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AdminService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "attendance.AdminService",
	HandlerType: (*AdminServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAdmin",
			Handler:    _AdminService_CreateAdmin_Handler,
		},
		{
			MethodName: "LoginAdmin",
			Handler:    _AdminService_LoginAdmin_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "attendance.proto",
}

const (
	AttendanceService_MarkAttendance_FullMethodName          = "/attendance.AttendanceService/MarkAttendance"
	AttendanceService_GetAttendanceForStudent_FullMethodName = "/attendance.AttendanceService/GetAttendanceForStudent"
)

// AttendanceServiceClient is the client API for AttendanceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AttendanceServiceClient interface {
	MarkAttendance(ctx context.Context, in *MarkAttendanceRequest, opts ...grpc.CallOption) (*AttendanceResponse, error)
	GetAttendanceForStudent(ctx context.Context, in *StudentAttendanceRequest, opts ...grpc.CallOption) (*AttendanceList, error)
}

type attendanceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAttendanceServiceClient(cc grpc.ClientConnInterface) AttendanceServiceClient {
	return &attendanceServiceClient{cc}
}

func (c *attendanceServiceClient) MarkAttendance(ctx context.Context, in *MarkAttendanceRequest, opts ...grpc.CallOption) (*AttendanceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AttendanceResponse)
	err := c.cc.Invoke(ctx, AttendanceService_MarkAttendance_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attendanceServiceClient) GetAttendanceForStudent(ctx context.Context, in *StudentAttendanceRequest, opts ...grpc.CallOption) (*AttendanceList, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AttendanceList)
	err := c.cc.Invoke(ctx, AttendanceService_GetAttendanceForStudent_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AttendanceServiceServer is the server API for AttendanceService service.
// All implementations must embed UnimplementedAttendanceServiceServer
// for forward compatibility.
type AttendanceServiceServer interface {
	MarkAttendance(context.Context, *MarkAttendanceRequest) (*AttendanceResponse, error)
	GetAttendanceForStudent(context.Context, *StudentAttendanceRequest) (*AttendanceList, error)
	mustEmbedUnimplementedAttendanceServiceServer()
}

// UnimplementedAttendanceServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAttendanceServiceServer struct{}

func (UnimplementedAttendanceServiceServer) MarkAttendance(context.Context, *MarkAttendanceRequest) (*AttendanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MarkAttendance not implemented")
}
func (UnimplementedAttendanceServiceServer) GetAttendanceForStudent(context.Context, *StudentAttendanceRequest) (*AttendanceList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAttendanceForStudent not implemented")
}
func (UnimplementedAttendanceServiceServer) mustEmbedUnimplementedAttendanceServiceServer() {}
func (UnimplementedAttendanceServiceServer) testEmbeddedByValue()                           {}

// UnsafeAttendanceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AttendanceServiceServer will
// result in compilation errors.
type UnsafeAttendanceServiceServer interface {
	mustEmbedUnimplementedAttendanceServiceServer()
}

func RegisterAttendanceServiceServer(s grpc.ServiceRegistrar, srv AttendanceServiceServer) {
	// If the following call pancis, it indicates UnimplementedAttendanceServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AttendanceService_ServiceDesc, srv)
}

func _AttendanceService_MarkAttendance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MarkAttendanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttendanceServiceServer).MarkAttendance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AttendanceService_MarkAttendance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttendanceServiceServer).MarkAttendance(ctx, req.(*MarkAttendanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AttendanceService_GetAttendanceForStudent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StudentAttendanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttendanceServiceServer).GetAttendanceForStudent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AttendanceService_GetAttendanceForStudent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttendanceServiceServer).GetAttendanceForStudent(ctx, req.(*StudentAttendanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AttendanceService_ServiceDesc is the grpc.ServiceDesc for AttendanceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AttendanceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "attendance.AttendanceService",
	HandlerType: (*AttendanceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MarkAttendance",
			Handler:    _AttendanceService_MarkAttendance_Handler,
		},
		{
			MethodName: "GetAttendanceForStudent",
			Handler:    _AttendanceService_GetAttendanceForStudent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "attendance.proto",
}
