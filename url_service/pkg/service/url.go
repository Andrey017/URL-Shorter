package service

import (
	"errors"
	"time"
	"url_service"
	"url_service/pkg/repository"

	"github.com/dgrijalva/jwt-go"
	"github.com/skip2/go-qrcode"
	"github.com/speps/go-hashids"
	"github.com/spf13/viper"
)

const signingKey = ("fhkjsdbfhwbfwe7823n")

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

type URLService struct {
	repo repository.URLS
}

func NewURLService(repo repository.URLS) *URLService {
	return &URLService{
		repo: repo,
	}
}

func generateShortURL(url string) string {
	var shortURL string

	hd := hashids.NewData()
	hd.Salt = url

	h, _ := hashids.NewWithData(hd)

	now := time.Now()

	shortURL, _ = h.Encode([]int{int(now.Unix())})

	return shortURL
}

func (s *URLService) CreateURL(userId int, url url_service.Urls) (int, string, error) {
	url.User_id = userId

	url.Shorter_url = generateShortURL(url.Real_url)

	return s.repo.CreateURL(url)
}

func (s *URLService) GetURLHash(hashURl, userAgent, ipClient string) (url_service.Urls, error) {
	url, err := s.repo.GetURLHash(hashURl)

	if err != nil {
		return url, err
	}

	if url.Id != 0 {
		var region url_service.RegionsCode

		stat, statDetail := parseUserAgent(userAgent)

		region.Code = parseIP(ipClient)

		//idRegion := s.repo.LoadIdRegion(region)

		stat.Url_id = url.Id
		statDetail.City = s.repo.LoadIdRegion(region)

		_, err = s.repo.AddStat(stat, statDetail)

		if err != nil {
			return url, err
		}
	}

	return url, nil
}

func (s *URLService) GetAllURL(userId int) ([]url_service.Urls, error) {
	return s.repo.GetAllURL(userId)
}

func (s *URLService) GetURLById(id, userId int) (url_service.Urls, error) {
	return s.repo.GetURLById(id, userId)
}

func (s *URLService) UpdateURL(id, userId int, url url_service.UpdateURL) error {
	if err := url.Validate(); err != nil {
		return err
	}

	shortURL := generateShortURL(*url.Real_url)

	url.Shorter_url = &shortURL

	return s.repo.UpdateURL(id, userId, url)
}

func (s *URLService) DeleteURL(id, userId int) error {
	return s.repo.DeleteURL(id, userId)
}

func (s *URLService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(accessToken *jwt.Token) (interface{}, error) {
		if _, ok := accessToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Invalid signin method")
		}

		return []byte(signingKey), nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)

	if !ok {
		return 0, errors.New("Token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}

func (s *URLService) GenerateQRURL(shortURL string) error {
	name_file := "qrCode/" + shortURL + ".png"
	err := qrcode.WriteFile(viper.GetString("server.url_server")+shortURL, qrcode.Medium, 256, name_file)

	return err
}
