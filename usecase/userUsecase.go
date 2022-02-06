package usecase

import (
	"go-gorm-test/domain/repository"
	"go-gorm-test/util"
	"net/http"
)

type UserUseCase interface {
	UserLoginUseCase(Email string, Password string) (resp interface{}, statuscode int, err error)
	UserGetUseCase() (resp interface{}, statuscode int, err error)
	UserCreateUseCase(Email string, Password string) (resp interface{}, statuscode int, err error)
	UserUpdateUseCase(ID int, Email string, Password string) (resp interface{}, statuscode int, err error)
	UserDeleteUseCase(ID int) (resp interface{}, statuscode int, err error)
}

type userUseCaceImpl struct {
	Ur repository.UserRepository
	Pu util.PasswordUtil
	Au util.AuthUtil
}

func NewUserUseCase(ur repository.UserRepository, pu util.PasswordUtil, au util.AuthUtil) UserUseCase {
	return userUseCaceImpl{ur, pu, au}
}

type userloginresult struct {
	Token string `json:"token"`
}

func (uu userUseCaceImpl) UserLoginUseCase(Email string, Password string) (resp interface{}, statuscode int, err error) {
	findUser, err := uu.Ur.UserFindEmail(Email)
	if err != nil {
		return nil, http.StatusUnauthorized, util.ErrorEmailOrPassIsWrong
	}

	if findUser.Email == "" {
		return nil, http.StatusUnauthorized, util.ErrorEmailOrPassIsWrong
	}

	err = uu.Pu.PasswordVerify(findUser.CryptedPassword, Password)
	if err != nil {
		return nil, http.StatusUnauthorized, util.ErrorEmailOrPassIsWrong
	}
	token, err := uu.Au.GetToken(findUser.Email)
	if err != nil {
		return nil, http.StatusUnauthorized, util.ErrorEmailOrPassIsWrong
	}
	resp = &userloginresult{
		Token: *token,
	}
	return resp, http.StatusOK, nil

}

func (uu userUseCaceImpl) UserGetUseCase() (resp interface{}, statuscode int, err error) {
	resp, err = uu.Ur.UserFindAll()
	if err != nil {
		return nil, http.StatusInternalServerError, util.ErrorServerError
	}
	return resp, http.StatusOK, nil
}

func (uu userUseCaceImpl) UserCreateUseCase(Email string, Password string) (resp interface{}, statuscode int, err error) {
	encryptPassword, err := uu.Pu.PasswordGenerate(Password)
	if err != nil {
		return nil, http.StatusInternalServerError, util.ErrorServerError
	}
	resp, err = uu.Ur.UserCreate(Email, string(encryptPassword))
	if err != nil {
		return nil, http.StatusInternalServerError, util.ErrorServerError
	}
	return resp, http.StatusOK, nil
}

func (uu userUseCaceImpl) UserUpdateUseCase(ID int, Email string, Password string) (resp interface{}, statuscode int, err error) {
	encryptPassword, err := uu.Pu.PasswordGenerate(Password)
	if err != nil {
		return nil, http.StatusInternalServerError, util.ErrorServerError
	}
	resp, err = uu.Ur.UserUpdate(ID, Email, string(encryptPassword))
	if err != nil {
		return nil, http.StatusInternalServerError, util.ErrorServerError
	}
	return resp, http.StatusOK, nil
}

func (uu userUseCaceImpl) UserDeleteUseCase(ID int) (resp interface{}, statuscode int, err error) {
	resp, err = uu.Ur.UserDelete(ID)
	if err != nil {
		return nil, http.StatusInternalServerError, util.ErrorServerError
	}
	return resp, http.StatusOK, nil
}
