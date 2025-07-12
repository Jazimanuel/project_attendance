-- name: RecordAttendance :one
INSERT INTO attendance (student_id, course_id, check_in)
VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdateCheckOut :exec
UPDATE attendance
SET check_out = $1
WHERE id = $2;

-- name: ListAttendanceByStudent :many
SELECT * FROM attendance
WHERE student_id = $1
ORDER BY created_at DESC;
