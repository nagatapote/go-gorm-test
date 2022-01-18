package repository

import (
	"go-gorm-test/domain/models"

	"github.com/jinzhu/gorm"
)

type UserRepository interface {
	UserGet() (*[]models.User, error)
	UserPost(Email string, PasswordDigest string) (*models.User, error)
	UserUpdate(ID int, Email string, PasswordDigest string) (*models.User, error)
	UserDelete(ID int) (*models.User, error)
}

type userRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepositoryImpl{db}
}

func (ur *userRepositoryImpl) UserGet() (*[]models.User, error) {
	findUsers := []models.User{}
	err := ur.db.Find(&findUsers).Error
	if err != nil {
		return nil, err
	}
	return &findUsers, nil
}

func (ur *userRepositoryImpl) UserPost(Email string, PasswordDigest string) (*models.User, error) {
	insertUser := models.User{}
	insertUser.Email = Email
	insertUser.PasswordDigest = PasswordDigest
	err := ur.db.Create(&insertUser).Error
	if err != nil {
		return nil, err
	}
	return &insertUser, nil
}

func (ur *userRepositoryImpl) UserUpdate(ID int, Email string, PasswordDigest string) (*models.User, error) {
	findUser := models.User{}
	err := ur.db.Where("id = ?", ID).First(&findUser).Error
	if err != nil {
		return nil, err
	}
	updateUser := models.User{}
	updateUser.Email = Email
	updateUser.PasswordDigest = PasswordDigest
	err = ur.db.Model(&findUser).Update(&updateUser).Error
	if err != nil {
		return nil, err
	}
	err = ur.db.Where("id = ?", ID).First(&findUser).Error
	if err != nil {
		return nil, err
	}
	return &findUser, nil
}

func (ur *userRepositoryImpl) UserDelete(ID int) (*models.User, error) {
	findUser := models.User{}
	err := ur.db.Where("id = ?", ID).First(&findUser).Error
	if err != nil {
		return nil, err
	}
	err = ur.db.Delete(&findUser).Error
	if err != nil {
		return nil, err
	}

	return &findUser, nil
}
