package service

import (
	"auth_service"
	"auth_service/pkg/repository"
)

type Authorization interface {
	CreateUser(user auth_service.User) (int, error)
	GenerateToken(username, password string) (string, error)
	//ParseToken(accessToken string) (int, error)
}

type Service struct {
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
