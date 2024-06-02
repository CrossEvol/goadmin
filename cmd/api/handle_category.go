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

func (app *application) GetCategory(w http.ResponseWriter, r *http.Request) {
	IdStr := r.PathValue("id")
	id, err := strconv.ParseInt(IdStr, 10, 32)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}
	category, err := app.q.GetCategory(*app.ctx, int(id))
	if err != nil {
		app.notFound(w, r)
		return
	}
	response.JSON(w, http.StatusOK, vo.NewCategoryVO(&category))
}

func (app *application) GetCategoryList(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	hasIds := query.Has("id")
	if hasIds {
		app.GetCategoriesByIDs(w, r)
		return
	}
	app.logger.Info(utils.ToJson(query))
	categorys, err := app.q.GetCategories(*app.ctx)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	response.JSON(w, http.StatusOK, vo.NewCategoryVOs(categorys))
}

func (app *application) GetCategoriesByIDs(w http.ResponseWriter, r *http.Request) {
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
	categoriesByIDs, err := app.q.GetCategoriesByIDs(*app.ctx, ids)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	response.JSON(w, http.StatusOK, vo.NewCategoryVOs(categoriesByIDs))
}

func (app *application) CreateCategory(w http.ResponseWriter, r *http.Request) {
	createCategoryDTO := dto.CreateCategoryDTO{}
	err := request.DecodeJSON(w, r, &createCategoryDTO)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}
	createCategoryParams := dto.NewCreateCategoryParams(&createCategoryDTO)
	result, err := app.q.CreateCategory(*app.ctx, *createCategoryParams)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	category, err := app.q.GetCategory(*app.ctx, int(lastInsertId))
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	response.JSON(w, http.StatusOK, vo.NewCategoryVO(&category))
}

func (app *application) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	IdStr := r.PathValue("id")
	id, err := strconv.ParseInt(IdStr, 10, 32)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}
	updateCategoryDTO := dto.UpdateCategoryDTO{}
	err = request.DecodeJSON(w, r, &updateCategoryDTO)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}
	updateCategoryParams := dto.NewUpdateCategoryParams(&updateCategoryDTO)
	updateCategoryParams.ID = int(id)
	_, err = app.q.UpdateCategory(*app.ctx, *updateCategoryParams)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	category, err := app.q.GetCategory(*app.ctx, updateCategoryParams.ID)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	response.JSON(w, http.StatusOK, vo.NewCategoryVO(&category))
}

func (app *application) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	IdStr := r.PathValue("id")
	id, err := strconv.ParseInt(IdStr, 10, 32)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}
	category, err := app.q.GetCategory(*app.ctx, int(id))
	if err != nil {
		app.notFound(w, r)
		return
	}
	err = app.q.DeleteCategory(*app.ctx, int(id))
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	response.JSON(w, http.StatusOK, vo.NewCategoryVO(&category))
}
