package repository

import (
	"example_project/domain"
)

type DBInterface interface {
	GetUser(name string) (*domain.User, error)
	UpdateUserAge(name string, age int64) error
}
