package models

import (
	"errors"

	"gorm.io/gorm"
)

type Token struct {
	gorm.Model
	Token     string        `json:"token"`
	UserID    uint
	User      User          `json:"-"`
	Active	  bool
}

func (t Token) Validete() error {
	if t.IsExpired() {
		return errors.New("token is expired")
	}
	if len(t.Token) != 64 {
		return errors.New("invalid token")
	}

	return nil
}

func (t Token) IsExpired() bool {
	// TODO: Implement
	return false
}
