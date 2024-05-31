
-- name: GetGoadminRoleUser :one
SELECT * FROM `goadmin_role_users`
WHERE role_id = ? LIMIT 1;

-- name: GetGoadminRoleUsers :many
SELECT * FROM `goadmin_role_users`;

-- name: CountGoadminRoleUsers :one
SELECT count(*) FROM `goadmin_role_users`;

-- name: CreateGoadminRoleUser :execresult
INSERT INTO `goadmin_role_users` (
`user_id`,`created_at`,`updated_at`
) VALUES (
? ,? ,? 
);

-- name: UpdateGoadminRoleUser :exec
UPDATE `goadmin_role_users`
SET 
  
  `user_id` = CASE WHEN sqlc.arg('user_id') IS NOT NULL THEN sqlc.arg('user_id') ELSE `user_id` END,
  `created_at` = CASE WHEN sqlc.arg('created_at') IS NOT NULL THEN sqlc.arg('created_at') ELSE `created_at` END,
  `updated_at` = CASE WHEN sqlc.arg('updated_at') IS NOT NULL THEN sqlc.arg('updated_at') ELSE `updated_at` END
WHERE role_id = ?;

-- name: DeleteGoadminRoleUser :exec
DELETE FROM `goadmin_role_users`
WHERE role_id = ?;
