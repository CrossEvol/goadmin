
-- name: GetGoadminRoleUser :one
SELECT * FROM `goadmin_role_users`
WHERE created_at = ? LIMIT 1;

-- name: GetGoadminRoleUsers :many
SELECT * FROM `goadmin_role_users`;

-- name: CountGoadminRoleUsers :one
SELECT count(*) FROM `goadmin_role_users`;

-- name: CreateGoadminRoleUser :execresult
INSERT INTO `goadmin_role_users` (
`role_id`,`updated_at`,`user_id`
) VALUES (
? ,? ,? 
);

-- name: UpdateGoadminRoleUser :exec
UPDATE `goadmin_role_users`
SET 
  
  `role_id` = CASE WHEN sqlc.arg('role_id') IS NOT NULL THEN sqlc.arg('role_id') ELSE `role_id` END,
  `updated_at` = CASE WHEN sqlc.arg('updated_at') IS NOT NULL THEN sqlc.arg('updated_at') ELSE `updated_at` END,
  `user_id` = CASE WHEN sqlc.arg('user_id') IS NOT NULL THEN sqlc.arg('user_id') ELSE `user_id` END
WHERE created_at = ?;

-- name: DeleteGoadminRoleUser :exec
DELETE FROM `goadmin_role_users`
WHERE created_at = ?;
