package service

import (
	"context"
	"database/sql"
	"fmt"

	"Team2_Attendance_Project/internal/db"
	"Team2_Attendance_Project/internal/service/attendancepb"
)

type Server struct {
	attendancepb.UnimplementedAttendanceServiceServer
	queries *db.Queries
}

func NewServer(dbConn *sql.DB) *Server {
	return &Server{
		queries: db.New(dbConn),
	}
}

func (s *Server) CreateStudent(ctx context.Context, req *attendancepb.CreateStudentRequest) (*attendancepb.StudentResponse, error) {
	student, err := s.queries.CreateStudent(ctx, db.CreateStudentParams{
		Name:  req.Name,
		Email: req.Email,
	})
	if err != nil {
		return nil, err
	}

	var createdAt string
	if student.CreatedAt.Valid {
		createdAt = student.CreatedAt.Time.Format("2006-01-02 15:04:05")
	}

	return &attendancepb.StudentResponse{
		Id:        int32(student.ID),
		Name:      student.Name,
		Email:     student.Email,
		CreatedAt: createdAt,
	}, nil
}

func (s *Server) GetStudentByID(ctx context.Context, req *attendancepb.GetStudentRequest) (*attendancepb.StudentResponse, error) {
	student, err := s.queries.GetStudentByID(ctx, int32(req.Id))
	if err != nil {
		return nil, err
	}

	var createdAt string
	if student.CreatedAt.Valid {
		createdAt = student.CreatedAt.Time.Format("2006-01-02 15:04:05")
	}

	return &attendancepb.StudentResponse{
		Id:        int32(student.ID),
		Name:      student.Name,
		Email:     student.Email,
		CreatedAt: createdAt,
	}, nil
}

func (s *Server) ListStudents(ctx context.Context, _ *attendancepb.Empty) (*attendancepb.StudentList, error) {
	students, err := s.queries.ListStudents(ctx)
	if err != nil {
		return nil, err
	}

	var protoStudents []*attendancepb.StudentResponse
	for _, student := range students {
		var createdAt string
		if student.CreatedAt.Valid {
			createdAt = student.CreatedAt.Time.Format("2006-01-02 15:04:05")
		}

		protoStudents = append(protoStudents, &attendancepb.StudentResponse{
			Id:        int32(student.ID),
			Name:      student.Name,
			Email:     student.Email,
			CreatedAt: createdAt,
		})
	}

	return &attendancepb.StudentList{
		Students: protoStudents,
	}, nil
}

func (s *Server) UpdateStudent(ctx context.Context, req *attendancepb.UpdateStudentRequest) (*attendancepb.StudentResponse, error) {
	student, err := s.queries.UpdateStudent(ctx, db.UpdateStudentParams{
		ID:    req.Id,
		Name:  req.Name,
		Email: req.Email,
	})

	if err != nil {
		return nil, fmt.Errorf("update student error")
	}

	var createdAt string
	if student.CreatedAt.Valid {
		createdAt = student.CreatedAt.Time.Format("2006-01-02 15:04:05")
	}

	return &attendancepb.StudentResponse{
		Id:        student.ID,
		Name:      student.Name,
		Email:     student.Email,
		CreatedAt: createdAt,
	}, nil
}

func (s *Server) DeleteStudent(ctx context.Context, req *attendancepb.GetStudentRequest) error {
	err := s.queries.DeleteStudent(ctx, req.Id)

	if err != nil {
		return fmt.Errorf("delete student error")

	}
	return nil
}

func (s *Server) CreateCourse(ctx context.Context, req *attendancepb.CreateCourseRequest) (*attendancepb.CourseResponse, error) {
	course, err := s.queries.CreateCourse(ctx, db.CreateCourseParams{
		Title:      req.Title,
		Instructor: req.Instructor,
	})

	if err != nil {
		return nil, err
	}

	var createdAt string
	if course.CreatedAt.Valid {
		createdAt = course.CreatedAt.Time.Format("2006-01-02 15:04:05")
	}

	return &attendancepb.CourseResponse{
		Id:         int32(course.ID),
		Title:      course.Title,
		Instructor: req.Instructor,
		CreatedAt:  createdAt,
	}, nil
}

func (s *Server) GetCourseByID(ctx context.Context, req *attendancepb.GetCourseRequest) (*attendancepb.CourseResponse, error) {
	course, err := s.queries.GetCourseByID(ctx, int32(req.Id))
	if err != nil {
		return nil, err
	}

	var createdAt string
	if course.CreatedAt.Valid {
		createdAt = course.CreatedAt.Time.Format("2006-01-02 15:04:05")
	}

	return &attendancepb.CourseResponse{
		Id:         int32(course.ID),
		Title:      course.Title,
		Instructor: course.Instructor,
		CreatedAt:  createdAt,
	}, nil
}

