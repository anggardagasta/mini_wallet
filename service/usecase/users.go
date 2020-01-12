package usecase

import (
	"encoding/base64"
	"errors"
	"github.com/anggardagasta/mini_wallet/models"
	"github.com/anggardagasta/mini_wallet/service"
	"github.com/anggardagasta/mini_wallet/service/repository/constant"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type serviceUsersUsecase struct {
	serviceUsersRepo service.IServiceUsersRepository
}

func NewServiceUsersUsecase(serviceUserRepo service.IServiceUsersRepository) service.IServiceUsersUseCase {
	return serviceUsersUsecase{serviceUsersRepo: serviceUserRepo}
}

func (uc serviceUsersUsecase) GeneratingToken(userID int64, username string, email string) (result string, err error) {
	claims := models.MyClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    constant.APPLICATION_NAME,
			ExpiresAt: time.Now().Add(constant.LOGIN_EXPIRATION_DURATION).Unix(),
		},
		ID:       userID,
		Username: username,
		Email:    email,
	}
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)
	signedToken, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (uc serviceUsersUsecase) RegisterUser(form models.FormRegister) (result models.AuthResult, err error) {
	checkUser, err := uc.serviceUsersRepo.GetUserByUsernameOrEmail(form)
	if err != nil {
		return result, err
	}
	if checkUser.ID.Int64 > int64(0) {
		return result, errors.New("username or email is already exist")
	}

	form.Password = base64.StdEncoding.EncodeToString([]byte(form.Password))
	userID, err := uc.serviceUsersRepo.InsertUser(form)
	if err != nil {
		return result, err
	}

	token, err := uc.GeneratingToken(userID, form.Username, form.Email)
	if err != nil {
		return result, err
	}

	result.Token = token

	return result, nil
}

func (uc serviceUsersUsecase) Auth(form models.FormAuth) (result models.AuthResult, err error) {
	var user models.GetUserScanner
	if form.Username != "" {
		getUser, err := uc.serviceUsersRepo.GetUserByUsername(form.Username)
		if err != nil {
			return result, err
		}
		user = getUser
	} else {
		getUser, err := uc.serviceUsersRepo.GetUserByEmail(form.Email)
		if err != nil {
			return result, err
		}
		user = getUser
	}
	if user.ID.Int64 == 0 {
		return result, errors.New("user not found")
	}

	userPassword, err := base64.StdEncoding.DecodeString(user.Password.String)
	if err != nil {
		return result, err
	}
	if form.Password == string(userPassword) {
		token, err := uc.GeneratingToken(user.ID.Int64, user.Username.String, user.Password.String)
		if err != nil {
			return result, err
		}

		result.Token = token
	} else {
		return result, errors.New("invalid password")
	}
	return result, nil
}
