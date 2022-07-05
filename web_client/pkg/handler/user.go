package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func showLoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{
		"title":  "Авторизация",
		"logged": false,
	})
}

func performLogin(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	var sameSiteCookie string

	if username != "" && password != "" {
		token := MakeRequest(username, password)

		if token != "" {
			c.SetCookie("token", token, 3600, "", sameSiteCookie, false, true)

			c.Redirect(http.StatusMovedPermanently, "/")
		} else {
			c.HTML(http.StatusOK, "login.html", gin.H{
				"ErrorTitle":   "Ошибка авторизации",
				"ErrorMessage": "Неверный логин и пароль",
			})
		}
	} else {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"ErrorTitle":   "Ошибка авторизации",
			"ErrorMessage": "Неверный логин и пароль",
		})
	}
}

func MakeRequest(email, password string) string {
	message := map[string]interface{}{
		"username": email,
		"password": password,
	}

	bytesRepres, err := json.Marshal(message)

	if err != nil {
		log.Fatalln(err)
	}

	resp, err := http.Post("http://192.168.17.199:8011/auth/sign-in", "application/json", bytes.NewBuffer(bytesRepres))

	if err != nil {
		log.Fatalln(err)
	}

	var result map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&result)

	var token string

	for _, val := range result {
		token = fmt.Sprint("", val)
	}

	return token
}

func showRegistrationPage(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", gin.H{
		"title": "Регистрация",
	})
}

func registrationUser(c *gin.Context) {
	name := c.PostForm("name")
	username := c.PostForm("username")
	password := c.PostForm("password")

	if name != "" && username != "" && password != "" {

		respRegUser(name, username, password)

		showLoginPage(c)

		c.Redirect(http.StatusMovedPermanently, "/u/login")
	} else {
		c.HTML(http.StatusOK, "register.html", gin.H{
			"ErrorTitle":   "Ошибка регистрации",
			"ErrorMessage": "Заполните все поля",
		})
	}
}

func respRegUser(name, username, password string) {
	message := map[string]interface{}{
		"name":     name,
		"username": username,
		"password": password,
	}

	bytesRepres, err := json.Marshal(message)

	if err != nil {
		log.Fatalln(err)
	}

	req, err := http.NewRequest("POST", "http://192.168.17.199:8011/auth/sign-up/", bytes.NewBuffer(bytesRepres))

	if err != nil {
		log.Fatalln(err)
	}

	client := &http.Client{}

	_, err = client.Do(req)

	if err != nil {
		log.Fatalln(err)
	}
}

func logoutUser(c *gin.Context) {
	var sameSiteCookie string

	c.SetCookie("token", "", 10, "", sameSiteCookie, false, true)
	c.Redirect(http.StatusMovedPermanently, "/u/login")
}
