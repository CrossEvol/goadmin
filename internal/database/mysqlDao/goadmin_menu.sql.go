// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: goadmin_menu.sql

package mysqlDao

import (
	"context"
	"database/sql"
)

const CountGoadminMenus = `-- name: CountGoadminMenus :one
SELECT count(*) FROM ` + "`" + `goadmin_menu` + "`" + `
`

func (q *Queries) CountGoadminMenus(ctx context.Context) (int64, error) {
	row := q.queryRow(ctx, q.countGoadminMenusStmt, CountGoadminMenus)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const CreateGoadminMenu = `-- name: CreateGoadminMenu :execresult
INSERT INTO ` + "`" + `goadmin_menu` + "`" + ` (
` + "`" + `parent_id` + "`" + `,` + "`" + `type` + "`" + `,` + "`" + `order` + "`" + `,` + "`" + `title` + "`" + `,` + "`" + `icon` + "`" + `,` + "`" + `uri` + "`" + `,` + "`" + `header` + "`" + `,` + "`" + `plugin_name` + "`" + `,` + "`" + `uuid` + "`" + `,` + "`" + `created_at` + "`" + `,` + "`" + `updated_at` + "`" + `
) VALUES (
? ,? ,? ,? ,? ,? ,? ,? ,? ,? ,? 
)
`

type CreateGoadminMenuParams struct {
	ParentID   uint32         `db:"parent_id" json:"parent_id"`
	Type       uint32         `db:"type" json:"type"`
	Order      uint32         `db:"order" json:"order"`
	Title      string         `db:"title" json:"title"`
	Icon       string         `db:"icon" json:"icon"`
	Uri        string         `db:"uri" json:"uri"`
	Header     sql.NullString `db:"header" json:"header"`
	PluginName string         `db:"plugin_name" json:"plugin_name"`
	Uuid       sql.NullString `db:"uuid" json:"uuid"`
	CreatedAt  sql.NullTime   `db:"created_at" json:"created_at"`
	UpdatedAt  sql.NullTime   `db:"updated_at" json:"updated_at"`
}

func (q *Queries) CreateGoadminMenu(ctx context.Context, arg CreateGoadminMenuParams) (sql.Result, error) {
	return q.exec(ctx, q.createGoadminMenuStmt, CreateGoadminMenu,
		arg.ParentID,
		arg.Type,
		arg.Order,
		arg.Title,
		arg.Icon,
		arg.Uri,
		arg.Header,
		arg.PluginName,
		arg.Uuid,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
}

const DeleteGoadminMenu = `-- name: DeleteGoadminMenu :exec
DELETE FROM ` + "`" + `goadmin_menu` + "`" + `
WHERE id = ?
`

func (q *Queries) DeleteGoadminMenu(ctx context.Context, id uint32) error {
	_, err := q.exec(ctx, q.deleteGoadminMenuStmt, DeleteGoadminMenu, id)
	return err
}

const GetGoadminMenu = `-- name: GetGoadminMenu :one
SELECT id, parent_id, type, ` + "`" + `order` + "`" + `, title, icon, uri, header, plugin_name, uuid, created_at, updated_at FROM ` + "`" + `goadmin_menu` + "`" + `
WHERE id = ? LIMIT 1
`

func (q *Queries) GetGoadminMenu(ctx context.Context, id uint32) (GoadminMenu, error) {
	row := q.queryRow(ctx, q.getGoadminMenuStmt, GetGoadminMenu, id)
	var i GoadminMenu
	err := row.Scan(
		&i.ID,
		&i.ParentID,
		&i.Type,
		&i.Order,
		&i.Title,
		&i.Icon,
		&i.Uri,
		&i.Header,
		&i.PluginName,
		&i.Uuid,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const GetGoadminMenus = `-- name: GetGoadminMenus :many
SELECT id, parent_id, type, ` + "`" + `order` + "`" + `, title, icon, uri, header, plugin_name, uuid, created_at, updated_at FROM ` + "`" + `goadmin_menu` + "`" + `
`

func (q *Queries) GetGoadminMenus(ctx context.Context) ([]GoadminMenu, error) {
	rows, err := q.query(ctx, q.getGoadminMenusStmt, GetGoadminMenus)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GoadminMenu
	for rows.Next() {
		var i GoadminMenu
		if err := rows.Scan(
			&i.ID,
			&i.ParentID,
			&i.Type,
			&i.Order,
			&i.Title,
			&i.Icon,
			&i.Uri,
			&i.Header,
			&i.PluginName,
			&i.Uuid,
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

const UpdateGoadminMenu = `-- name: UpdateGoadminMenu :exec
UPDATE ` + "`" + `goadmin_menu` + "`" + `
SET 
  
  ` + "`" + `parent_id` + "`" + ` = CASE WHEN ? IS NOT NULL THEN ? ELSE ` + "`" + `parent_id` + "`" + ` END,
  ` + "`" + `type` + "`" + ` = CASE WHEN ? IS NOT NULL THEN ? ELSE ` + "`" + `type` + "`" + ` END,
  ` + "`" + `order` + "`" + ` = CASE WHEN ? IS NOT NULL THEN ? ELSE ` + "`" + `order` + "`" + ` END,
  ` + "`" + `title` + "`" + ` = CASE WHEN ? IS NOT NULL THEN ? ELSE ` + "`" + `title` + "`" + ` END,
  ` + "`" + `icon` + "`" + ` = CASE WHEN ? IS NOT NULL THEN ? ELSE ` + "`" + `icon` + "`" + ` END,
  ` + "`" + `uri` + "`" + ` = CASE WHEN ? IS NOT NULL THEN ? ELSE ` + "`" + `uri` + "`" + ` END,
  ` + "`" + `header` + "`" + ` = CASE WHEN ? IS NOT NULL THEN ? ELSE ` + "`" + `header` + "`" + ` END,
  ` + "`" + `plugin_name` + "`" + ` = CASE WHEN ? IS NOT NULL THEN ? ELSE ` + "`" + `plugin_name` + "`" + ` END,
  ` + "`" + `uuid` + "`" + ` = CASE WHEN ? IS NOT NULL THEN ? ELSE ` + "`" + `uuid` + "`" + ` END,
  ` + "`" + `created_at` + "`" + ` = CASE WHEN ? IS NOT NULL THEN ? ELSE ` + "`" + `created_at` + "`" + ` END,
  ` + "`" + `updated_at` + "`" + ` = CASE WHEN ? IS NOT NULL THEN ? ELSE ` + "`" + `updated_at` + "`" + ` END
WHERE id = ?
`

type UpdateGoadminMenuParams struct {
	ParentID   uint32         `db:"parent_id" json:"parent_id"`
	Type       uint32         `db:"type" json:"type"`
	Order      uint32         `db:"order" json:"order"`
	Title      string         `db:"title" json:"title"`
	Icon       string         `db:"icon" json:"icon"`
	Uri        string         `db:"uri" json:"uri"`
	Header     sql.NullString `db:"header" json:"header"`
	PluginName string         `db:"plugin_name" json:"plugin_name"`
	Uuid       sql.NullString `db:"uuid" json:"uuid"`
	CreatedAt  sql.NullTime   `db:"created_at" json:"created_at"`
	UpdatedAt  sql.NullTime   `db:"updated_at" json:"updated_at"`
	ID         uint32         `db:"id" json:"id"`
}

func (q *Queries) UpdateGoadminMenu(ctx context.Context, arg UpdateGoadminMenuParams) error {
	_, err := q.exec(ctx, q.updateGoadminMenuStmt, UpdateGoadminMenu,
		arg.ParentID,
		arg.ParentID,
		arg.Type,
		arg.Type,
		arg.Order,
		arg.Order,
		arg.Title,
		arg.Title,
		arg.Icon,
		arg.Icon,
		arg.Uri,
		arg.Uri,
		arg.Header,
		arg.Header,
		arg.PluginName,
		arg.PluginName,
		arg.Uuid,
		arg.Uuid,
		arg.CreatedAt,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.UpdatedAt,
		arg.ID,
	)
	return err
}
