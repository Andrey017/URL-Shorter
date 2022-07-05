package handler

import (
	"auth_service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SignInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) SignUp(c *gin.Context) {
	var input auth_service.User

	errBinding := c.BindJSON(&input)

	if errBinding != nil {
		newErrorResponse(c, http.StatusBadRequest, errBinding.Error())
	}

	if input.Name != "" && input.Username != "" && input.Password != "" {

		id, err := h.service.Authorization.CreateUser(input)

		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
		}

		c.JSON(http.StatusOK, map[string]interface{}{
			"id": id,
		})
	} else {
		c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Error empty login or password",
		})
	}
}

func (h *Handler) SignIn(ctx *gin.Context) {
	var input SignInInput

	errBinding := ctx.BindJSON(&input)

	if errBinding != nil {
		newErrorResponse(ctx, http.StatusBadRequest, errBinding.Error())
	}

	if input.Username != "" && input.Password != "" {

		token, err := h.service.Authorization.GenerateToken(input.Username, input.Password)

		if err != nil {
			newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		}

		ctx.JSON(http.StatusOK, map[string]interface{}{
			"token": token,
		})
	} else {
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"message": "Error empty login or password",
		})
	}
}
