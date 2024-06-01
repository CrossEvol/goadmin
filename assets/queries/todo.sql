
-- name: GetTodo :one
SELECT * FROM `todo`
WHERE id = ? LIMIT 1;

-- name: GetTodos :many
SELECT * FROM `todo`;

-- name: GetTodosByIDs :many
SELECT *
FROM `todo`
WHERE id IN (sqlc.slice('ids'));

-- name: CountTodos :one
SELECT count(*) FROM `todo`;

-- name: CreateTodo :execresult
INSERT INTO `todo` (
`amount`,`category_id`,`content`,`created_at`,`deadline`,`priority`,`score`,`status`,`title`,`updated_at`
) VALUES (
? ,? ,? ,? ,? ,? ,? ,? ,? ,? 
);

-- name: UpdateTodo :execresult
UPDATE `todo`
SET 
  `amount` = CASE WHEN sqlc.arg('amount') IS NOT NULL THEN sqlc.arg('amount') ELSE `amount` END,
  `category_id` = CASE WHEN sqlc.arg('category_id') IS NOT NULL THEN sqlc.arg('category_id') ELSE `category_id` END,
  `content` = CASE WHEN sqlc.arg('content') IS NOT NULL THEN sqlc.arg('content') ELSE `content` END,
  `created_at` = CASE WHEN sqlc.arg('created_at') IS NOT NULL THEN sqlc.arg('created_at') ELSE `created_at` END,
  `deadline` = CASE WHEN sqlc.arg('deadline') IS NOT NULL THEN sqlc.arg('deadline') ELSE `deadline` END,
  
  `priority` = CASE WHEN sqlc.arg('priority') IS NOT NULL THEN sqlc.arg('priority') ELSE `priority` END,
  `score` = CASE WHEN sqlc.arg('score') IS NOT NULL THEN sqlc.arg('score') ELSE `score` END,
  `status` = CASE WHEN sqlc.arg('status') IS NOT NULL THEN sqlc.arg('status') ELSE `status` END,
  `title` = CASE WHEN sqlc.arg('title') IS NOT NULL THEN sqlc.arg('title') ELSE `title` END,
  `updated_at` = CASE WHEN sqlc.arg('updated_at') IS NOT NULL THEN sqlc.arg('updated_at') ELSE `updated_at` END
WHERE id = ?;

-- name: DeleteTodo :exec
DELETE FROM `todo`
WHERE id = ?;
