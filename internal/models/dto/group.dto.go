package dto

import "github.com/crossevol/goadmin/internal/database/mysqlDao"

type CreateGroupDTO struct {
	Desc string `db:"desc" json:"desc"`
	Name string `db:"name" json:"name"`
}

type UpdateGroupDTO struct {
	Desc string `db:"desc" json:"desc"`
	Name string `db:"name" json:"name"`
}

func NewCreateGroupParams(dto *CreateGroupDTO) *mysqlDao.CreateGroupParams {
	return &mysqlDao.CreateGroupParams{
		Desc: dto.Desc,
		Name: dto.Name,
	}
}

func NewUpdateGroupParams(dto *UpdateGroupDTO) *mysqlDao.UpdateGroupParams {
	return &mysqlDao.UpdateGroupParams{
		Desc: dto.Desc,
		Name: dto.Name,
	}
}
