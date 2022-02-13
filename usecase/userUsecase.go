package usecase

import (
	"go-gorm-test/domain/repository"
	"go-gorm-test/util"
	"log"
	"net/http"
	"time"

	"github.com/pquerna/otp/totp"
)

type UserUseCase interface {
	UserLoginUseCase(Email string, Password string) (resp interface{}, statuscode int, err error)
	UserTotpUseCase(Email string, Totp string) (resp interface{}, statuscode int, err error)
	UserGetAllUseCase() (resp interface{}, statuscode int, err error)
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

	key, err := totp.Generate(totp.GenerateOpts{
			Issuer: "Example.com",
			AccountName: Email,
	})
	if err != nil {
		return nil, http.StatusUnauthorized, util.ErrorEmailOrPassIsWrong
	}

	secret := key.Secret()

	passcode, err := totp.GenerateCode(secret, time.Now())
	if err != nil {
		return nil, http.StatusUnauthorized, util.ErrorEmailOrPassIsWrong
	}
	_, err = uu.Ur.UserTotpUpdate(Email, secret)
	if err != nil {
		return nil, http.StatusUnauthorized, util.ErrorEmailOrPassIsWrong
	}
	log.Println(passcode)
	return nil, http.StatusOK, nil
}

func (uu userUseCaceImpl) UserTotpUseCase(Email string, Totp string) (resp interface{}, statuscode int, err error) {
	findUser, err := uu.Ur.UserFindEmail(Email)
	if err != nil {
		return nil, http.StatusUnauthorized, util.ErrorEmailOrPassIsWrong
	}

	if findUser.Email == "" {
		return nil, http.StatusUnauthorized, util.ErrorEmailOrPassIsWrong
	}
	valid := totp.Validate(Totp, findUser.TotpSecret)
	if !valid {
		return nil, http.StatusUnauthorized, util.ErrorTotpIsWrong
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

func (uu userUseCaceImpl) UserGetAllUseCase() (resp interface{}, statuscode int, err error) {
	resp, err = uu.Ur.UserGetAll()
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
