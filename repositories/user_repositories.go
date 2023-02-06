package repositories

import (
	"net/http"

	"git.garena.com/sea-labs-id/batch-02/andreas-timothy/entry-task-mo-be/helpers"
	"git.garena.com/sea-labs-id/batch-02/andreas-timothy/entry-task-mo-be/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindUserByEmail(email string) (*models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

type URConfig struct {
	DB *gorm.DB
}

func NewUserRepository(c *URConfig) UserRepository {
	return &userRepository{
		db: c.DB,
	}
}

func (u *userRepository) FindUserByEmail(email string) (*models.User, error) {
	var user *models.User
	if err := u.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, helpers.FailedResponse(http.StatusBadRequest, "Failed to Fetch User")
	}
	return user, nil
}
