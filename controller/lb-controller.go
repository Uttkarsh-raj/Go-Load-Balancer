package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/Uttkarsh-raj/go-load-balancer/models"
	"github.com/gin-gonic/gin"
)

func AddNewServer(loadBalancer *models.LoadBalancer) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()
		var body map[string]string
		err := json.NewDecoder(ctx.Request.Body).Decode(&body)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Unable to decode the body."})
		}
		addr := body["address"]
		err = loadBalancer.AddServer(0, addr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"success": true, "message": "Server added successfully"})
	}
}
