package repository

import (
	"errors"

	"github.com/suhas-koheda/mvi-rs/internal/models"
	initializers "github.com/suhas-koheda/mvi-rs/pkg/initialisers"
	"gorm.io/gorm"
)

func FetchUserByEmail(email string) (*models.User, error) {
	var user models.User
	result := initializers.DB.Where("email = ?", email).First(&user)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	return &user, nil
}

func CreateUser(user *models.User) error {
	result := initializers.DB.Create(user)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("no user was created")
	}

	return nil
}

func CheckUserExists(email string) (bool, error) {
	var count int64
	result := initializers.DB.Model(&models.User{}).Where("email = ?", email).Count(&count)

	if result.Error != nil {
		return false, result.Error
	}

	return count > 0, nil
}

func GetUserRole(email string) (models.Role, error) {
	var user models.User
	result := initializers.DB.Where("email = ?", email).Select("role").First(&user)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return "", errors.New("user not found")
		}
		return "", result.Error
	}

	return user.Role, nil
}

func UpdateUser(user *models.User) error {
	result := initializers.DB.Save(user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func DeleteUser(email string) error {
	result := initializers.DB.Where("email = ?", email).Delete(&models.User{})
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("user not found")
	}

	return nil
}
