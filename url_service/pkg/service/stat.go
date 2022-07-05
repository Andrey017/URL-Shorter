package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
	"url_service"

	"github.com/pattfy/useragent"
)

type GeoIP struct {
	Response        string `json:"response"`
	CountryCode     string `json:"country_code"`
	CountryName     string `json:"country_name"`
	RegionName      string `json:"region_name"`
	CityName        string `json:"city_name"`
	Region          Region `json:"region"`
	CreditsConsumed int64  `json:"credits_consumed"`
}

type Region struct {
	Name         string       `json:"name"`
	Code         string       `json:"code"`
	Translations Translations `json:"translations"`
}

type Translations struct {
	Ru string `json:"ru"`
}

func parseUserAgent(userAgent string) (url_service.Stats, url_service.StatsDetail) {
	var stat url_service.Stats
	var statDetail url_service.StatsDetail

	dataAgent := useragent.New(userAgent)

	browser := dataAgent.Browser
	platform := dataAgent.Platform
	mobil := dataAgent.IsMobile()

	stat.Counttransition = 1
	stat.Date = time.Now().Unix() * 1000
	stat.DateDB = time.Now().Format("2006-01-02")
	statDetail.City = 0
	statDetail.Browser = browser.Name()
	statDetail.Os = platform.Name()
	statDetail.IsMobile = mobil

	fmt.Println(mobil)
	return stat, statDetail
}

func parseIP(ipClient string) string {
	var geo GeoIP

	response, err := http.Get("https://api.ip2location.com/v2/?ip=" + ipClient + "&package=WS3&addon=continent,country,region,city&lang=ru&key=J6AMVRBRLY")
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(body, &geo)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(geo)

	return geo.Region.Code
}

func (s *URLService) GetStatURLid(urlId int) ([]url_service.Stats, error) {
	return s.repo.GetStatURLid(urlId)
}

func (s *URLService) GetStatBrowser(urlId int) ([]url_service.StatsBrowser, error) {
	return s.repo.GetStatBrowser(urlId)
}

func (s *URLService) GetOsURL(urlId int) ([]url_service.StatsOs, error) {
	return s.repo.GetOsURL(urlId)
}

func (s *URLService) GetCountType(urlId int) ([]url_service.StatCountType, error) {
	return s.repo.GetCountType(urlId)
}

func (s *URLService) GetCountIsMobile(urlId int) ([]url_service.StatCountIsMobile, error) {
	return s.repo.GetCountIsMobile(urlId)
}

func (s *URLService) GetCountIsTab(urlId int) ([]url_service.StatCountIsMobile, error) {
	return s.repo.GetCountIsTab(urlId)
}

func (s *URLService) GetCountRegion(urlId int) ([]url_service.RegionsCode, error) {
	return s.repo.GetCountRegion(urlId)
}
