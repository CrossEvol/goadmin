
-- name: GetGoadminRoleMenu :one
SELECT * FROM `goadmin_role_menu`
WHERE role_id = ? LIMIT 1;

-- name: GetGoadminRoleMenus :many
SELECT * FROM `goadmin_role_menu`;

-- name: CountGoadminRoleMenus :one
SELECT count(*) FROM `goadmin_role_menu`;

-- name: CreateGoadminRoleMenu :execresult
INSERT INTO `goadmin_role_menu` (
`menu_id`,`created_at`,`updated_at`
) VALUES (
? ,? ,? 
);

-- name: UpdateGoadminRoleMenu :exec
UPDATE `goadmin_role_menu`
SET 
  
  `menu_id` = CASE WHEN sqlc.arg('menu_id') IS NOT NULL THEN sqlc.arg('menu_id') ELSE `menu_id` END,
  `created_at` = CASE WHEN sqlc.arg('created_at') IS NOT NULL THEN sqlc.arg('created_at') ELSE `created_at` END,
  `updated_at` = CASE WHEN sqlc.arg('updated_at') IS NOT NULL THEN sqlc.arg('updated_at') ELSE `updated_at` END
WHERE role_id = ?;

-- name: DeleteGoadminRoleMenu :exec
DELETE FROM `goadmin_role_menu`
WHERE role_id = ?;
