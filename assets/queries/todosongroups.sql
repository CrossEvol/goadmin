-- name: GetTodosongroup :one
SELECT *
FROM `todosongroups`
WHERE todo_id = ?
  AND group_id = ? LIMIT 1;

-- name: GetTodosongroups :many
SELECT *
FROM `todosongroups`;

-- name: CountTodosongroups :one
SELECT count(*)
FROM `todosongroups`;

-- name: CreateTodosongroup :execresult
INSERT INTO `todosongroups` (`todo_id`, `group_id`)
VALUES (?, ?);

-- name: UpdateTodosongroup :execresult
UPDATE `todosongroups`
SET `todo_id`  = CASE
                     WHEN sqlc.arg('todo_id') IS NOT NULL THEN sqlc.arg('todo_id')
                     ELSE `todo_id` END,
    `group_id` = CASE
                     WHEN sqlc.arg('group_id') IS NOT NULL THEN sqlc.arg('group_id')
                     ELSE `group_id`
        END
WHERE assigned_at = ?;

-- name: DeleteTodosongroup :exec
DELETE
FROM `todosongroups`
WHERE todo_id = ?
  AND group_id = ?;
