-- name: GetTodoDetail :one
SELECT *
FROM `todo_detail`
WHERE todo_id = ?
LIMIT 1;

-- name: GetTodoDetails :many
SELECT *
FROM `todo_detail`;

-- name: GetTodoDetailsByIDs :many
SELECT *
FROM `todo_detail`
WHERE todo_id IN (sqlc.slice('ids'));

-- name: CountTodoDetailsByTodoID :one
SELECT count(*)
FROM `todo_detail`
WHERE todo_id = ?;


-- name: CountTodoDetails :one
SELECT count(*)
FROM `todo_detail`;

-- name: CreateTodoDetail :execresult
INSERT INTO `todo_detail` (`desc`, `img_url`, `todo_id`)
VALUES (?, ?, ?);

-- name: UpdateTodoDetail :execresult
UPDATE `todo_detail`
SET `desc`    = CASE WHEN sqlc.arg('desc') IS NOT NULL THEN sqlc.arg('desc') ELSE `desc` END,
    `img_url` = CASE WHEN sqlc.arg('img_url') IS NOT NULL THEN sqlc.arg('img_url') ELSE `img_url` END
WHERE todo_id = ?;

-- name: DeleteTodoDetail :exec
DELETE
FROM `todo_detail`
WHERE todo_id = ?;
