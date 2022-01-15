package usecase

import (
	"go-sqlboiler-test/domain/repository"
	"go-sqlboiler-test/util"
	"net/http"
)

type UserUseCase interface {
	UserGetUseCase() (resp interface{}, statuscode int, err error)
	// UserPostUseCase(ID float64, Email string) (resp interface{}, statuscode int, err error)
	// UserPutUseCase(ID float64) (resp interface{}, statuscode int, err error)
	// UserDeleteUseCase(ID float64) (resp interface{}, statuscode int, err error)
}

type userUseCaceImpl struct {
	Ur repository.UserRepository
}

func NewUserUseCase(ur repository.UserRepository) UserUseCase {
	return userUseCaceImpl{ur}
}

func (uu userUseCaceImpl) UserGetUseCase() (resp interface{}, statuscode int, err error) {
	resp, err = uu.Ur.UserGet()
	if err != nil {
		return nil, http.StatusInternalServerError, util.ErrorDupricationError
	}
	return resp, http.StatusOK, nil
}
