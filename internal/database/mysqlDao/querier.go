// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package mysqlDao

import (
	"context"
	"database/sql"
	"time"
)

type Querier interface {
	CountCategories(ctx context.Context) (int64, error)
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
	CountGroups(ctx context.Context) (int64, error)
	CountTodoDetails(ctx context.Context) (int64, error)
	CountTodoDetailsByTodoID(ctx context.Context, todoID int) (int64, error)
	CountTodoTags(ctx context.Context) (int64, error)
	CountTodos(ctx context.Context) (int64, error)
	CountTodosongroups(ctx context.Context) (int64, error)
	CreateCategory(ctx context.Context, arg CreateCategoryParams) (sql.Result, error)
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
	CreateGroup(ctx context.Context, arg CreateGroupParams) (sql.Result, error)
	CreateTodo(ctx context.Context, arg CreateTodoParams) (sql.Result, error)
	CreateTodoDetail(ctx context.Context, arg CreateTodoDetailParams) (sql.Result, error)
	CreateTodoTag(ctx context.Context, arg CreateTodoTagParams) (sql.Result, error)
	CreateTodosongroup(ctx context.Context, arg CreateTodosongroupParams) (sql.Result, error)
	DeleteCategory(ctx context.Context, id int) error
	DeleteGoadminMenu(ctx context.Context, id uint32) error
	DeleteGoadminOperationLog(ctx context.Context, id uint32) error
	DeleteGoadminPermission(ctx context.Context, id uint32) error
	DeleteGoadminRole(ctx context.Context, id uint32) error
	DeleteGoadminRoleMenu(ctx context.Context, createdAt time.Time) error
	DeleteGoadminRolePermission(ctx context.Context, createdAt time.Time) error
	DeleteGoadminRoleUser(ctx context.Context, createdAt time.Time) error
	DeleteGoadminSession(ctx context.Context, id uint32) error
	DeleteGoadminSite(ctx context.Context, id uint32) error
	DeleteGoadminUser(ctx context.Context, id uint32) error
	DeleteGoadminUserPermission(ctx context.Context, createdAt time.Time) error
	DeleteGroup(ctx context.Context, id int) error
	DeleteTodo(ctx context.Context, id int) error
	DeleteTodoDetail(ctx context.Context, todoID int) error
	DeleteTodoTag(ctx context.Context, id int) error
	DeleteTodosongroup(ctx context.Context, arg DeleteTodosongroupParams) error
	GetCategories(ctx context.Context) ([]Category, error)
	GetCategoriesByIDs(ctx context.Context, ids []int) ([]Category, error)
	GetCategory(ctx context.Context, id int) (Category, error)
	GetGoadminMenu(ctx context.Context, id uint32) (GoadminMenu, error)
	GetGoadminMenus(ctx context.Context) ([]GoadminMenu, error)
	GetGoadminOperationLog(ctx context.Context, id uint32) (GoadminOperationLog, error)
	GetGoadminOperationLogs(ctx context.Context) ([]GoadminOperationLog, error)
	GetGoadminPermission(ctx context.Context, id uint32) (GoadminPermission, error)
	GetGoadminPermissions(ctx context.Context) ([]GoadminPermission, error)
	GetGoadminRole(ctx context.Context, id uint32) (GoadminRole, error)
	GetGoadminRoleMenu(ctx context.Context, createdAt time.Time) (GoadminRoleMenu, error)
	GetGoadminRoleMenus(ctx context.Context) ([]GoadminRoleMenu, error)
	GetGoadminRolePermission(ctx context.Context, createdAt time.Time) (GoadminRolePermission, error)
	GetGoadminRolePermissions(ctx context.Context) ([]GoadminRolePermission, error)
	GetGoadminRoleUser(ctx context.Context, createdAt time.Time) (GoadminRoleUser, error)
	GetGoadminRoleUsers(ctx context.Context) ([]GoadminRoleUser, error)
	GetGoadminRoles(ctx context.Context) ([]GoadminRole, error)
	GetGoadminSession(ctx context.Context, id uint32) (GoadminSession, error)
	GetGoadminSessions(ctx context.Context) ([]GoadminSession, error)
	GetGoadminSite(ctx context.Context, id uint32) (GoadminSite, error)
	GetGoadminSites(ctx context.Context) ([]GoadminSite, error)
	GetGoadminUser(ctx context.Context, id uint32) (GoadminUser, error)
	GetGoadminUserByEmail(ctx context.Context, email string) (GoadminUser, error)
	GetGoadminUserPermission(ctx context.Context, createdAt time.Time) (GoadminUserPermission, error)
	GetGoadminUserPermissions(ctx context.Context) ([]GoadminUserPermission, error)
	GetGoadminUsers(ctx context.Context) ([]GoadminUser, error)
	GetGroup(ctx context.Context, id int) (Group, error)
	GetGroups(ctx context.Context) ([]Group, error)
	GetGroupsByIDs(ctx context.Context, ids []int) ([]Group, error)
	GetTodo(ctx context.Context, id int) (Todo, error)
	GetTodoDetail(ctx context.Context, todoID int) (TodoDetail, error)
	GetTodoDetails(ctx context.Context) ([]TodoDetail, error)
	GetTodoDetailsByIDs(ctx context.Context, ids []int) ([]TodoDetail, error)
	GetTodoTag(ctx context.Context, id int) (TodoTag, error)
	GetTodoTags(ctx context.Context) ([]TodoTag, error)
	GetTodoTagsByIDs(ctx context.Context, ids []int) ([]TodoTag, error)
	GetTodos(ctx context.Context) ([]Todo, error)
	GetTodosByIDs(ctx context.Context, ids []int) ([]Todo, error)
	GetTodosongroup(ctx context.Context, arg GetTodosongroupParams) (Todosongroup, error)
	GetTodosongroups(ctx context.Context) ([]Todosongroup, error)
	UpdateCategory(ctx context.Context, arg UpdateCategoryParams) (sql.Result, error)
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
	UpdateGroup(ctx context.Context, arg UpdateGroupParams) (sql.Result, error)
	UpdateTodo(ctx context.Context, arg UpdateTodoParams) (sql.Result, error)
	UpdateTodoDetail(ctx context.Context, arg UpdateTodoDetailParams) (sql.Result, error)
	UpdateTodoTag(ctx context.Context, arg UpdateTodoTagParams) (sql.Result, error)
	UpdateTodosongroup(ctx context.Context, arg UpdateTodosongroupParams) (sql.Result, error)
}

var _ Querier = (*Queries)(nil)
