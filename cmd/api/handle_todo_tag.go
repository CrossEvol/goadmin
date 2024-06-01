package main

import "net/http"

func (app *application) GetTodoTag(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is a protected handler"))
}

func (app *application) GetTodoTagList(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is a protected handler"))
}

func (app *application) GetTodoTagByIDs(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is a protected handler"))
}

func (app *application) CreateTodoTag(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is a protected handler"))
}

func (app *application) UpdateTodoTag(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is a protected handler"))
}

func (app *application) DeleteTodoTag(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is a protected handler"))
}
