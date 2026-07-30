-- name: Listusers :many
SELECT * FROM users
ORDER BY name;

-- name: Detailusers :one
SELECT * FROM users
WHERE id = $1
ORDER BY name;