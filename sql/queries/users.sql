-- name: CreateUser :one
INSERT INTO USERS (id,created_at,name) 
VALUES ($1,$2,$3)
RETURNING *;