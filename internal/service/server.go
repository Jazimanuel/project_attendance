package service

import (
  "context"
  "database/sql"

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
