
-- name: GetGoadminPermission :one
SELECT * FROM `goadmin_permissions`
WHERE id = ? LIMIT 1;

-- name: GetGoadminPermissions :many
SELECT * FROM `goadmin_permissions`;

-- name: CountGoadminPermissions :one
SELECT count(*) FROM `goadmin_permissions`;

-- name: CreateGoadminPermission :execresult
INSERT INTO `goadmin_permissions` (
`created_at`,`http_method`,`http_path`,`name`,`slug`,`updated_at`
) VALUES (
? ,? ,? ,? ,? ,? 
);

-- name: UpdateGoadminPermission :exec
UPDATE `goadmin_permissions`
SET 
  `created_at` = CASE WHEN sqlc.arg('created_at') IS NOT NULL THEN sqlc.arg('created_at') ELSE `created_at` END,
  `http_method` = CASE WHEN sqlc.arg('http_method') IS NOT NULL THEN sqlc.arg('http_method') ELSE `http_method` END,
  `http_path` = CASE WHEN sqlc.arg('http_path') IS NOT NULL THEN sqlc.arg('http_path') ELSE `http_path` END,
  
  `name` = CASE WHEN sqlc.arg('name') IS NOT NULL THEN sqlc.arg('name') ELSE `name` END,
  `slug` = CASE WHEN sqlc.arg('slug') IS NOT NULL THEN sqlc.arg('slug') ELSE `slug` END,
  `updated_at` = CASE WHEN sqlc.arg('updated_at') IS NOT NULL THEN sqlc.arg('updated_at') ELSE `updated_at` END
WHERE id = ?;

-- name: DeleteGoadminPermission :exec
DELETE FROM `goadmin_permissions`
WHERE id = ?;
