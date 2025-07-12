-- name: CreateCourse :one
INSERT INTO courses (title, instructor)
VALUES ($1, $2)
RETURNING *;

-- name: GetCourseByID :one
SELECT * FROM courses
WHERE id = $1;

-- name: ListCourses :many
SELECT * FROM courses
ORDER BY id;
