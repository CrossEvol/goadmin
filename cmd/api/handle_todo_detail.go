package main

import "net/http"

func (app *application) GetTodoDetail(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is a protected handler"))
}

func (app *application) GetTodoDetailList(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is a protected handler"))
}

func (app *application) GetTodoDetailByIDs(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is a protected handler"))
}

func (app *application) CreateTodoDetail(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is a protected handler"))
}

func (app *application) UpdateTodoDetail(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is a protected handler"))
}

func (app *application) DeleteTodoDetail(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is a protected handler"))
}
