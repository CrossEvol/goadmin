package vo

import "github.com/crossevol/goadmin/internal/database/mysqlDao"

type CategoryVO struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	ParentID int32  `json:"parent_id"`
}

func NewCategoryVO(category *mysqlDao.Category) *CategoryVO {
	return &CategoryVO{
		ID:       category.ID,
		Name:     category.Name,
		ParentID: category.ParentID.Int32,
	}
}

func NewCategoryVOs(categories []mysqlDao.Category) []*CategoryVO {
	var categoryVOs []*CategoryVO
	for _, category := range categories {
		categoryVOs = append(categoryVOs, NewCategoryVO(&category))
	}
	return categoryVOs
}
