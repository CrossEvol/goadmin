package dto

import (
	"database/sql"
	"github.com/crossevol/goadmin/internal/database/mysqlDao"
	"time"
)

type CreateTodoDTO struct {
	Amount     float64               `json:"amount"`
	CategoryID int32                 `json:"category_id,omitempty"`
	Content    string                `json:"content"`
	Deadline   int64                 `json:"deadline"`
	Priority   mysqlDao.TodoPriority `json:"priority"`
	Score      int                   `json:"score"`
	Status     mysqlDao.TodoStatus   `json:"status"`
	Title      string                `json:"title"`
}

func NewCreateTodoParams(dto *CreateTodoDTO) *mysqlDao.CreateTodoParams {
	params := mysqlDao.CreateTodoParams{
		Amount:    dto.Amount,
		Content:   dto.Content,
		Deadline:  time.UnixMilli(dto.Deadline),
		Priority:  dto.Priority,
		Score:     dto.Score,
		Status:    dto.Status,
		Title:     dto.Title,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if dto.CategoryID == 0 {
		params.CategoryID = sql.NullInt32{Valid: false, Int32: dto.CategoryID}
	} else {
		params.CategoryID = sql.NullInt32{Valid: true, Int32: dto.CategoryID}
	}
	return &params
}

type UpdateTodoDTO struct {
	Amount     float64               `db:"amount" json:"amount"`
	CategoryID int32                 `db:"category_id" json:"category_id"`
	Content    string                `db:"content" json:"content"`
	Deadline   int64                 `db:"deadline" json:"deadline"`
	Priority   mysqlDao.TodoPriority `db:"priority" json:"priority"`
	Score      int                   `db:"score" json:"score"`
	Status     mysqlDao.TodoStatus   `db:"status" json:"status"`
	Title      string                `db:"title" json:"title"`
}

func NewUpdateTodoParams(dto *UpdateTodoDTO) *mysqlDao.UpdateTodoParams {
	params := mysqlDao.UpdateTodoParams{
		Amount:    dto.Amount,
		Content:   dto.Content,
		Deadline:  time.UnixMilli(dto.Deadline),
		Priority:  dto.Priority,
		Score:     dto.Score,
		Status:    dto.Status,
		Title:     dto.Title,
		UpdatedAt: time.Now(),
	}
	if dto.CategoryID == 0 {
		params.CategoryID = sql.NullInt32{Valid: false, Int32: dto.CategoryID}
	} else {
		params.CategoryID = sql.NullInt32{Valid: true, Int32: dto.CategoryID}
	}
	return &params
}
