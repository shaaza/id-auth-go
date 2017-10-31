package service

import (
	"micro-auth/serializer"
	"micro-auth/domain"
)

type UserService interface {
	Register(reqData *serializer.SignupRequest) (string, *domain.Error)
}

type UserServiceImpl struct {}


func (us UserServiceImpl) Register(reqData *serializer.SignupRequest) (string, *domain.Error) {
	return "user", nil
}
