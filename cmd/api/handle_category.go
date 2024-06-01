package main

import "net/http"

func (app *application) GetCategory(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is a protected handler"))
}

func (app *application) GetCategoryList(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is a protected handler"))
}

func (app *application) GetCategoryByIDs(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is a protected handler"))
}

func (app *application) CreateCategory(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is a protected handler"))
}

func (app *application) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is a protected handler"))
}

func (app *application) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is a protected handler"))
}
