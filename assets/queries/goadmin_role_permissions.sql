
-- name: GetGoadminRolePermission :one
SELECT * FROM `goadmin_role_permissions`
WHERE created_at = ? LIMIT 1;

-- name: GetGoadminRolePermissions :many
SELECT * FROM `goadmin_role_permissions`;

-- name: CountGoadminRolePermissions :one
SELECT count(*) FROM `goadmin_role_permissions`;

-- name: CreateGoadminRolePermission :execresult
INSERT INTO `goadmin_role_permissions` (
`permission_id`,`role_id`,`updated_at`
) VALUES (
? ,? ,? 
);

-- name: UpdateGoadminRolePermission :exec
UPDATE `goadmin_role_permissions`
SET 
  
  `permission_id` = CASE WHEN sqlc.arg('permission_id') IS NOT NULL THEN sqlc.arg('permission_id') ELSE `permission_id` END,
  `role_id` = CASE WHEN sqlc.arg('role_id') IS NOT NULL THEN sqlc.arg('role_id') ELSE `role_id` END,
  `updated_at` = CASE WHEN sqlc.arg('updated_at') IS NOT NULL THEN sqlc.arg('updated_at') ELSE `updated_at` END
WHERE created_at = ?;

-- name: DeleteGoadminRolePermission :exec
DELETE FROM `goadmin_role_permissions`
WHERE created_at = ?;
