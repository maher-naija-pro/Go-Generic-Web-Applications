package repository


import (
	"web_server/internal/models"
)

type DatabaseRepo interface {
	AllUsers() bool

	InsertUser(res models.User) (int, error)
}
