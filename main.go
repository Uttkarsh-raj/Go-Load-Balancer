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
	// balancer.AddServer(5, "item1")
	// balancer.AddServer(7, "item2")
	// balancer.AddServer(2, "item3")
	// balancer.AddServer(8, "item4")
	// for i := 0; i < 4; i++ {
	// 	item := balancer.GetServerWithMinimumServerAddr()
	// 	fmt.Printf("The item is %s with %d connections\n", item.ServerAddr, item.Connections)
	// }
}
