package repository

import (
	"go-gorm-test/domain/models"

	"github.com/jinzhu/gorm"
)

type UserRepository interface {
	UserFindEmail(Email string) (*models.User, error)
	UserFindID(ID string) (*models.User, error)
	UserFindAll() (*[]models.User, error)
	UserCreate(Email string, Password string) (*models.User, error)
	UserUpdate(ID int, Email string, Password string) (*models.User, error)
	UserDelete(ID int) (*models.User, error)
}

type userRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepositoryImpl{db}
}

func (ur *userRepositoryImpl) UserFindEmail(Email string) (*models.User, error) {
	findUser := models.User{}
	err := ur.db.Where("Email = ?", Email).First(&findUser).Error
	if err != nil {
		return nil, err
	}
	return &findUser, nil
}
func (ur *userRepositoryImpl) UserFindID(ID string) (*models.User, error) {
	findUser := models.User{}
	err := ur.db.Where("id = ?", ID).First(&findUser).Error
	if err != nil {
		return nil, err
	}
	return &findUser, nil
}

func (ur *userRepositoryImpl) UserFindAll() (*[]models.User, error) {
	findUsers := []models.User{}
	err := ur.db.Find(&findUsers).Error
	if err != nil {
		return nil, err
	}
	return &findUsers, nil
}

func (ur *userRepositoryImpl) UserCreate(Email string, userEncryptPassword string) (*models.User, error) {
	insertUser := models.User{}
	insertUser.Email = Email
	insertUser.CryptedPassword = userEncryptPassword
	err := ur.db.Create(&insertUser).Error
	if err != nil {
		return nil, err
	}
	return &insertUser, nil
}

func (ur *userRepositoryImpl) UserUpdate(ID int, Email string, userEncryptPassword string) (*models.User, error) {
	findUser := models.User{}
	err := ur.db.Where("id = ?", ID).First(&findUser).Error
	if err != nil {
		return nil, err
	}
	updateUser := models.User{}
	updateUser.Email = Email
	updateUser.CryptedPassword = userEncryptPassword
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
