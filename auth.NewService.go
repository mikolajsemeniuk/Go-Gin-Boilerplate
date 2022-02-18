package auth

import "check/repositories/auth"

func NewService() Service {
	return &service{
		repository: auth.NewRepository(),
	}
}
