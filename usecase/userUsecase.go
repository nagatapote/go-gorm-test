package usecase

import (
	"bytes"
	"fmt"
	"go-gorm-test/domain/repository"
	"go-gorm-test/util"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type UserUseCase interface {
	UserLoginUseCase(Email string, Password string) (resp interface{}, statuscode int, err error)
	UserGetUseCase() (resp interface{}, statuscode int, err error)
	UserCreateUseCase(Email string, Password string) (resp interface{}, statuscode int, err error)
	UserUploadUseCase(file *multipart.FileHeader) (resp interface{}, statuscode int, err error)
	UserDownloadUseCase(Filename string) (resp interface{}, statuscode int, err error)
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

func (uu userUseCaceImpl) UserUploadUseCase(file *multipart.FileHeader) (resp interface{}, statuscode int, err error) {
	filedata, err := file.Open()
	if err != nil {
		return nil, http.StatusInternalServerError, util.ErrorServerError
	}
	defer filedata.Close()
	sess, err := session.NewSessionWithOptions(session.Options{
		Config: aws.Config{Region: aws.String(os.Getenv("REGION"))},
		// Profile: "default", NOTE: 書かなくても"default"になる。
	})
	if err != nil {
		return nil, http.StatusInternalServerError, util.ErrorServerError
	}
	fileModel := strings.Split(file.Filename, ".")
	// fileName := fileModel[0]
	extension := fileModel[1]
	filename := fmt.Sprintf("uploaded_%d.%s", time.Now().UnixNano(), extension)
	if err != nil {
		return nil, http.StatusInternalServerError, util.ErrorServerError
	}
	if err != nil {
		return nil, http.StatusInternalServerError, util.ErrorServerError
	}
	u := s3manager.NewUploader(sess)
	_, err = u.Upload(&s3manager.UploadInput{
		Bucket: aws.String(os.Getenv("BUCKET_NAME")),
		Body:   filedata,
		Key:    aws.String(filename),
	})
	if err != nil {
		return nil, http.StatusInternalServerError, util.ErrorServerError
	}
	resp = string("https://" + os.Getenv("BUCKET_NAME") + ".s3." + os.Getenv("REGION") + ".amazonaws.com/" + filename)
	return resp, http.StatusOK, nil
}

func (uu userUseCaceImpl) UserDownloadUseCase(Finename string) (resp interface{}, statuscode int, err error) {
	sess, err := session.NewSessionWithOptions(session.Options{
		Config: aws.Config{Region: aws.String(os.Getenv("REGION"))},
		// Profile: "default", NOTE: 書かなくても"default"になる。
	})
	if err != nil {
		return nil, http.StatusInternalServerError, util.ErrorServerError
	}
	d := s3manager.NewDownloader(sess)
	buf := aws.NewWriteAtBuffer([]byte{})
	_, err = d.Download(buf, &s3.GetObjectInput{
		Bucket: aws.String(os.Getenv("BUCKET_NAME")),
		Key:    aws.String(Finename),
	})
	if err != nil {
		return nil, http.StatusInternalServerError, util.ErrorServerError
	}
	// TODO: ブラウザでダウンロードできるようにする。
	resp = bytes.NewBuffer(buf.Bytes()).String()
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
