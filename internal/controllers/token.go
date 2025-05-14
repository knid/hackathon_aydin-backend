package controllers

import (
	"fmt"

	"github.com/knid/timezen/internal/models"
	"github.com/knid/timezen/internal/utils"
)

func (c *Controller) CreateToken(user models.User) (models.Token, error) {

	token := models.Token{
		Token: utils.GenerateRandomString(64),
		User: user,
		Active: true,
	}
	c.DB.Create(&token)
	return token, nil
}

func (c *Controller) GetTokenFromToken(t string) (models.Token, error) {
	var token models.Token
	result := c.DB.Where("token = ? AND active = true", t).First(&token)
	if result.Error != nil {
		return models.Token{}, result.Error
	}

	fmt.Println(token)
	return token, nil
}
