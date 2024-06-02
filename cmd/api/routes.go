package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()

	mux.NotFound(app.notFound)
	mux.MethodNotAllowed(app.methodNotAllowed)

	mux.Use(app.logAccess)
	mux.Use(app.recoverPanic)
	mux.Use(app.authenticate)

	mux.Get("/status", app.status)
	mux.Post("/users", app.createUser)
	mux.Post("/authentication-tokens", app.createAuthenticationToken)

	mux.Route("/todo", func(r chi.Router) {
		r.Get("/{id}", app.GetTodo)
		r.Get("/", app.GetTodoList)
		r.Post("/", app.CreateTodo)
		r.Patch("/{id}", app.UpdateTodo)
		r.Delete("/{id}", app.DeleteTodo)
	})

	mux.Route("/todo-tag", func(r chi.Router) {
		r.Get("/{id}", app.GetTodoTag)
		r.Get("/", app.GetTodoTagList)
		r.Post("/", app.CreateTodoTag)
		r.Patch("/{id}", app.UpdateTodoTag)
		r.Delete("/{id}", app.DeleteTodoTag)
	})

	mux.Route("/todo-detail", func(r chi.Router) {
		r.Get("/{todo_id}", app.GetTodoDetail)
		r.Get("/", app.GetTodoDetailList)
		r.Post("/", app.CreateTodoDetail)
		r.Patch("/{todo_id}", app.UpdateTodoDetail)
		r.Delete("/{todo_id}", app.DeleteTodoDetail)
	})

	mux.Route("/category", func(r chi.Router) {
		r.Get("/{id}", app.GetCategory)
		r.Get("/", app.GetCategoryList)
		r.Post("/", app.CreateCategory)
		r.Patch("/{id}", app.UpdateCategory)
		r.Delete("/{id}", app.DeleteCategory)
	})

	mux.Group(func(mux chi.Router) {
		mux.Use(app.requireAuthenticatedUser)

		mux.Get("/auth/me", app.me)

		mux.Get("/protected", app.protected)
	})

	mux.Group(func(mux chi.Router) {
		mux.Use(app.requireBasicAuthentication)

		mux.Get("/basic-auth-protected", app.protected)
	})

	mux.Post("/auth/login", app.login)

	return mux
}
