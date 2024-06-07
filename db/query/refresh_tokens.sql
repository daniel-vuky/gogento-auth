-- name: GetRefreshToken :one
SELECT *
FROM refresh_tokens
WHERE customer_id = $1;

-- name: InsertRefreshToken :one
INSERT INTO refresh_tokens
    (customer_id, refresh_token, user_agent, client_ip, is_blocked, expired_at)
VALUES
    ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: UpdateRefreshToken :one
UPDATE refresh_tokens
SET
    refresh_token = COALESCE(sqlc.narg(refresh_token), refresh_token),
    user_agent = COALESCE(sqlc.narg(user_agent), user_agent),
    client_ip = COALESCE(sqlc.narg(client_ip), client_ip),
    is_blocked = COALESCE(sqlc.narg(is_blocked), is_blocked),
    expired_at = COALESCE(sqlc.narg(expired_at), expired_at)
WHERE customer_id = $1
RETURNING *;

-- name: DeleteRefreshToken :one
DELETE FROM refresh_tokens
WHERE customer_id = $1
RETURNING *;