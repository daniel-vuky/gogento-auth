-- name: GetAuthorizationRule :one
SELECT *
FROM authorization_rules
WHERE rule_id = $1
FOR NO KEY UPDATE;

-- name: InsertAuthorizationRule :one
INSERT
INTO authorization_rules
    (role_id, is_administrator, permission_code, is_allowed)
VALUES
    ($1, $2, $3, $4)
RETURNING *;

-- name: InsertMultipleAuthorizationRules :many
INSERT
INTO authorization_rules
    (role_id, is_administrator, permission_code, is_allowed)
VALUES
    (UNNEST($1::BIGINT[]), UNNEST($2::BOOLEAN[]), UNNEST($3::varchar(255)[]), UNNEST($4::BOOLEAN[]))
RETURNING *;

-- name: UpdateAuthorizationRule :one
UPDATE authorization_rules
SET
    role_id = COALESCE(sqlc.narg(role_id), role_id),
    is_administrator = COALESCE(sqlc.narg(is_administrator), is_administrator),
    permission_code = COALESCE(sqlc.narg(permission_code), permission_code),
    is_allowed = COALESCE(sqlc.narg(is_allowed), is_allowed)
WHERE rule_id = $1
RETURNING *;

-- name: DeleteAuthorizationRule :one
DELETE FROM authorization_rules
WHERE rule_id = $1
RETURNING *;

-- name: GetAuthorizationRuleByRole :many
SELECT *
FROM authorization_rules
WHERE role_id = $1
FOR NO KEY UPDATE;

-- name: IsAllowed :one
SELECT is_allowed
FROM authorization_rules
WHERE role_id = $1 AND permission_code = $2
FOR NO KEY UPDATE;