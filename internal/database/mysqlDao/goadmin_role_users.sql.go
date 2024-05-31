// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: goadmin_role_users.sql

package mysqlDao

import (
	"context"
	"database/sql"
)

const CountGoadminRoleUsers = `-- name: CountGoadminRoleUsers :one
SELECT count(*) FROM ` + "`" + `goadmin_role_users` + "`" + `
`

func (q *Queries) CountGoadminRoleUsers(ctx context.Context) (int64, error) {
	row := q.queryRow(ctx, q.countGoadminRoleUsersStmt, CountGoadminRoleUsers)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const CreateGoadminRoleUser = `-- name: CreateGoadminRoleUser :execresult
INSERT INTO ` + "`" + `goadmin_role_users` + "`" + ` (
` + "`" + `user_id` + "`" + `,` + "`" + `created_at` + "`" + `,` + "`" + `updated_at` + "`" + `
) VALUES (
? ,? ,? 
)
`

type CreateGoadminRoleUserParams struct {
	UserID    uint32       `db:"user_id" json:"user_id"`
	CreatedAt sql.NullTime `db:"created_at" json:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at" json:"updated_at"`
}

func (q *Queries) CreateGoadminRoleUser(ctx context.Context, arg CreateGoadminRoleUserParams) (sql.Result, error) {
	return q.exec(ctx, q.createGoadminRoleUserStmt, CreateGoadminRoleUser, arg.UserID, arg.CreatedAt, arg.UpdatedAt)
}

const DeleteGoadminRoleUser = `-- name: DeleteGoadminRoleUser :exec
DELETE FROM ` + "`" + `goadmin_role_users` + "`" + `
WHERE role_id = ?
`

func (q *Queries) DeleteGoadminRoleUser(ctx context.Context, roleID uint32) error {
	_, err := q.exec(ctx, q.deleteGoadminRoleUserStmt, DeleteGoadminRoleUser, roleID)
	return err
}

const GetGoadminRoleUser = `-- name: GetGoadminRoleUser :one
SELECT role_id, user_id, created_at, updated_at FROM ` + "`" + `goadmin_role_users` + "`" + `
WHERE role_id = ? LIMIT 1
`

func (q *Queries) GetGoadminRoleUser(ctx context.Context, roleID uint32) (GoadminRoleUser, error) {
	row := q.queryRow(ctx, q.getGoadminRoleUserStmt, GetGoadminRoleUser, roleID)
	var i GoadminRoleUser
	err := row.Scan(
		&i.RoleID,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const GetGoadminRoleUsers = `-- name: GetGoadminRoleUsers :many
SELECT role_id, user_id, created_at, updated_at FROM ` + "`" + `goadmin_role_users` + "`" + `
`

func (q *Queries) GetGoadminRoleUsers(ctx context.Context) ([]GoadminRoleUser, error) {
	rows, err := q.query(ctx, q.getGoadminRoleUsersStmt, GetGoadminRoleUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GoadminRoleUser
	for rows.Next() {
		var i GoadminRoleUser
		if err := rows.Scan(
			&i.RoleID,
			&i.UserID,
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

const UpdateGoadminRoleUser = `-- name: UpdateGoadminRoleUser :exec
UPDATE ` + "`" + `goadmin_role_users` + "`" + `
SET 
  
  ` + "`" + `user_id` + "`" + ` = CASE WHEN ? IS NOT NULL THEN ? ELSE ` + "`" + `user_id` + "`" + ` END,
  ` + "`" + `created_at` + "`" + ` = CASE WHEN ? IS NOT NULL THEN ? ELSE ` + "`" + `created_at` + "`" + ` END,
  ` + "`" + `updated_at` + "`" + ` = CASE WHEN ? IS NOT NULL THEN ? ELSE ` + "`" + `updated_at` + "`" + ` END
WHERE role_id = ?
`

type UpdateGoadminRoleUserParams struct {
	UserID    uint32       `db:"user_id" json:"user_id"`
	CreatedAt sql.NullTime `db:"created_at" json:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at" json:"updated_at"`
	RoleID    uint32       `db:"role_id" json:"role_id"`
}

func (q *Queries) UpdateGoadminRoleUser(ctx context.Context, arg UpdateGoadminRoleUserParams) error {
	_, err := q.exec(ctx, q.updateGoadminRoleUserStmt, UpdateGoadminRoleUser,
		arg.UserID,
		arg.UserID,
		arg.CreatedAt,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.UpdatedAt,
		arg.RoleID,
	)
	return err
}