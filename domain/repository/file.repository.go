package repository

import (
	"go-gorm-test/domain/models"

	"github.com/jinzhu/gorm"
)

type FileRepository interface {
	FileCreate(Uploadname string, Filename string) (*models.File, error)
	// FileDownload()
}

type fileRepositoryImpl struct {
	db *gorm.DB
}

func NewFileRepository(db *gorm.DB) FileRepository {
	return &fileRepositoryImpl{db}
}

func (fr *fileRepositoryImpl) FileCreate(Uploadname string, Filename string) (*models.File, error) {
	insertFile := models.File{}
	insertFile.UploadName = Uploadname
	insertFile.FilesName = Filename
	err := fr.db.Create(&insertFile).Error
	if err != nil {
		return nil, err
	}
	return &insertFile, nil
}
