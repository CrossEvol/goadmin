package main

import (
	"github.com/crossevol/goadmin/internal/models/dto"
	"github.com/crossevol/goadmin/internal/models/vo"
	"github.com/crossevol/goadmin/internal/request"
	"github.com/crossevol/goadmin/internal/response"
	"github.com/crossevol/goadmin/internal/utils"
	"net/http"
	"strconv"
	"strings"
)

func (app *application) GetGroup(w http.ResponseWriter, r *http.Request) {
	IdStr := r.PathValue("id")
	id, err := strconv.ParseInt(IdStr, 10, 32)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}
	group, err := app.q.GetGroup(*app.ctx, int(id))
	if err != nil {
		app.notFound(w, r)
		return
	}
	response.JSON(w, http.StatusOK, vo.NewGroupVO(&group))
}

func (app *application) GetGroupList(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	hasIds := query.Has("id")
	if hasIds {
		app.GetGroupByIDs(w, r)
		return
	}
	app.logger.Info(utils.ToJson(query))
	groups, err := app.q.GetGroups(*app.ctx)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	response.JSON(w, http.StatusOK, vo.NewGroupVOs(groups))
}

func (app *application) GetGroupByIDs(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	idsStr := query.Get("id")
	idsArr := strings.Split(idsStr, ",")
	var ids []int
	for _, s := range idsArr {
		id, err := strconv.ParseInt(s, 10, 32)
		if err != nil {
			app.badRequest(w, r, err)
			return
		}
		ids = append(ids, int(id))
	}
	groupsByIDs, err := app.q.GetGroupsByIDs(*app.ctx, ids)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	response.JSON(w, http.StatusOK, vo.NewGroupVOs(groupsByIDs))
}

func (app *application) CreateGroup(w http.ResponseWriter, r *http.Request) {
	createGroupDTO := dto.CreateGroupDTO{}
	err := request.DecodeJSON(w, r, &createGroupDTO)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}
	createGroupParams := dto.NewCreateGroupParams(&createGroupDTO)
	result, err := app.q.CreateGroup(*app.ctx, *createGroupParams)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	group, err := app.q.GetGroup(*app.ctx, int(lastInsertId))
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	response.JSON(w, http.StatusOK, vo.NewGroupVO(&group))
}

func (app *application) UpdateGroup(w http.ResponseWriter, r *http.Request) {
	IdStr := r.PathValue("id")
	id, err := strconv.ParseInt(IdStr, 10, 32)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}
	updateGroupDTO := dto.UpdateGroupDTO{}
	err = request.DecodeJSON(w, r, &updateGroupDTO)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}
	updateGroupParams := dto.NewUpdateGroupParams(&updateGroupDTO)
	updateGroupParams.ID = int(id)
	_, err = app.q.UpdateGroup(*app.ctx, *updateGroupParams)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	group, err := app.q.GetGroup(*app.ctx, updateGroupParams.ID)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	response.JSON(w, http.StatusOK, vo.NewGroupVO(&group))
}

func (app *application) DeleteGroup(w http.ResponseWriter, r *http.Request) {
	IdStr := r.PathValue("id")
	id, err := strconv.ParseInt(IdStr, 10, 32)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}
	group, err := app.q.GetGroup(*app.ctx, int(id))
	if err != nil {
		app.notFound(w, r)
		return
	}
	err = app.q.DeleteGroup(*app.ctx, int(id))
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	response.JSON(w, http.StatusOK, vo.NewGroupVO(&group))
}
