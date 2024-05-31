
-- name: GetGoadminMenu :one
SELECT * FROM `goadmin_menu`
WHERE id = ? LIMIT 1;

-- name: GetGoadminMenus :many
SELECT * FROM `goadmin_menu`;

-- name: CountGoadminMenus :one
SELECT count(*) FROM `goadmin_menu`;

-- name: CreateGoadminMenu :execresult
INSERT INTO `goadmin_menu` (
`parent_id`,`type`,`order`,`title`,`icon`,`uri`,`header`,`plugin_name`,`uuid`,`created_at`,`updated_at`
) VALUES (
? ,? ,? ,? ,? ,? ,? ,? ,? ,? ,? 
);

-- name: UpdateGoadminMenu :exec
UPDATE `goadmin_menu`
SET 
  
  `parent_id` = CASE WHEN sqlc.arg('parent_id') IS NOT NULL THEN sqlc.arg('parent_id') ELSE `parent_id` END,
  `type` = CASE WHEN sqlc.arg('type') IS NOT NULL THEN sqlc.arg('type') ELSE `type` END,
  `order` = CASE WHEN sqlc.arg('order') IS NOT NULL THEN sqlc.arg('order') ELSE `order` END,
  `title` = CASE WHEN sqlc.arg('title') IS NOT NULL THEN sqlc.arg('title') ELSE `title` END,
  `icon` = CASE WHEN sqlc.arg('icon') IS NOT NULL THEN sqlc.arg('icon') ELSE `icon` END,
  `uri` = CASE WHEN sqlc.arg('uri') IS NOT NULL THEN sqlc.arg('uri') ELSE `uri` END,
  `header` = CASE WHEN sqlc.arg('header') IS NOT NULL THEN sqlc.arg('header') ELSE `header` END,
  `plugin_name` = CASE WHEN sqlc.arg('plugin_name') IS NOT NULL THEN sqlc.arg('plugin_name') ELSE `plugin_name` END,
  `uuid` = CASE WHEN sqlc.arg('uuid') IS NOT NULL THEN sqlc.arg('uuid') ELSE `uuid` END,
  `created_at` = CASE WHEN sqlc.arg('created_at') IS NOT NULL THEN sqlc.arg('created_at') ELSE `created_at` END,
  `updated_at` = CASE WHEN sqlc.arg('updated_at') IS NOT NULL THEN sqlc.arg('updated_at') ELSE `updated_at` END
WHERE id = ?;

-- name: DeleteGoadminMenu :exec
DELETE FROM `goadmin_menu`
WHERE id = ?;
