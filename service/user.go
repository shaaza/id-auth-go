package service

import (
	"micro-auth/db"
	"micro-auth/domain"
	"micro-auth/serializer"
)

type UserService interface {
	Register(reqData *serializer.SignupRequest) *domain.Error
}

type UserServiceImpl struct {
	Database db.DB
}

func (us UserServiceImpl) Register(reqData *serializer.SignupRequest) *domain.Error {
	user := domain.User{
		Username:    reqData.Username,
		Password:    reqData.Password,
		FirstName:   reqData.FirstName,
		LastName:    reqData.LastName,
		PhoneNumber: reqData.PhoneNumber,
	}
	us.Database.NewUser(user)
	return nil
}
