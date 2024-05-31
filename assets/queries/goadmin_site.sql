
-- name: GetGoadminSite :one
SELECT * FROM `goadmin_site`
WHERE id = ? LIMIT 1;

-- name: GetGoadminSites :many
SELECT * FROM `goadmin_site`;

-- name: CountGoadminSites :one
SELECT count(*) FROM `goadmin_site`;

-- name: CreateGoadminSite :execresult
INSERT INTO `goadmin_site` (
`key`,`value`,`description`,`state`,`created_at`,`updated_at`
) VALUES (
? ,? ,? ,? ,? ,? 
);

-- name: UpdateGoadminSite :exec
UPDATE `goadmin_site`
SET 
  
  `key` = CASE WHEN sqlc.arg('key') IS NOT NULL THEN sqlc.arg('key') ELSE `key` END,
  `value` = CASE WHEN sqlc.arg('value') IS NOT NULL THEN sqlc.arg('value') ELSE `value` END,
  `description` = CASE WHEN sqlc.arg('description') IS NOT NULL THEN sqlc.arg('description') ELSE `description` END,
  `state` = CASE WHEN sqlc.arg('state') IS NOT NULL THEN sqlc.arg('state') ELSE `state` END,
  `created_at` = CASE WHEN sqlc.arg('created_at') IS NOT NULL THEN sqlc.arg('created_at') ELSE `created_at` END,
  `updated_at` = CASE WHEN sqlc.arg('updated_at') IS NOT NULL THEN sqlc.arg('updated_at') ELSE `updated_at` END
WHERE id = ?;

-- name: DeleteGoadminSite :exec
DELETE FROM `goadmin_site`
WHERE id = ?;
