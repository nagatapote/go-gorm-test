package usecase

import (
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

type FileUseCase interface {
	FileGetAllUseCase() (resp interface{}, statuscode int, err error)
	FileUploadUseCase(File *multipart.FileHeader) (resp interface{}, statuscode int, err error)
	FileDownloadUseCase(Filename string) (resp []byte, statuscode int, err error)
}

type fileUseCaceImpl struct {
	Fr repository.FileRepository
}

func NewFileUseCase(fr repository.FileRepository) FileUseCase {
	return fileUseCaceImpl{fr}
}

func (fu fileUseCaceImpl) FileGetAllUseCase() (resp interface{}, statuscode int, err error) {
	resp, err = fu.Fr.FileGetAll()
	if err != nil {
		return nil, http.StatusInternalServerError, util.ErrorServerError
	}
	return resp, http.StatusOK, nil
}

func (fu fileUseCaceImpl) FileUploadUseCase(File *multipart.FileHeader) (resp interface{}, statuscode int, err error) {
	filedata, err := File.Open()
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
	filemodel := strings.Split(File.Filename, ".")
	filename := File.Filename
	extension := filemodel[1]
	uploadname := fmt.Sprintf("uploaded_%d.%s", time.Now().UnixNano(), extension)
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
		Key:    aws.String(uploadname),
	})
	if err != nil {
		return nil, http.StatusInternalServerError, util.ErrorServerError
	}
	resp, err = fu.Fr.FileCreate(uploadname, filename)
	if err != nil {
		return nil, http.StatusInternalServerError, util.ErrorServerError
	}
	// resp = string("https://" + os.Getenv("BUCKET_NAME") + ".s3." + os.Getenv("REGION") + ".amazonaws.com/" + uploadName)
	return resp, http.StatusOK, nil
}

func (fu fileUseCaceImpl) FileDownloadUseCase(Finename string) (resp []byte, statuscode int, err error) {
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

	resp = buf.Bytes()
	return resp, http.StatusOK, nil
}
