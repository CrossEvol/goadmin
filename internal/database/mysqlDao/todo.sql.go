// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: todo.sql

package mysqlDao

import (
	"context"
	"database/sql"
	"strings"
	"time"
)

const CountTodos = `-- name: CountTodos :one
SELECT count(*)
FROM ` + "`" + `todo` + "`" + `
`

func (q *Queries) CountTodos(ctx context.Context) (int64, error) {
	row := q.queryRow(ctx, q.countTodosStmt, CountTodos)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const CreateTodo = `-- name: CreateTodo :execresult
INSERT INTO ` + "`" + `todo` + "`" + ` (` + "`" + `amount` + "`" + `, ` + "`" + `category_id` + "`" + `, ` + "`" + `content` + "`" + `, ` + "`" + `created_at` + "`" + `, ` + "`" + `deadline` + "`" + `, ` + "`" + `priority` + "`" + `, ` + "`" + `score` + "`" + `, ` + "`" + `status` + "`" + `,
                    ` + "`" + `title` + "`" + `, ` + "`" + `updated_at` + "`" + `)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
`

type CreateTodoParams struct {
	Amount     float64       `db:"amount" json:"amount"`
	CategoryID sql.NullInt32 `db:"category_id" json:"category_id"`
	Content    string        `db:"content" json:"content"`
	CreatedAt  time.Time     `db:"created_at" json:"created_at"`
	Deadline   time.Time     `db:"deadline" json:"deadline"`
	Priority   TodoPriority  `db:"priority" json:"priority"`
	Score      int           `db:"score" json:"score"`
	Status     TodoStatus    `db:"status" json:"status"`
	Title      string        `db:"title" json:"title"`
	UpdatedAt  time.Time     `db:"updated_at" json:"updated_at"`
}

func (q *Queries) CreateTodo(ctx context.Context, arg CreateTodoParams) (sql.Result, error) {
	return q.exec(ctx, q.createTodoStmt, CreateTodo,
		arg.Amount,
		arg.CategoryID,
		arg.Content,
		arg.CreatedAt,
		arg.Deadline,
		arg.Priority,
		arg.Score,
		arg.Status,
		arg.Title,
		arg.UpdatedAt,
	)
}

const DeleteTodo = `-- name: DeleteTodo :exec
DELETE
FROM ` + "`" + `todo` + "`" + `
WHERE id = ?
`

func (q *Queries) DeleteTodo(ctx context.Context, id int) error {
	_, err := q.exec(ctx, q.deleteTodoStmt, DeleteTodo, id)
	return err
}

const GetTodo = `-- name: GetTodo :one
SELECT id, title, score, amount, status, created_at, updated_at, deadline, priority, content, category_id
FROM ` + "`" + `todo` + "`" + `
WHERE id = ?
LIMIT 1
`

func (q *Queries) GetTodo(ctx context.Context, id int) (Todo, error) {
	row := q.queryRow(ctx, q.getTodoStmt, GetTodo, id)
	var i Todo
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Score,
		&i.Amount,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Deadline,
		&i.Priority,
		&i.Content,
		&i.CategoryID,
	)
	return i, err
}

const GetTodos = `-- name: GetTodos :many
SELECT id, title, score, amount, status, created_at, updated_at, deadline, priority, content, category_id
FROM ` + "`" + `todo` + "`" + `
`

func (q *Queries) GetTodos(ctx context.Context) ([]Todo, error) {
	rows, err := q.query(ctx, q.getTodosStmt, GetTodos)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Todo
	for rows.Next() {
		var i Todo
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Score,
			&i.Amount,
			&i.Status,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Deadline,
			&i.Priority,
			&i.Content,
			&i.CategoryID,
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

const GetTodosByIDs = `-- name: GetTodosByIDs :many
SELECT id, title, score, amount, status, created_at, updated_at, deadline, priority, content, category_id
FROM ` + "`" + `todo` + "`" + `
WHERE id IN (/*SLICE:ids*/?)
`

func (q *Queries) GetTodosByIDs(ctx context.Context, ids []int) ([]Todo, error) {
	query := GetTodosByIDs
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
	var items []Todo
	for rows.Next() {
		var i Todo
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Score,
			&i.Amount,
			&i.Status,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Deadline,
			&i.Priority,
			&i.Content,
			&i.CategoryID,
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

const UpdateTodo = `-- name: UpdateTodo :execresult
UPDATE ` + "`" + `todo` + "`" + `
SET ` + "`" + `amount` + "`" + `      = CASE WHEN ? IS NOT NULL THEN ? ELSE ` + "`" + `amount` + "`" + ` END,
    ` + "`" + `category_id` + "`" + ` = CASE WHEN ? IS NOT NULL THEN ? ELSE ` + "`" + `category_id` + "`" + ` END,
    ` + "`" + `content` + "`" + `     = CASE WHEN ? IS NOT NULL THEN ? ELSE ` + "`" + `content` + "`" + ` END,
    ` + "`" + `deadline` + "`" + `    = CASE WHEN ? IS NOT NULL THEN ? ELSE ` + "`" + `deadline` + "`" + ` END,
    ` + "`" + `priority` + "`" + `    = CASE WHEN ? IS NOT NULL THEN ? ELSE ` + "`" + `priority` + "`" + ` END,
    ` + "`" + `score` + "`" + `       = CASE WHEN ? IS NOT NULL THEN ? ELSE ` + "`" + `score` + "`" + ` END,
    ` + "`" + `status` + "`" + `      = CASE WHEN ? IS NOT NULL THEN ? ELSE ` + "`" + `status` + "`" + ` END,
    ` + "`" + `title` + "`" + `       = CASE WHEN ? IS NOT NULL THEN ? ELSE ` + "`" + `title` + "`" + ` END,
    ` + "`" + `updated_at` + "`" + `  = CASE WHEN ? IS NOT NULL THEN ? ELSE ` + "`" + `updated_at` + "`" + ` END
WHERE id = ?
`

type UpdateTodoParams struct {
	Amount     float64       `db:"amount" json:"amount"`
	CategoryID sql.NullInt32 `db:"category_id" json:"category_id"`
	Content    string        `db:"content" json:"content"`
	Deadline   time.Time     `db:"deadline" json:"deadline"`
	Priority   TodoPriority  `db:"priority" json:"priority"`
	Score      int           `db:"score" json:"score"`
	Status     TodoStatus    `db:"status" json:"status"`
	Title      string        `db:"title" json:"title"`
	UpdatedAt  time.Time     `db:"updated_at" json:"updated_at"`
	ID         int           `db:"id" json:"id"`
}

func (q *Queries) UpdateTodo(ctx context.Context, arg UpdateTodoParams) (sql.Result, error) {
	return q.exec(ctx, q.updateTodoStmt, UpdateTodo,
		arg.Amount,
		arg.Amount,
		arg.CategoryID,
		arg.CategoryID,
		arg.Content,
		arg.Content,
		arg.Deadline,
		arg.Deadline,
		arg.Priority,
		arg.Priority,
		arg.Score,
		arg.Score,
		arg.Status,
		arg.Status,
		arg.Title,
		arg.Title,
		arg.UpdatedAt,
		arg.UpdatedAt,
		arg.ID,
	)
}
