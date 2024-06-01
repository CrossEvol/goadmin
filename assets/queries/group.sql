-- name: GetGroup :one
SELECT *
FROM `group`
WHERE id = ? LIMIT 1;

-- name: GetGroups :many
SELECT *
FROM `group`;

-- name: GetGroupsByIDs :many
SELECT *
FROM `group`
WHERE id IN (sqlc.slice('ids'));

-- name: CountGroups :one
SELECT count(*)
FROM `group`;

-- name: CreateGroup :execresult
INSERT INTO `group` (`desc`, `name`)
VALUES (?, ?);

-- name: UpdateGroup :execresult
UPDATE `group`
SET `desc` = CASE WHEN sqlc.arg('desc') IS NOT NULL THEN sqlc.arg('desc') ELSE `desc` END,
    `name` = CASE WHEN sqlc.arg('name') IS NOT NULL THEN sqlc.arg('name') ELSE `name` END
WHERE id = ?;

-- name: DeleteGroup :exec
DELETE
FROM `group`
WHERE id = ?;
