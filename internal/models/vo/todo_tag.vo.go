package vo

import "github.com/crossevol/goadmin/internal/database/mysqlDao"

type TodoTagVO struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	TodoID    int    `json:"todo_id"`
	CreatedAt int64  `json:"created_at"`
}

func NewTodoTagVO(todoTag *mysqlDao.TodoTag) *TodoTagVO {
	return &TodoTagVO{
		ID:        todoTag.ID,
		Name:      todoTag.Name,
		TodoID:    todoTag.TodoID,
		CreatedAt: todoTag.CreatedAt.UnixMilli(),
	}
}

func NewTodoTagVOList(todoTagList []mysqlDao.TodoTag) []*TodoTagVO {
	var todoTagVOList []*TodoTagVO
	for _, todoTag := range todoTagList {
		todoTagVOList = append(todoTagVOList, NewTodoTagVO(&todoTag))
	}
	return todoTagVOList
}
