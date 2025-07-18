-- name: RecordAttendance :one
INSERT INTO attendance (student_id, course_id, status)
VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdateCheckOut :one
UPDATE attendance
SET check_out = $1
WHERE id = $2
RETURNING *;

-- name: ListAttendanceByStudent :many
SELECT * FROM attendance
WHERE student_id = $1
ORDER BY created_at DESC;
