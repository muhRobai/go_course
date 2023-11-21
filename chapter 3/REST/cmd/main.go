package main

import (
	"rest/controller"
	httppkg "rest/pkg/http"
	"rest/pkg/http/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	middleware.HealthCheck("note", router)

	controller.NewController(router)

	httppkg.Serve(router, ":8000")
}
