-- name: CreateAdmin :one
INSERT INTO admins (name, email, role)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetAdminByEmail :one
SELECT * FROM admins
WHERE email = $1;

-- name: ListAdmins :many
SELECT * FROM admins
ORDER BY id;