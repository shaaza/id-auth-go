package service

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
	"micro-auth/db"
	"micro-auth/domain"
	"micro-auth/serializer"
	"net/http"
)

const HASHING_COST = 10

type UserService interface {
	Register(reqData *serializer.SignupRequest) *domain.Error
	Login(reqData *serializer.LoginRequest) (domain.User, *domain.Error)
}

type UserServiceImpl struct {
	Database db.DB
}

func (us UserServiceImpl) Register(reqData *serializer.SignupRequest) *domain.Error {
	hashedPass, err := hashPassword(reqData.Password)
	if err != nil {
		log.Fatal(err)
		return domain.NewError(err.Error(), http.StatusBadRequest)
	}

	user := domain.User{
		Username:    reqData.Username,
		Password:    hashedPass,
		FirstName:   reqData.FirstName,
		LastName:    reqData.LastName,
		PhoneNumber: reqData.PhoneNumber,
	}
	us.Database.NewUser(user)
	return nil
}

func (us UserServiceImpl) Login(reqData *serializer.LoginRequest) (domain.User, *domain.Error) {
	user, err := us.Database.GetUser(reqData.Username)
	if err != nil {
		return domain.User{}, domain.NewError(fmt.Sprintf("Login failed: %s", err.Error()), err.Code())
	}

	passwordMatch := checkPasswordHash(user.Password, reqData.Password)
	if passwordMatch != true {
		return domain.User{}, domain.NewError("Invalid password", http.StatusUnauthorized)
	}

	return user, nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), HASHING_COST)
	return string(bytes), err
}

func checkPasswordHash(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
