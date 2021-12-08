package usecase

import (
	"example_project/api/repository"
	"example_project/domain"
	"log"
)

type ProcessUserUsecase interface {
	ProcessUser(login string) (*domain.User, error)
}

type UpdateUserAgeUsecase interface {
	UpdateUserAge(login string, newAge int64)
}

type processUserUsecaseImpl struct {
	dbConn repository.DBInterface
}

func (uc *processUserUsecaseImpl) ProcessUser(login string) (*domain.User, error) {
	log.Println("Start process user")

	user, err := uc.dbConn.GetUser(login)
	if err != nil {
		log.Printf("Get user error: %v", err)
		return nil, err
	}

	log.Printf("User info: %+v", user)
	return user, nil
}

func NewProcessUserUsecase(db repository.DBInterface) ProcessUserUsecase {
	return &processUserUsecaseImpl{
		dbConn: db,
	}
}

type updateUserAgeUsecaseImpl struct {
	dbConn repository.DBInterface
}

func (uc *updateUserAgeUsecaseImpl) UpdateUserAge(login string, newAge int64) {
	log.Println("Start update user age")
	if err := uc.dbConn.UpdateUserAge(login, newAge); err != nil {
		log.Printf("Update user age error: %v", err)
	}
}

func NewUpdateUserAgeUsecase(db repository.DBInterface) UpdateUserAgeUsecase {
	return &updateUserAgeUsecaseImpl{
		dbConn: db,
	}
}
