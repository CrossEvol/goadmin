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

func (app *application) GetTodoDetail(w http.ResponseWriter, r *http.Request) {
	todoIdStr := r.PathValue("todo_id")
	todoID, err := strconv.ParseInt(todoIdStr, 10, 32)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}
	todoDetail, err := app.q.GetTodoDetail(*app.ctx, int(todoID))
	if err != nil {
		app.notFound(w, r)
		return
	}
	response.JSON(w, http.StatusOK, vo.NewTodoDetailVO(&todoDetail))
}

func (app *application) GetTodoDetailList(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	hasIds := query.Has("todo_id")
	if hasIds {
		app.GetTodoDetailByIDs(w, r)
		return
	}
	app.logger.Info(utils.ToJson(query))
	todoDetails, err := app.q.GetTodoDetails(*app.ctx)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	response.JSON(w, http.StatusOK, vo.NewTodoDetailVOs(todoDetails))
}

func (app *application) GetTodoDetailByIDs(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	todoIdStr := query.Get("todo_id")
	todoIDsArr := strings.Split(todoIdStr, ",")
	var ids []int
	for _, s := range todoIDsArr {
		id, err := strconv.ParseInt(s, 10, 32)
		if err != nil {
			app.badRequest(w, r, err)
			return
		}
		ids = append(ids, int(id))
	}
	todoDetailsByIDs, err := app.q.GetTodoDetailsByIDs(*app.ctx, ids)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	response.JSON(w, http.StatusOK, vo.NewTodoDetailVOs(todoDetailsByIDs))
}

func (app *application) CreateTodoDetail(w http.ResponseWriter, r *http.Request) {
	createTodoDetailDTO := dto.CreateTodoDetailDTO{}
	err := request.DecodeJSON(w, r, &createTodoDetailDTO)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}
	createTodoDetailParams := dto.NewCreateTodoDetailParams(&createTodoDetailDTO)
	result, err := app.q.CreateTodoDetail(*app.ctx, *createTodoDetailParams)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	todoDetail, err := app.q.GetTodoDetail(*app.ctx, int(lastInsertId))
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	response.JSON(w, http.StatusOK, vo.NewTodoDetailVO(&todoDetail))
}

func (app *application) UpdateTodoDetail(w http.ResponseWriter, r *http.Request) {
	todoIdStr := r.PathValue("todo_id")
	todoID, err := strconv.ParseInt(todoIdStr, 10, 32)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}
	updateTodoDetailDTO := dto.UpdateTodoDetailDTO{}
	err = request.DecodeJSON(w, r, &updateTodoDetailDTO)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}
	updateTodoDetailParams := dto.NewUpdateTodoDetailParams(&updateTodoDetailDTO)
	updateTodoDetailParams.TodoID = int(todoID)
	_, err = app.q.UpdateTodoDetail(*app.ctx, *updateTodoDetailParams)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	todoDetail, err := app.q.GetTodoDetail(*app.ctx, updateTodoDetailParams.TodoID)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	response.JSON(w, http.StatusOK, vo.NewTodoDetailVO(&todoDetail))
}

func (app *application) DeleteTodoDetail(w http.ResponseWriter, r *http.Request) {
	IdStr := r.PathValue("todo_id")
	id, err := strconv.ParseInt(IdStr, 10, 32)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}
	todoDetail, err := app.q.GetTodoDetail(*app.ctx, int(id))
	if err != nil {
		app.notFound(w, r)
		return
	}
	err = app.q.DeleteTodoDetail(*app.ctx, int(id))
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	response.JSON(w, http.StatusOK, vo.NewTodoDetailVO(&todoDetail))
}
