
-- name: GetGoadminRoleMenu :one
SELECT * FROM `goadmin_role_menu`
WHERE created_at = ? LIMIT 1;

-- name: GetGoadminRoleMenus :many
SELECT * FROM `goadmin_role_menu`;

-- name: CountGoadminRoleMenus :one
SELECT count(*) FROM `goadmin_role_menu`;

-- name: CreateGoadminRoleMenu :execresult
INSERT INTO `goadmin_role_menu` (
`menu_id`,`role_id`,`updated_at`
) VALUES (
? ,? ,? 
);

-- name: UpdateGoadminRoleMenu :exec
UPDATE `goadmin_role_menu`
SET 
  
  `menu_id` = CASE WHEN sqlc.arg('menu_id') IS NOT NULL THEN sqlc.arg('menu_id') ELSE `menu_id` END,
  `role_id` = CASE WHEN sqlc.arg('role_id') IS NOT NULL THEN sqlc.arg('role_id') ELSE `role_id` END,
  `updated_at` = CASE WHEN sqlc.arg('updated_at') IS NOT NULL THEN sqlc.arg('updated_at') ELSE `updated_at` END
WHERE created_at = ?;

-- name: DeleteGoadminRoleMenu :exec
DELETE FROM `goadmin_role_menu`
WHERE created_at = ?;
