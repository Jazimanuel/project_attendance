PROJECT ATTENDANCE INITIAL TEMPLATE

pg_ctl -D "C:\pgsql\data" start -l logfile
psql -U postgres
CREATE DATABASE attendance_db;

mkdir Team2_Attendance_Project
cd Team2_Attendance_Project
mkdir db db\migrations db\sqlc proto internal internal\db internal\service

migrate create -ext sql -dir db/migrations -seq init_schema

#buat up db
migrate -source "file://C:/Users/Lenovo/Documents/PROJECT GOLANG/Team2_Attendance_Project/db/migrations" -database "postgres://postgres:@localhost:5432/attendance_db?sslmode=disable" up
#buat down
migrate -source "file://C:/Users/Lenovo/Documents/PROJECT GOLANG/Team2_Attendance_Project/db/migrations" -database "postgres://postgres:@localhost:5432/attendance_db?sslmode=disable" down

protoc --go_out=. --go-grpc_out=. proto/attendance.proto

protoc -I=proto -I=./googleapis --go_out=internal/service/attendancepb --go_opt=paths=source_relative --go-grpc_out=internal/service/attendancepb --go-grpc_opt=paths=source_relative --grpc-gateway_out=internal/service/attendancepb --grpc-gateway_opt=paths=source_relative proto/attendance.proto


go run main.go

grpcui -plaintext -port 9090 localhost:50051
