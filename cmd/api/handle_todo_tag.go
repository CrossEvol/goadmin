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

func (app *application) GetTodoTag(w http.ResponseWriter, r *http.Request) {
	IdStr := r.PathValue("id")
	id, err := strconv.ParseInt(IdStr, 10, 32)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}
	todoTag, err := app.q.GetTodoTag(*app.ctx, int(id))
	if err != nil {
		app.notFound(w, r)
		return
	}
	response.JSON(w, http.StatusOK, vo.NewTodoTagVO(&todoTag))
}

func (app *application) GetTodoTagList(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	hasIds := query.Has("id")
	if hasIds {
		app.GetTodoTagByIDs(w, r)
		return
	}
	app.logger.Info(utils.ToJson(query))
	todoTags, err := app.q.GetTodoTags(*app.ctx)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	response.JSON(w, http.StatusOK, vo.NewTodoTagVOs(todoTags))
}

func (app *application) GetTodoTagByIDs(w http.ResponseWriter, r *http.Request) {
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
	todoTagsByIDs, err := app.q.GetTodoTagsByIDs(*app.ctx, ids)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	response.JSON(w, http.StatusOK, vo.NewTodoTagVOs(todoTagsByIDs))
}

func (app *application) CreateTodoTag(w http.ResponseWriter, r *http.Request) {
	createTodoTagDTO := dto.CreateTodoTagDTO{}
	err := request.DecodeJSON(w, r, &createTodoTagDTO)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}
	createTodoTagParams := dto.NewCreateTodoTagParams(&createTodoTagDTO)
	result, err := app.q.CreateTodoTag(*app.ctx, *createTodoTagParams)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	todoTag, err := app.q.GetTodoTag(*app.ctx, int(lastInsertId))
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	response.JSON(w, http.StatusOK, vo.NewTodoTagVO(&todoTag))
}

func (app *application) UpdateTodoTag(w http.ResponseWriter, r *http.Request) {
	IdStr := r.PathValue("id")
	id, err := strconv.ParseInt(IdStr, 10, 32)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}
	updateTodoTagDTO := dto.UpdateTodoTagDTO{}
	err = request.DecodeJSON(w, r, &updateTodoTagDTO)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}
	updateTodoTagParams := dto.NewUpdateTodoTagParams(&updateTodoTagDTO)
	updateTodoTagParams.ID = int(id)
	_, err = app.q.UpdateTodoTag(*app.ctx, *updateTodoTagParams)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	todoTag, err := app.q.GetTodoTag(*app.ctx, updateTodoTagParams.ID)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	response.JSON(w, http.StatusOK, vo.NewTodoTagVO(&todoTag))
}

func (app *application) DeleteTodoTag(w http.ResponseWriter, r *http.Request) {
	IdStr := r.PathValue("id")
	id, err := strconv.ParseInt(IdStr, 10, 32)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}
	todoTag, err := app.q.GetTodoTag(*app.ctx, int(id))
	if err != nil {
		app.notFound(w, r)
		return
	}
	err = app.q.DeleteTodoTag(*app.ctx, int(id))
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	response.JSON(w, http.StatusOK, vo.NewTodoTagVO(&todoTag))
}
