package test

import (
	"testing"

	"github.com/Uttkarsh-raj/go-load-balancer/models"
)

func TestAddServer(t *testing.T) {
	lb := models.NewLoadBalancer()
	lb.AddServer(0, "http://testServer.com")
	lb.AddServer(1, "http://testServer.com")
	lb.AddServer(2, "http://testServer.com")

}
