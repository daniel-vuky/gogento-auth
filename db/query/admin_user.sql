-- name: GetAdminUser :one
SELECT *
FROM admin_users
WHERE email = $1
FOR NO KEY UPDATE;

-- name: InsertAdminUser :one
INSERT
INTO admin_users
    (email, role_id, firstname, lastname, hashed_password, is_active, lock_expires)
VALUES
    ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: UpdateAdminUser :one
UPDATE admin_users
SET
    role_id = COALESCE(sqlc.narg(role_id), role_id),
    firstname = COALESCE(sqlc.narg(firstname), firstname),
    lastname = COALESCE(sqlc.narg(lastname), lastname),
    hashed_password = COALESCE(sqlc.narg(hashed_password), hashed_password),
    is_active = COALESCE(sqlc.narg(is_active), is_active),
    lock_expires = COALESCE(sqlc.narg(lock_expires), lock_expires)
WHERE email = $1
RETURNING *;

-- name: DeleteAdminUser :one
DELETE FROM admin_users
WHERE email = $1
RETURNING *;