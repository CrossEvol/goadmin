
-- name: GetGoadminUserPermission :one
SELECT * FROM `goadmin_user_permissions`
WHERE user_id = ? LIMIT 1;

-- name: GetGoadminUserPermissions :many
SELECT * FROM `goadmin_user_permissions`;

-- name: CountGoadminUserPermissions :one
SELECT count(*) FROM `goadmin_user_permissions`;

-- name: CreateGoadminUserPermission :execresult
INSERT INTO `goadmin_user_permissions` (
`permission_id`,`created_at`,`updated_at`
) VALUES (
? ,? ,? 
);

-- name: UpdateGoadminUserPermission :exec
UPDATE `goadmin_user_permissions`
SET 
  
  `permission_id` = CASE WHEN sqlc.arg('permission_id') IS NOT NULL THEN sqlc.arg('permission_id') ELSE `permission_id` END,
  `created_at` = CASE WHEN sqlc.arg('created_at') IS NOT NULL THEN sqlc.arg('created_at') ELSE `created_at` END,
  `updated_at` = CASE WHEN sqlc.arg('updated_at') IS NOT NULL THEN sqlc.arg('updated_at') ELSE `updated_at` END
WHERE user_id = ?;

-- name: DeleteGoadminUserPermission :exec
DELETE FROM `goadmin_user_permissions`
WHERE user_id = ?;
