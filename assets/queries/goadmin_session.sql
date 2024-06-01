
-- name: GetGoadminSession :one
SELECT * FROM `goadmin_session`
WHERE id = ? LIMIT 1;

-- name: GetGoadminSessions :many
SELECT * FROM `goadmin_session`;

-- name: CountGoadminSessions :one
SELECT count(*) FROM `goadmin_session`;

-- name: CreateGoadminSession :execresult
INSERT INTO `goadmin_session` (
`created_at`,`sid`,`updated_at`,`values`
) VALUES (
? ,? ,? ,? 
);

-- name: UpdateGoadminSession :exec
UPDATE `goadmin_session`
SET 
  `created_at` = CASE WHEN sqlc.arg('created_at') IS NOT NULL THEN sqlc.arg('created_at') ELSE `created_at` END,
  
  `sid` = CASE WHEN sqlc.arg('sid') IS NOT NULL THEN sqlc.arg('sid') ELSE `sid` END,
  `updated_at` = CASE WHEN sqlc.arg('updated_at') IS NOT NULL THEN sqlc.arg('updated_at') ELSE `updated_at` END,
  `values` = CASE WHEN sqlc.arg('values') IS NOT NULL THEN sqlc.arg('values') ELSE `values` END
WHERE id = ?;

-- name: DeleteGoadminSession :exec
DELETE FROM `goadmin_session`
WHERE id = ?;
