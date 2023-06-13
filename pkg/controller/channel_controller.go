package controller

import (
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/valeelim/mahchat/pkg/dto"
	"github.com/valeelim/mahchat/pkg/service"
)

func (c *Controller) CreateChannelController(svc service.CreateChannel) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req dto.CreateChannelRequest
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

func (c *Controller) GetChannelByIDController(svc service.GetChannel) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		resp, err := svc(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, resp)
	}
}

func (c *Controller) ServeChannelController(svc service.ServeChannel) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var wg sync.WaitGroup
		wg.Add(1)
		close := svc(ctx, &wg)
		wg.Wait()
		defer func() {
			log.Println("controller closed")
			close()
		}()
	}
}

func (c *Controller) CreateGroupMessageController(svc service.CreateGroupMessage) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req dto.CreateGroupMessageRequest
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		if err := svc(ctx, req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Message created",
		})
	}
}

func (c *Controller) GetGroupMessageByChannelIDController(svc service.GetGroupMessages) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		resp, err := svc(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, resp)
	}
}
