package service

import (
	"micro-auth/domain"
	"micro-auth/serializer"
)

type UserService interface {
	Register(reqData *serializer.SignupRequest) *domain.Error
}

type UserServiceImpl struct{}

func (us UserServiceImpl) Register(reqData *serializer.SignupRequest) *domain.Error {
	return nil
}
