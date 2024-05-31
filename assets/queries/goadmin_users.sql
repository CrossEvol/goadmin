
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
`username`,`password`,`name`,`avatar`,`remember_token`,`created_at`,`updated_at`,`email`
) VALUES (
? ,? ,? ,? ,? ,? ,? ,? 
);

-- name: UpdateGoadminUser :exec
UPDATE `goadmin_users`
SET 
  
  `username` = CASE WHEN sqlc.arg('username') IS NOT NULL THEN sqlc.arg('username') ELSE `username` END,
  `password` = CASE WHEN sqlc.arg('password') IS NOT NULL THEN sqlc.arg('password') ELSE `password` END,
  `name` = CASE WHEN sqlc.arg('name') IS NOT NULL THEN sqlc.arg('name') ELSE `name` END,
  `avatar` = CASE WHEN sqlc.arg('avatar') IS NOT NULL THEN sqlc.arg('avatar') ELSE `avatar` END,
  `remember_token` = CASE WHEN sqlc.arg('remember_token') IS NOT NULL THEN sqlc.arg('remember_token') ELSE `remember_token` END,
  `created_at` = CASE WHEN sqlc.arg('created_at') IS NOT NULL THEN sqlc.arg('created_at') ELSE `created_at` END,
  `updated_at` = CASE WHEN sqlc.arg('updated_at') IS NOT NULL THEN sqlc.arg('updated_at') ELSE `updated_at` END,
  `email` = CASE WHEN sqlc.arg('email') IS NOT NULL THEN sqlc.arg('email') ELSE `email` END
WHERE id = ?;

-- name: DeleteGoadminUser :exec
DELETE FROM `goadmin_users`
WHERE id = ?;
