package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type StatsList struct {
	IsMobile bool `json:"ismobile"`
	Count    int  `json:"count"`
}

func loadURLStat(url_id int, token string) string {
	var bearer = "Bearer " + token

	stat := fmt.Sprintf("http://192.168.17.199:8011/api/stat/count/%d", url_id)

	req, err := http.NewRequest("GET", stat, nil)

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

	statList := string(b)

	return statList
}

func loadBrowserStat(url_id int, token string) string {
	var bearer = "Bearer " + token

	stat := fmt.Sprintf("http://192.168.17.199:8011/api/stat/browser/%d", url_id)

	req, err := http.NewRequest("GET", stat, nil)

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

	statList := string(b)

	return statList
}

func loadDataOS(url_id int, token string) string {
	var bearer = "Bearer " + token

	stat := fmt.Sprintf("http://192.168.17.199:8011/api/stat/os/%d", url_id)

	req, err := http.NewRequest("GET", stat, nil)

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

	statList := string(b)

	return statList
}

func loadCountIsMobile(url_id int, token string) string {
	var bearer = "Bearer " + token

	stat := fmt.Sprintf("http://192.168.17.199:8011/api/stat/countIsMobile/%d", url_id)

	req, err := http.NewRequest("GET", stat, nil)

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

	statList := string(b)

	return statList
}

func loadCountIsTab(url_id int, token string) string {
	var bearer = "Bearer " + token

	stat := fmt.Sprintf("http://192.168.17.199:8011/api/stat/countIsTab/%d", url_id)

	req, err := http.NewRequest("GET", stat, nil)

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

	statList := string(b)

	return statList
}

func loadDataType(url_id int, token string) []StatsList {
	var statList []StatsList
	var bearer = "Bearer " + token

	stat := fmt.Sprintf("http://192.168.17.199:8011/api/stat/countType/%d", url_id)

	req, err := http.NewRequest("GET", stat, nil)

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

	err = json.Unmarshal(b, &statList)

	if err != nil {
		fmt.Println(err)
	}

	return statList
}

func loadCountRegion(url_id int, token string) string {
	var bearer = "Bearer " + token

	stat := fmt.Sprintf("http://192.168.17.199:8011/api/stat/countRegion/%d", url_id)

	req, err := http.NewRequest("GET", stat, nil)

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

	statList := string(b)

	return statList
}
