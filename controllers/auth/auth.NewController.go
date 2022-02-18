package auth

import "check/services/auth"

func NewController() Controller {
	return &controller{
		service: auth.NewService(),
	}
}
