-- name: GetAuthorizationRole :one
SELECT *
FROM authorization_roles
WHERE role_id = $1
FOR NO KEY UPDATE;

-- name: InsertAuthorizationRole :one
INSERT
INTO authorization_roles
    (role_name, description)
VALUES
    ($1, $2)
RETURNING *;

-- name: UpdateAuthorizationRole :one
UPDATE authorization_roles
SET
    role_name = COALESCE(sqlc.narg(role_name), role_name),
    description = COALESCE(sqlc.narg(description), description)
WHERE role_id = $1
RETURNING *;

-- name: DeleteAuthorizationRole :one
DELETE FROM authorization_roles
WHERE role_id = $1
RETURNING *;