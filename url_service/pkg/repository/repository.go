package repository

import (
	"database/sql"
	"url_service"
)

type URLS interface {
	CreateURL(url url_service.Urls) (int, string, error)
	GetURLHash(hashURL string) (url_service.Urls, error)
	GetAllURL(userId int) ([]url_service.Urls, error)
	GetURLById(id, userId int) (url_service.Urls, error)
	UpdateURL(id, userId int, url url_service.UpdateURL) error
	DeleteURL(id, userId int) error

	AddStat(stat url_service.Stats, statDetail url_service.StatsDetail) (int, error)
	AddStatDetail(statDetail url_service.StatsDetail) (int, error)
	LoadIdRegion(stat url_service.RegionsCode) int
	CheckStat(stat url_service.Stats) int
	GetStatURLid(urlId int) ([]url_service.Stats, error)
	GetStatBrowser(urlId int) ([]url_service.StatsBrowser, error)
	GetOsURL(urlId int) ([]url_service.StatsOs, error)
	GetCountType(urlId int) ([]url_service.StatCountType, error)
	GetCountIsMobile(urlId int) ([]url_service.StatCountIsMobile, error)
	GetCountIsTab(urlId int) ([]url_service.StatCountIsMobile, error)
	GetCountRegion(urlId int) ([]url_service.RegionsCode, error)
}

type Repository struct {
	URLS
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		URLS: NewURLSQL(db),
	}
}
