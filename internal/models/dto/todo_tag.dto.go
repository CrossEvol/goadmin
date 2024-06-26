package dto

import (
	"github.com/crossevol/goadmin/internal/database/mysqlDao"
	"time"
)

type CreateTodoTagDTO struct {
	CreatedAt int64  `json:"created_at"`
	Name      string `json:"name"`
	TodoID    int    `json:"todo_id"`
}

func NewCreateTodoTagParams(dto *CreateTodoTagDTO) *mysqlDao.CreateTodoTagParams {
	params := mysqlDao.CreateTodoTagParams{Name: dto.Name, TodoID: dto.TodoID, CreatedAt: time.Now()}
	return &params
}

type UpdateTodoTagDTO struct {
	Name   string `json:"name"`
	TodoID int    `json:"todo_id"`
	ID     int    `json:"id"`
}

func NewUpdateTodoTagParams(dto *UpdateTodoTagDTO) *mysqlDao.UpdateTodoTagParams {
	return &mysqlDao.UpdateTodoTagParams{
		Name:   dto.Name,
		TodoID: dto.TodoID,
	}
}
