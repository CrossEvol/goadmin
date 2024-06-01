
-- name: GetTodoTag :one
SELECT * FROM `todo_tag`
WHERE id = ? LIMIT 1;

-- name: GetTodoTags :many
SELECT * FROM `todo_tag`;

-- name: GetTodoTagsByIDs :many
SELECT *
FROM `todo_tag`
WHERE id IN (sqlc.slice('ids'));

-- name: CountTodoTags :one
SELECT count(*) FROM `todo_tag`;

-- name: CreateTodoTag :execresult
INSERT INTO `todo_tag` (
`created_at`,`name`,`todo_id`
) VALUES (
? ,? ,? 
);

-- name: UpdateTodoTag :execresult
UPDATE `todo_tag`
SET 
  `created_at` = CASE WHEN sqlc.arg('created_at') IS NOT NULL THEN sqlc.arg('created_at') ELSE `created_at` END,
  
  `name` = CASE WHEN sqlc.arg('name') IS NOT NULL THEN sqlc.arg('name') ELSE `name` END,
  `todo_id` = CASE WHEN sqlc.arg('todo_id') IS NOT NULL THEN sqlc.arg('todo_id') ELSE `todo_id` END
WHERE id = ?;

-- name: DeleteTodoTag :exec
DELETE FROM `todo_tag`
WHERE id = ?;
