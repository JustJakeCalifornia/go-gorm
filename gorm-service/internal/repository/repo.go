package repository

import (
	"github.com/trite8q1/go-gorm/internal/entity"
)

// Implementierung des Repository Pattern (also Funktionen die mit der Datenbank arbeiten)
type UserRepository interface {
	AddUser() error
	GetUser() (entity.User, error)
}

func AddUser() error {
	user := entity.User{Name: "Max"}

	result := GetInstance().DB.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetUser() (entity.User, error) {
	var user entity.User

	result := GetInstance().DB.First(&user)
	if result.Error != nil {
		return entity.User{}, result.Error
	}
	return user, nil
}
