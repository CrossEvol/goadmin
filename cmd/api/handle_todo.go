package main

import (
	"github.com/crossevol/goadmin/internal/database/mysqlDao"
	"github.com/crossevol/goadmin/internal/request"
	"github.com/crossevol/goadmin/internal/response"
	"github.com/crossevol/goadmin/internal/utils"
	"net/http"
	"strconv"
	"strings"
)

func (app *application) GetTodo(w http.ResponseWriter, r *http.Request) {
	IdStr := r.PathValue("id")
	id, err := strconv.ParseInt(IdStr, 10, 32)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}
	todo, err := app.q.GetTodo(*app.ctx, int(id))
	if err != nil {
		app.notFound(w, r)
		return
	}
	response.JSON(w, http.StatusOK, todo)
}

func (app *application) GetTodoList(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	app.logger.Info(utils.ToJson(query))
	todos, err := app.q.GetTodos(*app.ctx)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	response.JSON(w, http.StatusOK, todos)
}

func (app *application) GetTodoByIDs(w http.ResponseWriter, r *http.Request) {
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
	todosByIDs, err := app.q.GetTodosByIDs(*app.ctx, ids)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	response.JSON(w, http.StatusOK, todosByIDs)
}

func (app *application) CreateTodo(w http.ResponseWriter, r *http.Request) {
	createTodoParams := mysqlDao.CreateTodoParams{}
	err := request.DecodeJSON(w, r, &createTodoParams)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}
	result, err := app.q.CreateTodo(*app.ctx, createTodoParams)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	todo, err := app.q.GetTodo(*app.ctx, int(lastInsertId))
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	response.JSON(w, http.StatusOK, todo)
}

func (app *application) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	updateTodoParams := mysqlDao.UpdateTodoParams{}
	err := request.DecodeJSON(w, r, &updateTodoParams)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}
	result, err := app.q.UpdateTodo(*app.ctx, updateTodoParams)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	todo, err := app.q.GetTodo(*app.ctx, int(lastInsertId))
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	response.JSON(w, http.StatusOK, todo)
}

func (app *application) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	IdStr := r.PathValue("id")
	id, err := strconv.ParseInt(IdStr, 10, 32)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}
	todo, err := app.q.GetTodo(*app.ctx, int(id))
	if err != nil {
		app.notFound(w, r)
		return
	}
	err = app.q.DeleteTodo(*app.ctx, int(id))
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	response.JSON(w, http.StatusOK, todo)
}
