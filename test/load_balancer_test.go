package test

import (
	"testing"

	"github.com/Uttkarsh-raj/go-load-balancer/models"
	"github.com/stretchr/testify/assert"
)

func TestAddServer(t *testing.T) {
	lb := models.NewLoadBalancer()
	lb.AddServer(0, "http://testServer1.com")
	lb.AddServer(1, "http://testServer2.com")
	lb.AddServer(2, "http://testServer3.com")
	assert.Equal(t, len(lb.ServerAddr), 3)
}
func TestAddSameServer(t *testing.T) {
	lb := models.NewLoadBalancer()
	lb.AddServer(0, "http://testServer.com")
	lb.AddServer(2, "http://testServer.com")
	assert.Equal(t, len(lb.ServerAddr), 1)
}

func TestDeleteServer(t *testing.T) {
	lb := models.NewLoadBalancer()
	lb.AddServer(0, "http://testServer.com")
	err := lb.DeleteServer("http://testServer.com")
	assert.NoError(t, err, nil)
}

func TestDeleteServerNotExists(t *testing.T) {
	lb := models.NewLoadBalancer()
	err := lb.DeleteServer("http://testServer.com")
	assert.Error(t, err, "error: The server does not exists")
}
