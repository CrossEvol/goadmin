-- name: GetCategory :one
SELECT *
FROM `category`
WHERE id = ? LIMIT 1;

-- name: GetCategories :many
SELECT *
FROM `category`;

-- name: GetCategoriesByIDs :many
SELECT *
FROM `category`
WHERE id IN (sqlc.slice('ids'));

-- name: CountCategories :one
SELECT count(*)
FROM `category`;

-- name: CreateCategory :execresult
INSERT INTO `category` (`name`, `parent_id`)
VALUES (?, ?);

-- name: UpdateCategory :exec
UPDATE `category`
SET `name`      = CASE WHEN sqlc.arg('name') IS NOT NULL THEN sqlc.arg('name') ELSE `name` END,
    `parent_id` = CASE WHEN sqlc.arg('parent_id') IS NOT NULL THEN sqlc.arg('parent_id') ELSE `parent_id` END
WHERE id = ?;

-- name: DeleteCategory :exec
DELETE
FROM `category`
WHERE id = ?;
