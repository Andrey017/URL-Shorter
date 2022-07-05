package handler

import "github.com/gin-gonic/gin"

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.LoadHTMLGlob("/home/andrey/go/web_client/templates/*")

	router.GET("/", getHome)

	userRoutes := router.Group("/u")
	{
		userRoutes.GET("/login", showLoginPage)
		userRoutes.POST("/login", performLogin)

		userRoutes.GET("/logout", logoutUser)

		userRoutes.GET("/reg", showRegistrationPage)
		userRoutes.POST("/reg", registrationUser)
	}

	url := router.Group("/api")
	{
		url.GET("/urls/:url_id", getUrlId)

		url.GET("/create", showURLCreate)
		url.POST("/create", createURL)

		url.GET("/update/:url_id", showUpdateURL)
		url.POST("/update/:url_id", updateURL)

		url.GET("/delete/:url_id", deleteURL)
	}

	return router
}
