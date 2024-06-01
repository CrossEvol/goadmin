package vo

import (
	"github.com/crossevol/goadmin/internal/database/mysqlDao"
)

type TodoVO struct {
	ID         int                   `json:"id"`
	Title      string                `json:"title"`
	Score      int                   `json:"score"`
	Amount     float64               `json:"amount"`
	Status     mysqlDao.TodoStatus   `json:"status"`
	CreatedAt  int64                 `json:"created_at"`
	UpdatedAt  int64                 `json:"updated_at"`
	Deadline   int64                 `json:"deadline"`
	Priority   mysqlDao.TodoPriority `json:"priority"`
	Content    string                `json:"content"`
	CategoryID int32                 `json:"category_id,omitempty"`
}

func NewTodoVO(todo *mysqlDao.Todo) *TodoVO {
	return &TodoVO{
		ID:         todo.ID,
		Title:      todo.Title,
		Score:      todo.Score,
		Amount:     todo.Amount,
		Status:     todo.Status,
		CreatedAt:  todo.CreatedAt.UnixMilli(),
		UpdatedAt:  todo.UpdatedAt.UnixMilli(),
		Deadline:   todo.Deadline.UnixMilli(),
		Priority:   todo.Priority,
		Content:    todo.Content,
		CategoryID: todo.CategoryID.Int32,
	}
}

func NewTodoVOs(todoList []mysqlDao.Todo) []*TodoVO {
	var todoVoList []*TodoVO
	for _, todo := range todoList {
		todoVoList = append(todoVoList, NewTodoVO(&todo))
	}
	return todoVoList
}
