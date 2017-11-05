package service

import (
	"fmt"
	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"log"
	"micro-auth/db"
	"micro-auth/domain"
	"micro-auth/serializer"
	"net/http"
	"time"
)

const HASHING_COST = 10

type UserService interface {
	Register(reqData *serializer.SignupRequest) *domain.Error
	Login(reqData *serializer.LoginRequest) (domain.User, domain.Session, *domain.Error)
	Logout(sessionID string) *domain.Error
}

type UserServiceImpl struct {
	Database db.DB
}

func (us UserServiceImpl) Register(reqData *serializer.SignupRequest) *domain.Error {
	user, _ := us.Database.GetUser(reqData.Username)
	if user.Username != "" {
		return domain.NewError("user already exists", http.StatusConflict)
	}

	hashedPass, err := hashPassword(reqData.Password)
	if err != nil {
		log.Fatal(err)
		return domain.NewError(err.Error(), http.StatusBadRequest)
	}

	user = domain.User{
		Id:          uuid.NewV4().String(),
		Username:    reqData.Username,
		Password:    hashedPass,
		FirstName:   reqData.FirstName,
		LastName:    reqData.LastName,
		PhoneNumber: reqData.PhoneNumber,
	}
	rowCnt, registerErr := us.Database.NewUser(user)
	if registerErr != nil {
		return domain.NewError(fmt.Sprintf("could not create user: %s", registerErr.Error()), http.StatusInternalServerError)
	}

	if rowCnt == 0 {
		return domain.NewError(fmt.Sprintf("could not create user: %s", registerErr.Error()), http.StatusInternalServerError)
	}

	return nil
}

func (us UserServiceImpl) Login(reqData *serializer.LoginRequest) (domain.User, domain.Session, *domain.Error) {
	user, err := us.Database.GetUser(reqData.Username)
	if err != nil {
		return domain.User{}, domain.Session{}, domain.NewError(fmt.Sprintf("Login failed: %s", err.Error()), err.Code())
	}

	passwordMatch := checkPasswordHash(user.Password, reqData.Password)
	if passwordMatch != true {
		return domain.User{}, domain.Session{}, domain.NewError("Invalid password", http.StatusUnauthorized)
	}

	session := domain.Session{
		Id:        uuid.NewV4().String(),
		UserId:    user.Id,
		StartTime: time.Now(),
		Expiry:    int(time.Duration(time.Hour * 24 * 30).Seconds()),
		Valid:     true,
	}

	existingSession, err := us.Database.GetValidSession(user.Id)
	if err != nil && err.Code() != http.StatusNotFound {
		return domain.User{}, domain.Session{}, domain.NewError(fmt.Sprintf("Login failed: %s", err.Error()), err.Code())
	}

	if len(existingSession.Id) > 0 {
		return user, existingSession, &domain.Error{}
	}

	rowCnt, loginErr := us.Database.NewSession(session)
	if loginErr != nil {
		return domain.User{}, domain.Session{}, domain.NewError(fmt.Sprintf("could not create session: %s", loginErr.Error()), loginErr.Code())
	}

	if rowCnt == 0 {
		return domain.User{}, domain.Session{}, domain.NewError(fmt.Sprintf("could not create session: %s", loginErr.Error()), http.StatusInternalServerError)
	}

	return user, session, nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), HASHING_COST)
	return string(bytes), err
}

func checkPasswordHash(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (us UserServiceImpl) Logout(sessionID string) *domain.Error {
	rowCnt, loginErr := us.Database.InvalidateSession(sessionID)
	if loginErr != nil {
		return domain.NewError(fmt.Sprintf("could not invalidate session: %s", loginErr.Error()), loginErr.Code())
	}

	if rowCnt == 0 {
		return domain.NewError("session does not exist", http.StatusNotFound)
	}

	return nil
}
