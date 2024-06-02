package dto

import "github.com/crossevol/goadmin/internal/database/mysqlDao"

type CreateTodoDetailDTO struct {
	Desc   string `json:"desc"`
	ImgUrl string `json:"img_url"`
	TodoID int    `json:"todo_id"`
}

func NewCreateTodoDetailParams(dto *CreateTodoDetailDTO) *mysqlDao.CreateTodoDetailParams {
	return &mysqlDao.CreateTodoDetailParams{
		Desc:   dto.Desc,
		ImgUrl: dto.ImgUrl,
		TodoID: dto.TodoID,
	}
}

type UpdateTodoDetailDTO struct {
	Desc   string `json:"desc"`
	ImgUrl string `json:"img_url"`
}

func NewUpdateTodoDetailParams(dto *UpdateTodoDetailDTO) *mysqlDao.UpdateTodoDetailParams {
	return &mysqlDao.UpdateTodoDetailParams{
		Desc:   dto.Desc,
		ImgUrl: dto.ImgUrl,
	}
}
