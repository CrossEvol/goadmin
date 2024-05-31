
-- name: GetGoadminRole :one
SELECT * FROM `goadmin_roles`
WHERE id = ? LIMIT 1;

-- name: GetGoadminRoles :many
SELECT * FROM `goadmin_roles`;

-- name: CountGoadminRoles :one
SELECT count(*) FROM `goadmin_roles`;

-- name: CreateGoadminRole :execresult
INSERT INTO `goadmin_roles` (
`name`,`slug`,`created_at`,`updated_at`
) VALUES (
? ,? ,? ,? 
);

-- name: UpdateGoadminRole :exec
UPDATE `goadmin_roles`
SET 
  
  `name` = CASE WHEN sqlc.arg('name') IS NOT NULL THEN sqlc.arg('name') ELSE `name` END,
  `slug` = CASE WHEN sqlc.arg('slug') IS NOT NULL THEN sqlc.arg('slug') ELSE `slug` END,
  `created_at` = CASE WHEN sqlc.arg('created_at') IS NOT NULL THEN sqlc.arg('created_at') ELSE `created_at` END,
  `updated_at` = CASE WHEN sqlc.arg('updated_at') IS NOT NULL THEN sqlc.arg('updated_at') ELSE `updated_at` END
WHERE id = ?;

-- name: DeleteGoadminRole :exec
DELETE FROM `goadmin_roles`
WHERE id = ?;
