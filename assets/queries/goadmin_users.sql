
-- name: GetGoadminUserByEmail :one
SELECT * FROM `goadmin_users`
WHERE `email` = ? LIMIT 1;

-- name: GetGoadminUser :one
SELECT * FROM `goadmin_users`
WHERE id = ? LIMIT 1;

-- name: GetGoadminUsers :many
SELECT * FROM `goadmin_users`;

-- name: CountGoadminUsers :one
SELECT count(*) FROM `goadmin_users`;

-- name: CreateGoadminUser :execresult
INSERT INTO `goadmin_users` (
`avatar`,`created_at`,`email`,`name`,`password`,`remember_token`,`updated_at`,`username`
) VALUES (
? ,? ,? ,? ,? ,? ,? ,? 
);

-- name: UpdateGoadminUser :exec
UPDATE `goadmin_users`
SET 
  `avatar` = CASE WHEN sqlc.arg('avatar') IS NOT NULL THEN sqlc.arg('avatar') ELSE `avatar` END,
  `created_at` = CASE WHEN sqlc.arg('created_at') IS NOT NULL THEN sqlc.arg('created_at') ELSE `created_at` END,
  `email` = CASE WHEN sqlc.arg('email') IS NOT NULL THEN sqlc.arg('email') ELSE `email` END,
  
  `name` = CASE WHEN sqlc.arg('name') IS NOT NULL THEN sqlc.arg('name') ELSE `name` END,
  `password` = CASE WHEN sqlc.arg('password') IS NOT NULL THEN sqlc.arg('password') ELSE `password` END,
  `remember_token` = CASE WHEN sqlc.arg('remember_token') IS NOT NULL THEN sqlc.arg('remember_token') ELSE `remember_token` END,
  `updated_at` = CASE WHEN sqlc.arg('updated_at') IS NOT NULL THEN sqlc.arg('updated_at') ELSE `updated_at` END,
  `username` = CASE WHEN sqlc.arg('username') IS NOT NULL THEN sqlc.arg('username') ELSE `username` END
WHERE id = ?;

-- name: DeleteGoadminUser :exec
DELETE FROM `goadmin_users`
WHERE id = ?;
