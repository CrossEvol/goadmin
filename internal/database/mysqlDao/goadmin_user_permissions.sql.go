// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: goadmin_user_permissions.sql

package dao

import (
	"context"
	"database/sql"
)

const CountGoadminUserPermissions = `-- name: CountGoadminUserPermissions :one
SELECT count(*) FROM ` + "`" + `goadmin_user_permissions` + "`" + `
`

func (q *Queries) CountGoadminUserPermissions(ctx context.Context) (int64, error) {
	row := q.queryRow(ctx, q.countGoadminUserPermissionsStmt, CountGoadminUserPermissions)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const CreateGoadminUserPermission = `-- name: CreateGoadminUserPermission :execresult
INSERT INTO ` + "`" + `goadmin_user_permissions` + "`" + ` (
` + "`" + `permission_id` + "`" + `,` + "`" + `created_at` + "`" + `,` + "`" + `updated_at` + "`" + `
) VALUES (
? ,? ,? 
)
`

type CreateGoadminUserPermissionParams struct {
	PermissionID uint32       `db:"permission_id" json:"permission_id"`
	CreatedAt    sql.NullTime `db:"created_at" json:"created_at"`
	UpdatedAt    sql.NullTime `db:"updated_at" json:"updated_at"`
}

func (q *Queries) CreateGoadminUserPermission(ctx context.Context, arg CreateGoadminUserPermissionParams) (sql.Result, error) {
	return q.exec(ctx, q.createGoadminUserPermissionStmt, CreateGoadminUserPermission, arg.PermissionID, arg.CreatedAt, arg.UpdatedAt)
}

const DeleteGoadminUserPermission = `-- name: DeleteGoadminUserPermission :exec
DELETE FROM ` + "`" + `goadmin_user_permissions` + "`" + `
WHERE user_id = ?
`

func (q *Queries) DeleteGoadminUserPermission(ctx context.Context, userID uint32) error {
	_, err := q.exec(ctx, q.deleteGoadminUserPermissionStmt, DeleteGoadminUserPermission, userID)
	return err
}

const GetGoadminUserPermission = `-- name: GetGoadminUserPermission :one
SELECT user_id, permission_id, created_at, updated_at FROM ` + "`" + `goadmin_user_permissions` + "`" + `
WHERE user_id = ? LIMIT 1
`

func (q *Queries) GetGoadminUserPermission(ctx context.Context, userID uint32) (GoadminUserPermission, error) {
	row := q.queryRow(ctx, q.getGoadminUserPermissionStmt, GetGoadminUserPermission, userID)
	var i GoadminUserPermission
	err := row.Scan(
		&i.UserID,
		&i.PermissionID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const GetGoadminUserPermissions = `-- name: GetGoadminUserPermissions :many
SELECT user_id, permission_id, created_at, updated_at FROM ` + "`" + `goadmin_user_permissions` + "`" + `
`

func (q *Queries) GetGoadminUserPermissions(ctx context.Context) ([]GoadminUserPermission, error) {
	rows, err := q.query(ctx, q.getGoadminUserPermissionsStmt, GetGoadminUserPermissions)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GoadminUserPermission
	for rows.Next() {
		var i GoadminUserPermission
		if err := rows.Scan(
			&i.UserID,
			&i.PermissionID,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const UpdateGoadminUserPermission = `-- name: UpdateGoadminUserPermission :exec
UPDATE ` + "`" + `goadmin_user_permissions` + "`" + `
SET 
  
  ` + "`" + `permission_id` + "`" + ` = CASE WHEN ? IS NOT NULL THEN ? ELSE ` + "`" + `permission_id` + "`" + ` END,
  ` + "`" + `created_at` + "`" + ` = CASE WHEN ? IS NOT NULL THEN ? ELSE ` + "`" + `created_at` + "`" + ` END,
  ` + "`" + `updated_at` + "`" + ` = CASE WHEN ? IS NOT NULL THEN ? ELSE ` + "`" + `updated_at` + "`" + ` END
WHERE user_id = ?
`

type UpdateGoadminUserPermissionParams struct {
	PermissionID uint32       `db:"permission_id" json:"permission_id"`
	CreatedAt    sql.NullTime `db:"created_at" json:"created_at"`
	UpdatedAt    sql.NullTime `db:"updated_at" json:"updated_at"`
	UserID       uint32       `db:"user_id" json:"user_id"`
}

func (q *Queries) UpdateGoadminUserPermission(ctx context.Context, arg UpdateGoadminUserPermissionParams) error {
	_, err := q.exec(ctx, q.updateGoadminUserPermissionStmt, UpdateGoadminUserPermission,
		arg.PermissionID,
		arg.PermissionID,
		arg.CreatedAt,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.UpdatedAt,
		arg.UserID,
	)
	return err
}
