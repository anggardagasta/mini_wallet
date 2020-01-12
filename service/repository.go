package service

import "github.com/anggardagasta/mini_wallet/models"

type IServiceUsersRepository interface {
	GetUserByUsernameOrEmail(form models.FormRegister) (result models.GetUserScanner, err error)
	GetUserByUsername(username string) (result models.GetUserScanner, err error)
	GetUserByEmail(email string) (result models.GetUserScanner, err error)
	InsertUser(input models.FormRegister) (id int64, err error)
}
