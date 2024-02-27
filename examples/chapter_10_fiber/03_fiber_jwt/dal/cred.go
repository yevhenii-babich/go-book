package dal

import (
	"basic-jwt-auth/models"
	"errors"
)

// FindByCredentials Симулює виклик бази даних
func FindByCredentials(email, password string) (*models.User, error) {
	// Тут ви б запитували вашу базу даних для користувача з даною електронною поштою
	if email == "test@mail.com" && password == "test12345" {
		return &models.User{
			ID:       1,
			Email:    "test@mail.com",
			Password: "test12345",
			NickName: "Test User",
		}, nil
	}
	return nil, errors.New("користувача не знайдено")
}
