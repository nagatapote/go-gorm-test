package repository

import (
	"context"
	"database/sql"
	"go-sqlboiler-test/domain/models"
	"go-sqlboiler-test/infrastructure/db"
)

type UserRepository interface {
	UserGet() (*models.UserSlice, error)
	// UserPost(ID float64, Email string) (*models.User, error)
	// UserPut(ID float64) (*models.User, error)
	// UserDelete(ID float64) (*models.User, error)
}

type userRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepositoryImpl{db}
}

func (ur *userRepositoryImpl) UserGet() (*models.UserSlice, error) {
	user, _ := models.Users().All(context.Background(), db.DB)

	return &user, nil
}

// func (ur *userRepositoryImpl) UserPost(ID float64, Email string) (*models.User, error) {
// 	user := models.User{
// 		Email:          null.StringFrom("test@example.com"),
// 		PasswordDigest: null.StringFrom("digested-password"),
// 	}
// 	user.Insert(context.Background(), db.DB, boil.Infer())

// 	return &user, nil
// }

// func (ur *userRepositoryImpl) UserPut(ID float64) (*models.User, error) {
// 	user := models.User{ID: 1}
// 	user.Email = null.StringFrom("update@example.com")
// 	user.Update(context.Background(), db.DB, boil.Infer())

// 	return &user, nil
// }

// func (ur *userRepositoryImpl) UserDelete(ID float64) (*models.User, error) {
// 	user := models.User{ID: 1}
// 	user.Delete(context.Background(), db.DB)

// 	return &user, nil
// }
