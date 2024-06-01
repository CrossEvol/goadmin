
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
`createdAt`,`name`,`todo_id`
) VALUES (
? ,? ,? 
);

-- name: UpdateTodoTag :execresult
UPDATE `todo_tag`
SET 
  `createdAt` = CASE WHEN sqlc.arg('createdAt') IS NOT NULL THEN sqlc.arg('createdAt') ELSE `createdAt` END,
  
  `name` = CASE WHEN sqlc.arg('name') IS NOT NULL THEN sqlc.arg('name') ELSE `name` END,
  `todo_id` = CASE WHEN sqlc.arg('todo_id') IS NOT NULL THEN sqlc.arg('todo_id') ELSE `todo_id` END
WHERE id = ?;

-- name: DeleteTodoTag :exec
DELETE FROM `todo_tag`
WHERE id = ?;
