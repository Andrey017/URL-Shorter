package handler

import (
	"url_service/pkg/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/:hashURL", h.RedirectURL)

	api := router.Group("/api", h.userIdentity)
	{
		urls := api.Group("/urls")
		{
			urls.POST("/", h.CreateURL)
			urls.GET("/", h.GetAllURL)
			urls.GET("/:id", h.GetURLById)
			urls.PUT("/:id", h.UpdateURL)
			urls.DELETE("/:id", h.DeleteURL)
		}

		stat := api.Group("/stat")
		{
			statCount := stat.Group("/count")
			{
				statCount.GET("/:id", h.GetStatURLid)
			}

			statBrowser := stat.Group("/browser")
			{
				statBrowser.GET("/:id", h.GetStatBrowser)
			}

			statOs := stat.Group("/os")
			{
				statOs.GET("/:id", h.GetOsURL)
			}

			statCountType := stat.Group("/countType")
			{
				statCountType.GET("/:id", h.GetCountType)
			}

			statCountIsMobile := stat.Group("/countIsMobile")
			{
				statCountIsMobile.GET("/:id", h.GetCountIsMobile)
			}

			statCountIsTab := stat.Group("/countIsTab")
			{
				statCountIsTab.GET("/:id", h.GetCountIsTab)
			}

			statCountRegion := stat.Group("/countRegion")
			{
				statCountRegion.GET("/:id", h.GetCountRegion)
			}
		}
	}

	return router
}
