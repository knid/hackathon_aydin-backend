package controllers

import (
	"errors"
	"net/http"

	"github.com/knid/timezen/internal/models"
	"github.com/knid/timezen/internal/utils"
)

func (c *Controller) GetUserByID(id uint) (models.User, error) {
	var user models.User

	result := c.DB.Where("id = ?", id).First(&user)
	if result.Error != nil {
		return models.User{}, errors.New("user not found")
	}

	return user, nil

}

func (c *Controller) GetUserByToken(token *models.Token) (models.User, error) { 
	var user models.User

	result := c.DB.Where("id = ?", token.UserID).First(&user)
	if result.Error != nil {
		return models.User{}, errors.New("token not found")
	}

	return user, nil
}

func (c *Controller) GetUserByEmailPassword(email, password string) (models.User, error) { 
	var user models.User

	result := c.DB.Where("email = ? AND password = ?", email, password).First(&user)
	if result.Error != nil {
		return models.User{}, errors.New("user not found")
	}

	return user, nil
}

func (c *Controller) GetUserFromRequest(r *http.Request) (models.User, error) {
	tokenStr, err := utils.ExtractTokenFromHeader(r)
	if err != nil {
		return models.User{}, err
	}
	token, err := c.GetTokenFromToken(tokenStr)
	if err != nil {
		return models.User{}, err
	}

	if err := token.Validete(); err != nil {
		return models.User{}, err
	}

	user, err := c.GetUserByToken(&token)
	if err != nil {
		return models.User{}, err
	}

	return user, nil

}

func (c *Controller) CreateUser(user models.User) (models.User, error) {
	result := c.DB.Create(&user)
	if result.Error != nil {
		return models.User{}, result.Error
	}
	return user, nil
}

func (c *Controller) UpdateUser(u models.User) (models.User, error) {
	user, err := c.GetUserByID(u.ID)
	if err != nil {
		return models.User{}, err
	}

	err = c.DB.Model(&user).Updates(u).Error
	if err != nil {
		return models.User{}, err
	}

	return u, nil
}

func (c *Controller) DeleteUser(user models.User) error {
	user, err := c.GetUserByID(user.ID)
	if err != nil {
		return err
	}

	return c.DB.Delete(&user).Error
}


