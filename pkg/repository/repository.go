package repository


import (
	"web_server/pkg/models"
)

type DatabaseRepo interface {
	AllUsers() bool

	InsertUser(res models.User) (int, error)
}