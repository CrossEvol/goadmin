package vo

import "github.com/crossevol/goadmin/internal/database/mysqlDao"

type GroupVO struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Desc string `json:"desc"`
}

func NewGroupVO(group *mysqlDao.Group) *GroupVO {
	return &GroupVO{
		ID:   group.ID,
		Name: group.Name,
		Desc: group.Desc,
	}
}

func NewGroupVOs(groups []mysqlDao.Group) []*GroupVO {
	var groupVOs []*GroupVO
	for _, group := range groups {
		groupVOs = append(groupVOs, NewGroupVO(&group))
	}
	return groupVOs
}
