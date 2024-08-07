package main

import (
	"fmt"
	"log"

	"github.com/Uttkarsh-raj/go-load-balancer/models"
	"github.com/Uttkarsh-raj/go-load-balancer/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Starting Server.....")
	router := gin.New()
	router.Use(gin.Logger())
	balancer := models.NewLoadBalancer()
	routes.SetRoutes(router, balancer)
	log.Fatal(router.Run(":3000"))
}
