package main

import (
  "context"
  "database/sql"
  "log"
  "net"
  "net/http"

  _ "github.com/lib/pq"
  "google.golang.org/grpc"
  "google.golang.org/grpc/reflection"
  "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

  "Team2_Attendance_Project/internal/service"
  "Team2_Attendance_Project/internal/service/attendancepb"
)

func main() {
  // Connect to PostgreSQL
  connStr := "postgres://postgres:@localhost:5432/attendance_db?sslmode=disable"
  dbConn, err := sql.Open("postgres", connStr)
  if err != nil {
    log.Fatalf("❌ cannot connect to DB: %v", err)
  }
  if err := dbConn.Ping(); err != nil {
    log.Fatalf("❌ database ping failed: %v", err)
  }
  log.Println("✅ Connected to PostgreSQL!")

  // Create gRPC server
  grpcServer := grpc.NewServer()
  attendancepb.RegisterAttendanceServiceServer(grpcServer, service.NewServer(dbConn))
  reflection.Register(grpcServer)

  // Start gRPC server in a goroutine
  go func() {
    listener, err := net.Listen("tcp", ":50051")
    if err != nil {
      log.Fatalf("❌ failed to listen on port 50051: %v", err)
    }
    log.Println("🚀 gRPC server is live at port 50051")
    if err := grpcServer.Serve(listener); err != nil {
      log.Fatalf("❌ failed to serve gRPC: %v", err)
    }
  }()

  // Start REST gateway in a goroutine
  go func() {
    ctx := context.Background()
    mux := runtime.NewServeMux()
    opts := []grpc.DialOption{grpc.WithInsecure()}

    err := attendancepb.RegisterAttendanceServiceHandlerFromEndpoint(ctx, mux, "localhost:50051", opts)
    if err != nil {
      log.Fatalf("❌ failed to register REST gateway: %v", err)
    }

    log.Println("🌐 REST gateway is live at port 8080")
    if err := http.ListenAndServe(":8080", mux); err != nil {
      log.Fatalf("❌ failed to serve REST gateway: %v", err)
    }
  }()

  // Block main forever
  select {}
}
