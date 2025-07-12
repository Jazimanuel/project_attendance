-- name: CreateStudent :one
INSERT INTO students (name, email)
VALUES ($1, $2)
RETURNING *;

-- name: GetStudentByID :one
SELECT * FROM students
WHERE id = $1;

-- name: ListStudents :many
SELECT * FROM students
ORDER BY id;
