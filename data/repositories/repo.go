package repositories

import (
	"gorm.io/gorm"
)

type UserRepository interface {
}

type userRepository struct {
	db *gorm.DB
}

// implement repositories

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

