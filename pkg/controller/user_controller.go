package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/valeelim/mahchat/pkg/dto"
	"github.com/valeelim/mahchat/pkg/service"
)

func (c *Controller) RegisterController(svc service.RegisterUser) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req dto.RegisterRequest
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		resp, err := svc(ctx, req)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, resp)
	}
}

func (c *Controller) LoginController(svc service.LoginUser) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req dto.LoginRequest
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		resp, err := svc(ctx, req)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, resp)
	}
}

func (c *Controller) GetAllUsersController(svc service.GetUsers) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		resp, err := svc(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
		ctx.JSON(http.StatusOK, resp)
	}
}
