package service

import "github.com/anggardagasta/mini_wallet/models"

type IServiceUsersUseCase interface {
	RegisterUser(form models.FormRegister) (result models.AuthResult, err error)
	Auth(form models.FormAuth) (result models.AuthResult, err error)
}
