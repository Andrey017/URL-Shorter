package service

import (
	"url_service"
	"url_service/pkg/repository"
)

type URLS interface {
	CreateURL(userId int, url url_service.Urls) (int, string, error)
	GetURLHash(hashURL, userAgent, ipClient string) (url_service.Urls, error)
	GetAllURL(userId int) ([]url_service.Urls, error)
	GetURLById(id, userId int) (url_service.Urls, error)
	UpdateURL(id, userId int, url url_service.UpdateURL) error
	DeleteURL(id, userId int) error

	GenerateQRURL(shortURL string) error

	ParseToken(accessToken string) (int, error)

	GetStatURLid(urlId int) ([]url_service.Stats, error)
	GetStatBrowser(urlId int) ([]url_service.StatsBrowser, error)
	GetOsURL(urlId int) ([]url_service.StatsOs, error)
	GetCountType(urlId int) ([]url_service.StatCountType, error)
	GetCountIsMobile(urlId int) ([]url_service.StatCountIsMobile, error)
	GetCountIsTab(urlId int) ([]url_service.StatCountIsMobile, error)
	GetCountRegion(urlId int) ([]url_service.RegionsCode, error)
}

type Service struct {
	URLS
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		URLS: NewURLService(repos.URLS),
	}
}
