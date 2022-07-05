package handler

import (
	"net/http"
	"strconv"
	"url_service"

	"github.com/gin-gonic/gin"
)

func (h *Handler) RedirectURL(c *gin.Context) {
	hashURL := c.Param("hashURL")

	userAgent := c.Request.UserAgent()

	ipClient := c.Request.Header.Get("X-Forwarded-For")

	url, err := h.service.GetURLHash(hashURL, userAgent, ipClient)

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.Redirect(http.StatusMovedPermanently, url.Real_url)
}

func (h *Handler) CreateURL(c *gin.Context) {
	userId, err := h.getUserId(c)

	if err != nil {
		return
	}

	var input url_service.Urls

	errBind := c.BindJSON(&input)

	if errBind != nil {
		newErrorResponse(c, http.StatusBadRequest, errBind.Error())
		return
	}

	id, hashURL, err := h.service.URLS.CreateURL(userId, input)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id":      id,
		"hashURL": hashURL,
	})
}

func (h *Handler) GetAllURL(c *gin.Context) {
	userId, err := h.getUserId(c)

	if err != nil {
		return
	}

	urlList, err := h.service.URLS.GetAllURL(userId)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, urlList)
}

func (h *Handler) GetURLById(c *gin.Context) {
	userId, err := h.getUserId(c)

	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid id param")
		return
	}

	url, err := h.service.URLS.GetURLById(id, userId)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, url)
}

func (h *Handler) UpdateURL(c *gin.Context) {
	userId, err := h.getUserId(c)

	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid id param")
		return
	}

	var input url_service.UpdateURL

	if errBind := c.BindJSON(&input); errBind != nil {
		newErrorResponse(c, http.StatusInternalServerError, errBind.Error())
		return
	}

	if errUpdate := h.service.URLS.UpdateURL(id, userId, input); errUpdate != nil {
		newErrorResponse(c, http.StatusInternalServerError, errUpdate.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Message: "OK",
	})
}

func (h *Handler) DeleteURL(c *gin.Context) {
	userId, err := h.getUserId(c)

	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid id param")
		return
	}

	if errDelete := h.service.URLS.DeleteURL(id, userId); errDelete != nil {
		newErrorResponse(c, http.StatusInternalServerError, errDelete.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Message: "OK",
	})
}
