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

-- name: UpdateStudent :one
UPDATE students
SET name = $2, email = $3
WHERE id = $1
RETURNING *;

-- name: DeleteStudent :exec
DELETE FROM students
WHERE id = $1;
