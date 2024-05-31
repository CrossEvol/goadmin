// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package dao

import (
	"context"
	"database/sql"
)

type Querier interface {
	CountGoadminMenus(ctx context.Context) (int64, error)
	CountGoadminOperationLogs(ctx context.Context) (int64, error)
	CountGoadminPermissions(ctx context.Context) (int64, error)
	CountGoadminRoleMenus(ctx context.Context) (int64, error)
	CountGoadminRolePermissions(ctx context.Context) (int64, error)
	CountGoadminRoleUsers(ctx context.Context) (int64, error)
	CountGoadminRoles(ctx context.Context) (int64, error)
	CountGoadminSessions(ctx context.Context) (int64, error)
	CountGoadminSites(ctx context.Context) (int64, error)
	CountGoadminUserPermissions(ctx context.Context) (int64, error)
	CountGoadminUsers(ctx context.Context) (int64, error)
	CreateGoadminMenu(ctx context.Context, arg CreateGoadminMenuParams) (sql.Result, error)
	CreateGoadminOperationLog(ctx context.Context, arg CreateGoadminOperationLogParams) (sql.Result, error)
	CreateGoadminPermission(ctx context.Context, arg CreateGoadminPermissionParams) (sql.Result, error)
	CreateGoadminRole(ctx context.Context, arg CreateGoadminRoleParams) (sql.Result, error)
	CreateGoadminRoleMenu(ctx context.Context, arg CreateGoadminRoleMenuParams) (sql.Result, error)
	CreateGoadminRolePermission(ctx context.Context, arg CreateGoadminRolePermissionParams) (sql.Result, error)
	CreateGoadminRoleUser(ctx context.Context, arg CreateGoadminRoleUserParams) (sql.Result, error)
	CreateGoadminSession(ctx context.Context, arg CreateGoadminSessionParams) (sql.Result, error)
	CreateGoadminSite(ctx context.Context, arg CreateGoadminSiteParams) (sql.Result, error)
	CreateGoadminUser(ctx context.Context, arg CreateGoadminUserParams) (sql.Result, error)
	CreateGoadminUserPermission(ctx context.Context, arg CreateGoadminUserPermissionParams) (sql.Result, error)
	DeleteGoadminMenu(ctx context.Context, id uint32) error
	DeleteGoadminOperationLog(ctx context.Context, id uint32) error
	DeleteGoadminPermission(ctx context.Context, id uint32) error
	DeleteGoadminRole(ctx context.Context, id uint32) error
	DeleteGoadminRoleMenu(ctx context.Context, roleID uint32) error
	DeleteGoadminRolePermission(ctx context.Context, roleID uint32) error
	DeleteGoadminRoleUser(ctx context.Context, roleID uint32) error
	DeleteGoadminSession(ctx context.Context, id uint32) error
	DeleteGoadminSite(ctx context.Context, id uint32) error
	DeleteGoadminUser(ctx context.Context, id uint32) error
	DeleteGoadminUserPermission(ctx context.Context, userID uint32) error
	GetGoadminMenu(ctx context.Context, id uint32) (GoadminMenu, error)
	GetGoadminMenus(ctx context.Context) ([]GoadminMenu, error)
	GetGoadminOperationLog(ctx context.Context, id uint32) (GoadminOperationLog, error)
	GetGoadminOperationLogs(ctx context.Context) ([]GoadminOperationLog, error)
	GetGoadminPermission(ctx context.Context, id uint32) (GoadminPermission, error)
	GetGoadminPermissions(ctx context.Context) ([]GoadminPermission, error)
	GetGoadminRole(ctx context.Context, id uint32) (GoadminRole, error)
	GetGoadminRoleMenu(ctx context.Context, roleID uint32) (GoadminRoleMenu, error)
	GetGoadminRoleMenus(ctx context.Context) ([]GoadminRoleMenu, error)
	GetGoadminRolePermission(ctx context.Context, roleID uint32) (GoadminRolePermission, error)
	GetGoadminRolePermissions(ctx context.Context) ([]GoadminRolePermission, error)
	GetGoadminRoleUser(ctx context.Context, roleID uint32) (GoadminRoleUser, error)
	GetGoadminRoleUsers(ctx context.Context) ([]GoadminRoleUser, error)
	GetGoadminRoles(ctx context.Context) ([]GoadminRole, error)
	GetGoadminSession(ctx context.Context, id uint32) (GoadminSession, error)
	GetGoadminSessions(ctx context.Context) ([]GoadminSession, error)
	GetGoadminSite(ctx context.Context, id uint32) (GoadminSite, error)
	GetGoadminSites(ctx context.Context) ([]GoadminSite, error)
	GetGoadminUser(ctx context.Context, id uint32) (GoadminUser, error)
	GetGoadminUserPermission(ctx context.Context, userID uint32) (GoadminUserPermission, error)
	GetGoadminUserPermissions(ctx context.Context) ([]GoadminUserPermission, error)
	GetGoadminUsers(ctx context.Context) ([]GoadminUser, error)
	UpdateGoadminMenu(ctx context.Context, arg UpdateGoadminMenuParams) error
	UpdateGoadminOperationLog(ctx context.Context, arg UpdateGoadminOperationLogParams) error
	UpdateGoadminPermission(ctx context.Context, arg UpdateGoadminPermissionParams) error
	UpdateGoadminRole(ctx context.Context, arg UpdateGoadminRoleParams) error
	UpdateGoadminRoleMenu(ctx context.Context, arg UpdateGoadminRoleMenuParams) error
	UpdateGoadminRolePermission(ctx context.Context, arg UpdateGoadminRolePermissionParams) error
	UpdateGoadminRoleUser(ctx context.Context, arg UpdateGoadminRoleUserParams) error
	UpdateGoadminSession(ctx context.Context, arg UpdateGoadminSessionParams) error
	UpdateGoadminSite(ctx context.Context, arg UpdateGoadminSiteParams) error
	UpdateGoadminUser(ctx context.Context, arg UpdateGoadminUserParams) error
	UpdateGoadminUserPermission(ctx context.Context, arg UpdateGoadminUserPermissionParams) error
}

var _ Querier = (*Queries)(nil)
