package main

import (
	"fmt"

	"github.com/Uttkarsh-raj/go-load-balancer/models"
)

func main() {
	balancer := models.NewLoadBalancer()
	balancer.AddServer(5, "item1")
	balancer.AddServer(7, "item2")
	balancer.AddServer(2, "item3")
	balancer.AddServer(8, "item4")
	for i := 0; i < 4; i++ {
		item := balancer.Connections.GetServerWithMinConnections()
		println(balancer.Connections.Len())
		fmt.Printf("The item is %d with %d connections\n", item.ServerIndex, item.Connections)
	}
}
