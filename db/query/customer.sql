-- name: GetCustomer :one
SELECT *
FROM customers
WHERE email = $1
FOR NO KEY UPDATE;

-- name: InsertCustomer :one
INSERT
INTO customers
    (email, firstname, lastname, gender, dob, hashed_password, password_changed_at)
VALUES
    ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: UpdateCustomer :one
UPDATE customers
SET
    firstname = COALESCE(sqlc.narg(firstname), firstname),
    lastname = COALESCE(sqlc.narg(lastname), lastname),
    gender = COALESCE(sqlc.narg(gender), gender),
    dob = COALESCE(sqlc.narg(dob), dob),
    hashed_password = COALESCE(sqlc.narg(hashed_password), hashed_password),
    password_changed_at = COALESCE(sqlc.narg(password_changed_at), password_changed_at)
WHERE email = $1
RETURNING *;