package vo

import "github.com/crossevol/goadmin/internal/database/mysqlDao"

type TodoDetailVO struct {
	ID     int    `json:"id"`
	Desc   string `json:"desc"`
	ImgUrl string `json:"img_url"`
	TodoID int    `json:"todo_id"`
}

func NewTodoDetailVO(todoDetail *mysqlDao.TodoDetail) *TodoDetailVO {
	return &TodoDetailVO{
		ID:     todoDetail.ID,
		Desc:   todoDetail.Desc,
		ImgUrl: todoDetail.ImgUrl,
		TodoID: todoDetail.TodoID,
	}
}

func NewTodoDetailVOs(todoDetailList []mysqlDao.TodoDetail) []*TodoDetailVO {
	var todoDetailVOs []*TodoDetailVO
	for _, todoDetail := range todoDetailList {
		todoDetailVOs = append(todoDetailVOs, NewTodoDetailVO(&todoDetail))
	}
	return todoDetailVOs
}
