package routes

import (
	"github.com/Uttkarsh-raj/go-load-balancer/controller"
	"github.com/Uttkarsh-raj/go-load-balancer/models"
	"github.com/gin-gonic/gin"
)

func SetRoutes(router *gin.Engine, loadBalancer *models.LoadBalancer) {
	router.Any("/*path", controller.ForwardRequest(loadBalancer))
}
