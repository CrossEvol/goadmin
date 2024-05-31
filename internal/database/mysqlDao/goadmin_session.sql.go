// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: goadmin_session.sql

package dao

import (
	"context"
	"database/sql"
)

const CountGoadminSessions = `-- name: CountGoadminSessions :one
SELECT count(*) FROM ` + "`" + `goadmin_session` + "`" + `
`

func (q *Queries) CountGoadminSessions(ctx context.Context) (int64, error) {
	row := q.queryRow(ctx, q.countGoadminSessionsStmt, CountGoadminSessions)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const CreateGoadminSession = `-- name: CreateGoadminSession :execresult
INSERT INTO ` + "`" + `goadmin_session` + "`" + ` (
` + "`" + `sid` + "`" + `,` + "`" + `values` + "`" + `,` + "`" + `created_at` + "`" + `,` + "`" + `updated_at` + "`" + `
) VALUES (
? ,? ,? ,? 
)
`

type CreateGoadminSessionParams struct {
	Sid       string       `db:"sid" json:"sid"`
	Values    string       `db:"values" json:"values"`
	CreatedAt sql.NullTime `db:"created_at" json:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at" json:"updated_at"`
}

func (q *Queries) CreateGoadminSession(ctx context.Context, arg CreateGoadminSessionParams) (sql.Result, error) {
	return q.exec(ctx, q.createGoadminSessionStmt, CreateGoadminSession,
		arg.Sid,
		arg.Values,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
}

const DeleteGoadminSession = `-- name: DeleteGoadminSession :exec
DELETE FROM ` + "`" + `goadmin_session` + "`" + `
WHERE id = ?
`

func (q *Queries) DeleteGoadminSession(ctx context.Context, id uint32) error {
	_, err := q.exec(ctx, q.deleteGoadminSessionStmt, DeleteGoadminSession, id)
	return err
}

const GetGoadminSession = `-- name: GetGoadminSession :one
SELECT id, sid, ` + "`" + `values` + "`" + `, created_at, updated_at FROM ` + "`" + `goadmin_session` + "`" + `
WHERE id = ? LIMIT 1
`

func (q *Queries) GetGoadminSession(ctx context.Context, id uint32) (GoadminSession, error) {
	row := q.queryRow(ctx, q.getGoadminSessionStmt, GetGoadminSession, id)
	var i GoadminSession
	err := row.Scan(
		&i.ID,
		&i.Sid,
		&i.Values,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const GetGoadminSessions = `-- name: GetGoadminSessions :many
SELECT id, sid, ` + "`" + `values` + "`" + `, created_at, updated_at FROM ` + "`" + `goadmin_session` + "`" + `
`

func (q *Queries) GetGoadminSessions(ctx context.Context) ([]GoadminSession, error) {
	rows, err := q.query(ctx, q.getGoadminSessionsStmt, GetGoadminSessions)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GoadminSession
	for rows.Next() {
		var i GoadminSession
		if err := rows.Scan(
			&i.ID,
			&i.Sid,
			&i.Values,
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

const UpdateGoadminSession = `-- name: UpdateGoadminSession :exec
UPDATE ` + "`" + `goadmin_session` + "`" + `
SET 
  
  ` + "`" + `sid` + "`" + ` = CASE WHEN ? IS NOT NULL THEN ? ELSE ` + "`" + `sid` + "`" + ` END,
  ` + "`" + `values` + "`" + ` = CASE WHEN ? IS NOT NULL THEN ? ELSE ` + "`" + `values` + "`" + ` END,
  ` + "`" + `created_at` + "`" + ` = CASE WHEN ? IS NOT NULL THEN ? ELSE ` + "`" + `created_at` + "`" + ` END,
  ` + "`" + `updated_at` + "`" + ` = CASE WHEN ? IS NOT NULL THEN ? ELSE ` + "`" + `updated_at` + "`" + ` END
WHERE id = ?
`

type UpdateGoadminSessionParams struct {
	Sid       string       `db:"sid" json:"sid"`
	Values    string       `db:"values" json:"values"`
	CreatedAt sql.NullTime `db:"created_at" json:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at" json:"updated_at"`
	ID        uint32       `db:"id" json:"id"`
}

func (q *Queries) UpdateGoadminSession(ctx context.Context, arg UpdateGoadminSessionParams) error {
	_, err := q.exec(ctx, q.updateGoadminSessionStmt, UpdateGoadminSession,
		arg.Sid,
		arg.Sid,
		arg.Values,
		arg.Values,
		arg.CreatedAt,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.UpdatedAt,
		arg.ID,
	)
	return err
}
