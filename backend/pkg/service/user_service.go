package service

import (
	"errors"
	"time"

	"github.com/SoumyaJha0410/backend/pkg/domain"
	"github.com/SoumyaJha0410/backend/pkg/repository"
	"github.com/SoumyaJha0410/backend/pkg/service/dto"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type IUserService interface {
	Login(username, password string) (string, error)
	SignUp(user dto.UserCreate) error
}

type UserService struct {
	userRepository repository.IUserRepository
}

func NewUserService(userRepository repository.IUserRepository) IUserService {
	return &UserService{userRepository}
}

func (service *UserService) Login(username, password string) (string, error) {
	jwtKey := "JWT_KEY"
	user,err:= service.userRepository.GetUserByUsername(username)
	if err != nil {
		return "", errors.New("no user found with the username: " + username)
	}

	err=bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(password))
	if err != nil {
		return "", errors.New("invalid password")
	}

	expirationTime:=time.Now().Add(24* time.Hour)
	claims:= &domain.Claims{
		Username: user.Username,
		Role: user.Role,
	    RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token:=jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	tokenString,err:= token.SignedString([]byte(jwtKey))
	if err!=nil{
		return "", errors.New("error signing the token: " + err.Error())
	}
	return tokenString, nil
}

func (service *UserService) SignUp(userCreate dto.UserCreate) error {
	err := validateUser(userCreate)
	if err!=nil{
		return err
	}
	user:=userCreateToUser(userCreate)
	hashedPassword,err:= bcrypt.GenerateFromPassword([]byte(user.Password),bcrypt.DefaultCost)
    if err!=nil{
		return err
	}

	user.Password = string(hashedPassword)
	return service.userRepository.SignUp(user)

}

func validateUser(user dto.UserCreate) error {
	if user.Username == "" {
		return errors.New("username can't be empty")
	}
	if user.Password == "" {
		return errors.New("category can't be empty")
	}
	if user.Role == "" {
		return errors.New("quantity can't be empty")
	}
	return nil
}

func userCreateToUser(userCreate dto.UserCreate) domain.User {
	return domain.User{
		Username: userCreate.Username,
		Password: userCreate.Password,
		Role:     userCreate.Role,
	}
}
