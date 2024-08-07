package test

import (
	"testing"

	"github.com/Uttkarsh-raj/go-load-balancer/models"
	"github.com/stretchr/testify/assert"
)

func TestRoundRobinTest(t *testing.T) {
	lb := models.NewLoadBalancer()
	servers := [3]string{"http://testServer1.com", "http://testServer2.com", "http://testServer3.com"}
	lb.AddServer(0, servers[0])
	lb.AddServer(0, servers[1])
	lb.AddServer(0, servers[2])
	for i := 0; i <= 5; i++ {
		add := lb.GetRoundRobinNextServerAddr()
		assert.Equal(t, add, servers[i%3])
	}
}

func TestDeleteWhileInRoundRobin(t *testing.T) {
	lb := models.NewLoadBalancer()
	servers := [3]string{"http://testServer1.com", "http://testServer2.com", "http://testServer3.com"}
	lb.AddServer(0, servers[0])
	lb.AddServer(0, servers[1])
	lb.AddServer(0, servers[2])
	for i := 0; i < 2; i++ {
		add := lb.GetRoundRobinNextServerAddr()
		assert.Equal(t, add, servers[i%3])
	}
	err := lb.DeleteServer(servers[2]) // If the server gets deleted before getting the request
	assert.NoError(t, err, nil)

	add := lb.GetRoundRobinNextServerAddr()
	assert.Equal(t, add, servers[0])
}
