// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: category.sql

package mysqlDao

import (
	"context"
	"database/sql"
	"strings"
)

const CountCategories = `-- name: CountCategories :one
SELECT count(*)
FROM ` + "`" + `category` + "`" + `
`

func (q *Queries) CountCategories(ctx context.Context) (int64, error) {
	row := q.queryRow(ctx, q.countCategoriesStmt, CountCategories)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const CreateCategory = `-- name: CreateCategory :execresult
INSERT INTO ` + "`" + `category` + "`" + ` (` + "`" + `name` + "`" + `, ` + "`" + `parent_id` + "`" + `)
VALUES (?, ?)
`

type CreateCategoryParams struct {
	Name     string        `db:"name" json:"name"`
	ParentID sql.NullInt32 `db:"parent_id" json:"parent_id"`
}

func (q *Queries) CreateCategory(ctx context.Context, arg CreateCategoryParams) (sql.Result, error) {
	return q.exec(ctx, q.createCategoryStmt, CreateCategory, arg.Name, arg.ParentID)
}

const DeleteCategory = `-- name: DeleteCategory :exec
DELETE
FROM ` + "`" + `category` + "`" + `
WHERE id = ?
`

func (q *Queries) DeleteCategory(ctx context.Context, id int) error {
	_, err := q.exec(ctx, q.deleteCategoryStmt, DeleteCategory, id)
	return err
}

const GetCategories = `-- name: GetCategories :many
SELECT id, name, parent_id
FROM ` + "`" + `category` + "`" + `
`

func (q *Queries) GetCategories(ctx context.Context) ([]Category, error) {
	rows, err := q.query(ctx, q.getCategoriesStmt, GetCategories)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Category
	for rows.Next() {
		var i Category
		if err := rows.Scan(&i.ID, &i.Name, &i.ParentID); err != nil {
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

const GetCategoriesByIDs = `-- name: GetCategoriesByIDs :many
SELECT id, name, parent_id
FROM ` + "`" + `category` + "`" + `
WHERE id IN (/*SLICE:ids*/?)
`

func (q *Queries) GetCategoriesByIDs(ctx context.Context, ids []int) ([]Category, error) {
	query := GetCategoriesByIDs
	var queryParams []interface{}
	if len(ids) > 0 {
		for _, v := range ids {
			queryParams = append(queryParams, v)
		}
		query = strings.Replace(query, "/*SLICE:ids*/?", strings.Repeat(",?", len(ids))[1:], 1)
	} else {
		query = strings.Replace(query, "/*SLICE:ids*/?", "NULL", 1)
	}
	rows, err := q.query(ctx, nil, query, queryParams...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Category
	for rows.Next() {
		var i Category
		if err := rows.Scan(&i.ID, &i.Name, &i.ParentID); err != nil {
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

const GetCategory = `-- name: GetCategory :one
SELECT id, name, parent_id
FROM ` + "`" + `category` + "`" + `
WHERE id = ? LIMIT 1
`

func (q *Queries) GetCategory(ctx context.Context, id int) (Category, error) {
	row := q.queryRow(ctx, q.getCategoryStmt, GetCategory, id)
	var i Category
	err := row.Scan(&i.ID, &i.Name, &i.ParentID)
	return i, err
}

const UpdateCategory = `-- name: UpdateCategory :exec
UPDATE ` + "`" + `category` + "`" + `
SET ` + "`" + `name` + "`" + `      = CASE WHEN ? IS NOT NULL THEN ? ELSE ` + "`" + `name` + "`" + ` END,
    ` + "`" + `parent_id` + "`" + ` = CASE WHEN ? IS NOT NULL THEN ? ELSE ` + "`" + `parent_id` + "`" + ` END
WHERE id = ?
`

type UpdateCategoryParams struct {
	Name     string        `db:"name" json:"name"`
	ParentID sql.NullInt32 `db:"parent_id" json:"parent_id"`
	ID       int           `db:"id" json:"id"`
}

func (q *Queries) UpdateCategory(ctx context.Context, arg UpdateCategoryParams) error {
	_, err := q.exec(ctx, q.updateCategoryStmt, UpdateCategory,
		arg.Name,
		arg.Name,
		arg.ParentID,
		arg.ParentID,
		arg.ID,
	)
	return err
}
