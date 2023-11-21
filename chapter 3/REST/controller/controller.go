package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type controller struct {
}

func NewController(r *gin.Engine) {
	route := r.Group("/v1/api")

	c := controller{}

	route.GET("/", c.getHandler)
}

func (c *controller) getHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"response": "ok",
	})
}
