package main

import (
	"context"
	"github.com/crossevol/goadmin/internal/database/mysqlDao"
	"net/http"
)

type contextKey string

const (
	authenticatedUserContextKey = contextKey("authenticatedUser")
)

func contextSetAuthenticatedUser(r *http.Request, user *mysqlDao.GoadminUser) *http.Request {
	ctx := context.WithValue(r.Context(), authenticatedUserContextKey, user)
	return r.WithContext(ctx)
}

func contextGetAuthenticatedUser(r *http.Request) *mysqlDao.GoadminUser {
	user, ok := r.Context().Value(authenticatedUserContextKey).(*mysqlDao.GoadminUser)
	if !ok {
		return nil
	}

	return user
}
