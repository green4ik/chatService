-- name: CreateUser :one
INSERT INTO USERS (id,created_at,name,api_key) 
VALUES ($1,$2,$3,
encode(sha256(random()::text::bytea), 'hex')
)
RETURNING *;