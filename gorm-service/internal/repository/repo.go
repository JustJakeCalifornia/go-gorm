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

	conn := getDatabase()
	dbInstance := Database{db: conn}
	result := dbInstance.db.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetUser() (entity.User, error) {
	var user entity.User

	conn := getDatabase()
	dbInstance := Database{db: conn}
	result := dbInstance.db.First(&user)
	if result.Error != nil {
		return entity.User{}, result.Error
	}
	return user, nil
}
