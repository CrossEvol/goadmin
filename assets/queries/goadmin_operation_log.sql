
-- name: GetGoadminOperationLog :one
SELECT * FROM `goadmin_operation_log`
WHERE id = ? LIMIT 1;

-- name: GetGoadminOperationLogs :many
SELECT * FROM `goadmin_operation_log`;

-- name: CountGoadminOperationLogs :one
SELECT count(*) FROM `goadmin_operation_log`;

-- name: CreateGoadminOperationLog :execresult
INSERT INTO `goadmin_operation_log` (
`created_at`,`input`,`ip`,`method`,`path`,`updated_at`,`user_id`
) VALUES (
? ,? ,? ,? ,? ,? ,? 
);

-- name: UpdateGoadminOperationLog :exec
UPDATE `goadmin_operation_log`
SET 
  `created_at` = CASE WHEN sqlc.arg('created_at') IS NOT NULL THEN sqlc.arg('created_at') ELSE `created_at` END,
  
  `input` = CASE WHEN sqlc.arg('input') IS NOT NULL THEN sqlc.arg('input') ELSE `input` END,
  `ip` = CASE WHEN sqlc.arg('ip') IS NOT NULL THEN sqlc.arg('ip') ELSE `ip` END,
  `method` = CASE WHEN sqlc.arg('method') IS NOT NULL THEN sqlc.arg('method') ELSE `method` END,
  `path` = CASE WHEN sqlc.arg('path') IS NOT NULL THEN sqlc.arg('path') ELSE `path` END,
  `updated_at` = CASE WHEN sqlc.arg('updated_at') IS NOT NULL THEN sqlc.arg('updated_at') ELSE `updated_at` END,
  `user_id` = CASE WHEN sqlc.arg('user_id') IS NOT NULL THEN sqlc.arg('user_id') ELSE `user_id` END
WHERE id = ?;

-- name: DeleteGoadminOperationLog :exec
DELETE FROM `goadmin_operation_log`
WHERE id = ?;
