package dto

import (
	"database/sql"
	"github.com/crossevol/goadmin/internal/database/mysqlDao"
)

type CreateCategoryDTO struct {
	Name     string `json:"name"`
	ParentID int32  `json:"parent_id"`
}

type UpdateCategoryDTO struct {
	Name     string `json:"name"`
	ParentID int32  `json:"parent_id"`
}

func NewCreateCategoryParams(dto *CreateCategoryDTO) *mysqlDao.CreateCategoryParams {
	params := &mysqlDao.CreateCategoryParams{
		Name: dto.Name,
	}
	if dto.ParentID == 0 {
		params.ParentID = sql.NullInt32{Int32: dto.ParentID, Valid: false}
	} else {
		params.ParentID = sql.NullInt32{Int32: dto.ParentID, Valid: true}
	}
	return params
}

func NewUpdateCategoryParams(dto *UpdateCategoryDTO) *mysqlDao.UpdateCategoryParams {
	params := &mysqlDao.UpdateCategoryParams{
		Name: dto.Name,
	}
	if dto.ParentID == 0 {
		params.ParentID = sql.NullInt32{Int32: dto.ParentID, Valid: false}
	} else {
		params.ParentID = sql.NullInt32{Int32: dto.ParentID, Valid: true}
	}
	return params
}