func (s *Server) ListCourses(ctx context.Context, _ *attendancepb.Empty) (*attendancepb.CourseList, error) {
	courses, err := s.queries.ListCourses(ctx)
	if err != nil {
		return nil, err
	}

	var protoCourses []*attendancepb.CourseResponse
	for _, course := range courses {
		var createdAt string
		if course.CreatedAt.Valid {
			createdAt = course.CreatedAt.Time.Format("2006-01-02 15:04:05")
		}

		protoCourses = append(protoCourses, &attendancepb.CourseResponse{
			Id:         int32(course.ID),
			Title:      course.Title,
			Instructor: course.Instructor,
			CreatedAt:  createdAt,
		})
	}

	return &attendancepb.CourseList{
		Courses: protoCourses,
	}, nil
}

func (s *Server) UpdateCourse(ctx context.Context, req *attendancepb.UpdateCourseRequest) (*attendancepb.CourseResponse, error) {
	course, err := s.queries.UpdateCourse(ctx, db.UpdateCourseParams{
		ID:         req.Id,
		Title:      req.Title,
		Instructor: req.Instructor,
	})

	if err != nil {
		return nil, fmt.Errorf("update course error")
	}

	var createdAt string
	if course.CreatedAt.Valid {
		createdAt = course.CreatedAt.Time.Format("2006-01-02 15:04:05")
	}

	return &attendancepb.CourseResponse{
		Id:         course.ID,
		Title:      course.Title,
		Instructor: course.Instructor,
		CreatedAt:  createdAt,
	}, nil
}

func (s *Server) DeleteCourse(ctx context.Context, req *attendancepb.GetCourseRequest) error {
	err := s.queries.DeleteCourse(ctx, req.Id)

	if err != nil {
		return fmt.Errorf("delete course error")

	}
	return nil
}

func (s *Server) CreateAdmin(ctx context.Context, req *attendancepb.CreateAdminRequest) (*attendancepb.AdminResponse, error) {
	admin, err := s.queries.CreateAdmin(ctx, db.CreateAdminParams{
		Name:  req.Name,
		Email: req.Email,
	})

	if err != nil {
		return nil, err
	}

	var createdAt string
	if admin.CreatedAt.Valid {
		createdAt = admin.CreatedAt.Time.Format("2006-01-02 15:04:05")
	}

	return &attendancepb.AdminResponse{
		Id:        int32(admin.ID),
		Name:      admin.Name,
		Email:     req.Email,
		CreatedAt: createdAt,
	}, nil
}

func (s *Server) GetAdminByID(ctx context.Context, req *attendancepb.GetAdminByEmail) (*attendancepb.AdminResponse, error) {
	admin, err := s.queries.GetAdminByEmail(ctx, req.Email)

	if err != nil {
		return nil, err
	}

	var createdAt string
	if admin.CreatedAt.Valid {
		createdAt = admin.CreatedAt.Time.Format("2006-01-02 15:04:05")
	}

	return &attendancepb.AdminResponse{
		Id:        int32(admin.ID),
		Name:      admin.Name,
		Email:     admin.Email,
		CreatedAt: createdAt,
	}, nil
}

func (s *Server) ListAdmins(ctx context.Context, _ *attendancepb.Empty) (*attendancepb.AdminList, error) {
	admins, err := s.queries.ListAdmins(ctx)
	if err != nil {
		return nil, err
	}

	var protoAdmins []*attendancepb.AdminResponse
	for _, admin := range admins {
		var createdAt string
		if admin.CreatedAt.Valid {
			createdAt = admin.CreatedAt.Time.Format("2006-01-02 15:04:05")
		}

		protoAdmins = append(protoAdmins, &attendancepb.AdminResponse{
			Id:        int32(admin.ID),
			Name:      admin.Name,
			Email:     admin.Email,
			CreatedAt: createdAt,
		})
	}

	return &attendancepb.AdminList{
		Admins: protoAdmins,
	}, nil
}

func (s *Server) UpdateAdmin(ctx context.Context, req *attendancepb.UpdateAdminRequest) (*attendancepb.AdminResponse, error) {
	admin, err := s.queries.UpdateAdmin(ctx, db.UpdateAdminParams{
		ID:    req.Id,
		Name:  req.Name,
		Email: req.Email,
	})

	if err != nil {
		return nil, fmt.Errorf("update admin error")
	}

	var createdAt string
	if admin.CreatedAt.Valid {
		createdAt = admin.CreatedAt.Time.Format("2006-01-02 15:04:05")
	}

	return &attendancepb.AdminResponse{
		Id:        admin.ID,
		Name:      admin.Name,
		Email:     admin.Email,
		CreatedAt: createdAt,
	}, nil
}

func (s *Server) DeleteAdmin(ctx context.Context, req *attendancepb.GetAdminByEmail) error {
	err := s.queries.DeleteAdmin(ctx, req.Email)

	if err != nil {
		return fmt.Errorf("delete admin error")

	}
	return nil
}
