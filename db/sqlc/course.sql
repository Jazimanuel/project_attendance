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

-- name: UpdateCourse :one
UPDATE courses
SET title = $2, instructor = $3
WHERE id = $1
    RETURNING *;

-- name: DeleteCourse :exec
DELETE FROM courses
WHERE id = $1;
