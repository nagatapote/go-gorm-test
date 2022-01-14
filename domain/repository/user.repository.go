package repository

import (
	"context"
	"fmt"
	"go-todo-test/domain/models"
	"go-todo-test/infrastructure/db"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func insert() {
	user := models.User{
		Email:          null.StringFrom("test@example.com"),
		PasswordDigest: null.StringFrom("digested-password"),
	}
	fmt.Printf("before user = %+v\n", user)
	user.Insert(context.Background(), db.DB, boil.Infer())
	fmt.Printf("after user = %+v\n", user)
}

func update() {
	user := models.User{ID: 1}
	user.Email = null.StringFrom("update@example.com")
	user.Update(context.Background(), db.DB, boil.Infer())
}

func delete() {
	user := models.User{ID: 1}
	user.Delete(context.Background(), db.DB)
}
