package repository

import (
	"database/sql"
	"fmt"
	"strings"
	"url_service"

	"github.com/spf13/viper"
)

type URLSQL struct {
	db *sql.DB
}

func NewURLSQL(db *sql.DB) *URLSQL {
	return &URLSQL{
		db: db,
	}
}

func (r *URLSQL) CreateURL(urls url_service.Urls) (int, string, error) {
	var idResult int64

	query := fmt.Sprintf("INSERT INTO %s (name, real_url, shorter_url, user_id) VALUES ('%s', '%s', '%s', %d)", urlTable, urls.Name, urls.Real_url, urls.Shorter_url, urls.User_id)

	result, err := r.db.Exec(query)

	if err != nil {
		return 0, "", err
	}

	idResult, err = result.LastInsertId()

	if err != nil {
		return 0, "", err
	}

	return int(idResult), urls.Shorter_url, nil
}

func (r *URLSQL) GetURLHash(hashURL string) (url_service.Urls, error) {
	var url url_service.Urls

	query := fmt.Sprintf("SELECT id, real_url FROM %s WHERE shorter_url='%s'", urlTable, hashURL)

	rows, err := r.db.Query(query)

	for rows.Next() {
		rows.Scan(&url.Id, &url.Real_url)
	}

	return url, err
}

func (r *URLSQL) GetAllURL(userId int) ([]url_service.Urls, error) {
	var urlList []url_service.Urls

	query := fmt.Sprintf("SELECT id, name, real_url, shorter_url FROM %s WHERE user_id=%d", urlTable, userId)

	rows, err := r.db.Query(query)

	for rows.Next() {
		var singleList url_service.Urls

		errScan := rows.Scan(&singleList.Id, &singleList.Name, &singleList.Real_url, &singleList.Shorter_url)

		if errScan != nil {
			return urlList, errScan
		}

		singleList.Shorter_url = viper.GetString("server.url_server") + singleList.Shorter_url

		urlList = append(urlList, singleList)
	}

	return urlList, err
}

func (r *URLSQL) GetURLById(id, userId int) (url_service.Urls, error) {
	var url url_service.Urls

	query := fmt.Sprintf("SELECT id, name, real_url, shorter_url FROM %s WHERE id=%d AND user_id=%d", urlTable, id, userId)

	rows, err := r.db.Query(query)

	for rows.Next() {
		rows.Scan(&url.Id, &url.Name, &url.Real_url, &url.Shorter_url)

		url.Shorter_url = viper.GetString("server.url_server") + url.Shorter_url
	}

	return url, err
}

func (r *URLSQL) UpdateURL(id, userId int, url url_service.UpdateURL) error {
	setValues := make([]string, 0)

	if url.Name != nil {
		setValues = append(setValues, fmt.Sprint("name='", *url.Name, "'"))
	}

	if url.Real_url != nil {
		setValues = append(setValues, fmt.Sprint("real_url='", *url.Real_url, "'"))
	}

	if url.Shorter_url != nil {
		setValues = append(setValues, fmt.Sprint("shorter_url='", *url.Shorter_url, "'"))
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id=%d AND user_id=%d", urlTable, setQuery, id, userId)

	_, err := r.db.Exec(query)

	return err
}

func (r *URLSQL) DeleteURL(id, userId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=%d AND user_id=%d", urlTable, id, userId)

	_, err := r.db.Exec(query)

	return err
}
