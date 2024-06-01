
-- name: GetGoadminUserPermission :one
SELECT * FROM `goadmin_user_permissions`
WHERE created_at = ? LIMIT 1;

-- name: GetGoadminUserPermissions :many
SELECT * FROM `goadmin_user_permissions`;

-- name: CountGoadminUserPermissions :one
SELECT count(*) FROM `goadmin_user_permissions`;

-- name: CreateGoadminUserPermission :execresult
INSERT INTO `goadmin_user_permissions` (
`permission_id`,`updated_at`,`user_id`
) VALUES (
? ,? ,? 
);

-- name: UpdateGoadminUserPermission :exec
UPDATE `goadmin_user_permissions`
SET 
  
  `permission_id` = CASE WHEN sqlc.arg('permission_id') IS NOT NULL THEN sqlc.arg('permission_id') ELSE `permission_id` END,
  `updated_at` = CASE WHEN sqlc.arg('updated_at') IS NOT NULL THEN sqlc.arg('updated_at') ELSE `updated_at` END,
  `user_id` = CASE WHEN sqlc.arg('user_id') IS NOT NULL THEN sqlc.arg('user_id') ELSE `user_id` END
WHERE created_at = ?;

-- name: DeleteGoadminUserPermission :exec
DELETE FROM `goadmin_user_permissions`
WHERE created_at = ?;
