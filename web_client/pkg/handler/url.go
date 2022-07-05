package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func getHome(c *gin.Context) {
	token, _ := c.Cookie("token")

	if token != "" {
		urls := loadURL(token)

		c.HTML(http.StatusOK, "index.html", gin.H{
			"title":   "Главная страница",
			"logged":  true,
			"payload": urls,
		})
	} else {
		c.Redirect(http.StatusMovedPermanently, "/u/login")
	}
}

func loadURL(token string) []Urls {
	var urlList []Urls

	var bearer = "Bearer " + token

	req, err := http.NewRequest("GET", "http://192.168.17.199:8011/api/urls/", nil)

	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Add("Authorization", bearer)

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		log.Fatalln(err)
	}

	b, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}

	err = json.Unmarshal(b, &urlList)

	if err != nil {
		fmt.Println(err)
	}

	return urlList
}

//-----------------------------------------------CREATE URL-------------------------------------------------------
func showURLCreate(c *gin.Context) {
	c.HTML(http.StatusOK, "create-url.html", gin.H{
		"title":  "Создать новую ссылку",
		"logged": true,
	})
}

func createURL(c *gin.Context) {
	name := c.PostForm("name")
	real_url := c.PostForm("real_url")

	token, _ := c.Cookie("token")

	if token != "" {

		message := map[string]interface{}{
			"name":     name,
			"real_url": real_url,
		}

		bytesRepres, err := json.Marshal(message)

		if err != nil {
			log.Fatalln(err)
		}

		var bearer = "Bearer " + token

		req, err := http.NewRequest("POST", "http://192.168.17.199:8011/api/urls/", bytes.NewBuffer(bytesRepres))

		if err != nil {
			log.Fatalln(err)
		}

		req.Header.Add("Authorization", bearer)

		client := &http.Client{}

		_, err = client.Do(req)

		if err != nil {
			log.Fatalln(err)
		}

		c.Redirect(http.StatusMovedPermanently, "/")
	} else {
		c.Redirect(http.StatusMovedPermanently, "/u/login")
	}
}

//-----------------------------------------------LOAD URL BY ID-------------------------------------------------------
func getUrlId(c *gin.Context) {
	if urlId, err := strconv.Atoi(c.Param("url_id")); err == nil {
		token, _ := c.Cookie("token")

		if token != "" {
			url_list := loadURLID(urlId, token)

			statList := loadURLStat(urlId, token) //`[{"value": 13,"date":1655563064000},{"value":1,"date":1655475670000}]`
			statBrowser := loadBrowserStat(urlId, token)
			statDataOS := loadDataOS(urlId, token)
			statDataType := loadDataType(urlId, token)
			statCountIsMobile := loadCountIsMobile(urlId, token)
			statCountIsTab := loadCountIsTab(urlId, token)
			countRegion := loadCountRegion(urlId, token)

			var namePng string

			var dataIsMobile int
			var dataIsTab int

			for _, temp := range statDataType {
				if temp.IsMobile {
					dataIsMobile = temp.Count
				} else {
					dataIsTab = temp.Count
				}
			}

			totalCount := dataIsMobile + dataIsTab

			hashId := strings.Split(url_list.Shorter_url, "/")

			for _, val := range hashId {
				namePng = val
			}

			c.HTML(http.StatusOK, "url.html", gin.H{
				"title":             url_list.Name,
				"logged":            true,
				"namePng":           namePng,
				"payload":           url_list,
				"stat":              statList,
				"statBrowser":       statBrowser,
				"statDataOS":        statDataOS,
				"statDataIsMobile":  dataIsMobile,
				"statDataIsTab":     dataIsTab,
				"statCountIsMobile": statCountIsMobile,
				"statCountIsTab":    statCountIsTab,
				"countRegion":       countRegion,
				"totalCount":        totalCount,
			})
		} else {
			c.Redirect(http.StatusMovedPermanently, "/u/login")
		}
	}
}

func loadURLID(url_id int, token string) Urls {
	var url_list Urls

	var bearer = "Bearer " + token

	url := fmt.Sprintf("http://192.168.17.199:8011/api/urls/%d", url_id)

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Add("Authorization", bearer)

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		log.Fatalln(err)
	}

	b, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}

	err = json.Unmarshal(b, &url_list)

	if err != nil {
		fmt.Println(err)
	}

	return url_list
}

//-----------------------------------------------UPDATE URL-------------------------------------------------------
func showUpdateURL(c *gin.Context) {

	if urlId, err := strconv.Atoi(c.Param("url_id")); err == nil {
		token, _ := c.Cookie("token")

		if token != "" {

			url_list := loadURLID(urlId, token)

			c.HTML(http.StatusOK, "update-url.html", gin.H{
				"title":   "Изменить ссылку",
				"payload": url_list,
				"logged":  true,
			})
		} else {
			c.Redirect(http.StatusMovedPermanently, "/u/login")
		}
	}
}

func updateURL(c *gin.Context) {
	name := c.PostForm("name")
	real_url := c.PostForm("real_url")

	token, _ := c.Cookie("token")

	urlId, _ := strconv.Atoi(c.Param("url_id"))

	respUpdateURL(name, real_url, urlId, token)

	c.Redirect(http.StatusMovedPermanently, "/")
}

func respUpdateURL(name, real_url string, url_id int, token string) {
	message := map[string]interface{}{
		"name":     name,
		"real_url": real_url,
	}

	bytesRepres, err := json.Marshal(message)

	if err != nil {
		log.Fatalln(err)
	}

	var bearer = "Bearer " + token

	url := fmt.Sprintf("http://192.168.17.199:8011/api/urls/%d", url_id)

	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(bytesRepres))

	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Add("Authorization", bearer)

	client := &http.Client{}

	_, err = client.Do(req)

	if err != nil {
		log.Fatalln(err)
	}
}

//-----------------------------------------------DELETE URL-------------------------------------------------------
func deleteURL(c *gin.Context) {
	token, _ := c.Cookie("token")

	if token != "" {

		urlId, _ := strconv.Atoi(c.Param("url_id"))

		responseDeleteURL(urlId, token)

		c.Redirect(http.StatusMovedPermanently, "/")
	} else {
		c.Redirect(http.StatusMovedPermanently, "/u/login")
	}
}

func responseDeleteURL(url_id int, token string) {
	var bearer = "Bearer " + token

	url := fmt.Sprintf("http://192.168.17.199:8011/api/urls/%d", url_id)

	req, err := http.NewRequest(http.MethodDelete, url, nil)

	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Add("Authorization", bearer)

	client := &http.Client{}

	_, err = client.Do(req)

	if err != nil {
		log.Fatalln(err)
	}
}
