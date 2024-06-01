// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: goadmin_role_menu.sql

package mysqlDao

import (
	"context"
	"database/sql"
	"time"
)

const CountGoadminRoleMenus = `-- name: CountGoadminRoleMenus :one
SELECT count(*) FROM ` + "`" + `goadmin_role_menu` + "`" + `
`

func (q *Queries) CountGoadminRoleMenus(ctx context.Context) (int64, error) {
	row := q.queryRow(ctx, q.countGoadminRoleMenusStmt, CountGoadminRoleMenus)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const CreateGoadminRoleMenu = `-- name: CreateGoadminRoleMenu :execresult
INSERT INTO ` + "`" + `goadmin_role_menu` + "`" + ` (
` + "`" + `menu_id` + "`" + `,` + "`" + `role_id` + "`" + `,` + "`" + `updated_at` + "`" + `
) VALUES (
? ,? ,? 
)
`

type CreateGoadminRoleMenuParams struct {
	MenuID    uint32    `db:"menu_id" json:"menu_id"`
	RoleID    uint32    `db:"role_id" json:"role_id"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

func (q *Queries) CreateGoadminRoleMenu(ctx context.Context, arg CreateGoadminRoleMenuParams) (sql.Result, error) {
	return q.exec(ctx, q.createGoadminRoleMenuStmt, CreateGoadminRoleMenu, arg.MenuID, arg.RoleID, arg.UpdatedAt)
}

const DeleteGoadminRoleMenu = `-- name: DeleteGoadminRoleMenu :exec
DELETE FROM ` + "`" + `goadmin_role_menu` + "`" + `
WHERE created_at = ?
`

func (q *Queries) DeleteGoadminRoleMenu(ctx context.Context, createdAt time.Time) error {
	_, err := q.exec(ctx, q.deleteGoadminRoleMenuStmt, DeleteGoadminRoleMenu, createdAt)
	return err
}

const GetGoadminRoleMenu = `-- name: GetGoadminRoleMenu :one
SELECT role_id, menu_id, created_at, updated_at FROM ` + "`" + `goadmin_role_menu` + "`" + `
WHERE created_at = ? LIMIT 1
`

func (q *Queries) GetGoadminRoleMenu(ctx context.Context, createdAt time.Time) (GoadminRoleMenu, error) {
	row := q.queryRow(ctx, q.getGoadminRoleMenuStmt, GetGoadminRoleMenu, createdAt)
	var i GoadminRoleMenu
	err := row.Scan(
		&i.RoleID,
		&i.MenuID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const GetGoadminRoleMenus = `-- name: GetGoadminRoleMenus :many
SELECT role_id, menu_id, created_at, updated_at FROM ` + "`" + `goadmin_role_menu` + "`" + `
`

func (q *Queries) GetGoadminRoleMenus(ctx context.Context) ([]GoadminRoleMenu, error) {
	rows, err := q.query(ctx, q.getGoadminRoleMenusStmt, GetGoadminRoleMenus)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GoadminRoleMenu
	for rows.Next() {
		var i GoadminRoleMenu
		if err := rows.Scan(
			&i.RoleID,
			&i.MenuID,
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

const UpdateGoadminRoleMenu = `-- name: UpdateGoadminRoleMenu :exec
UPDATE ` + "`" + `goadmin_role_menu` + "`" + `
SET 
  
  ` + "`" + `menu_id` + "`" + ` = CASE WHEN ? IS NOT NULL THEN ? ELSE ` + "`" + `menu_id` + "`" + ` END,
  ` + "`" + `role_id` + "`" + ` = CASE WHEN ? IS NOT NULL THEN ? ELSE ` + "`" + `role_id` + "`" + ` END,
  ` + "`" + `updated_at` + "`" + ` = CASE WHEN ? IS NOT NULL THEN ? ELSE ` + "`" + `updated_at` + "`" + ` END
WHERE created_at = ?
`

type UpdateGoadminRoleMenuParams struct {
	MenuID    uint32    `db:"menu_id" json:"menu_id"`
	RoleID    uint32    `db:"role_id" json:"role_id"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

func (q *Queries) UpdateGoadminRoleMenu(ctx context.Context, arg UpdateGoadminRoleMenuParams) error {
	_, err := q.exec(ctx, q.updateGoadminRoleMenuStmt, UpdateGoadminRoleMenu,
		arg.MenuID,
		arg.MenuID,
		arg.RoleID,
		arg.RoleID,
		arg.UpdatedAt,
		arg.UpdatedAt,
		arg.CreatedAt,
	)
	return err
}
