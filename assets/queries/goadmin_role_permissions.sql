
-- name: GetGoadminRolePermission :one
SELECT * FROM `goadmin_role_permissions`
WHERE role_id = ? LIMIT 1;

-- name: GetGoadminRolePermissions :many
SELECT * FROM `goadmin_role_permissions`;

-- name: CountGoadminRolePermissions :one
SELECT count(*) FROM `goadmin_role_permissions`;

-- name: CreateGoadminRolePermission :execresult
INSERT INTO `goadmin_role_permissions` (
`permission_id`,`created_at`,`updated_at`
) VALUES (
? ,? ,? 
);

-- name: UpdateGoadminRolePermission :exec
UPDATE `goadmin_role_permissions`
SET 
  
  `permission_id` = CASE WHEN sqlc.arg('permission_id') IS NOT NULL THEN sqlc.arg('permission_id') ELSE `permission_id` END,
  `created_at` = CASE WHEN sqlc.arg('created_at') IS NOT NULL THEN sqlc.arg('created_at') ELSE `created_at` END,
  `updated_at` = CASE WHEN sqlc.arg('updated_at') IS NOT NULL THEN sqlc.arg('updated_at') ELSE `updated_at` END
WHERE role_id = ?;

-- name: DeleteGoadminRolePermission :exec
DELETE FROM `goadmin_role_permissions`
WHERE role_id = ?;
