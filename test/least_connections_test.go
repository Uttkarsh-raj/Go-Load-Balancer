package test

import (
	"testing"

	"github.com/Uttkarsh-raj/go-load-balancer/models"
	"github.com/stretchr/testify/assert"
)

func TestGetServerWithLeastConnection(t *testing.T) {
	lb := models.NewLeastConnectionLoadBalancer()
	lb.AddServer(5, "http://testServer1.com")
	lb.AddServer(1, "http://testServer2.com")
	lb.AddServer(3, "http://testServer3.com")
	add := lb.GetServerWithMinimumServerAddr()
	assert.Equal(t, add, "http://testServer2.com")
}

func TestGetServerByAddress(t *testing.T) {
	lb := models.NewLeastConnectionLoadBalancer()
	lb.AddServer(5, "http://testServer1.com")
	lb.AddServer(1, "http://testServer2.com")
	addr := lb.Connections.GetItemByServerAddr("http://testServer1.com")
	assert.Equal(t, addr.ServerAddr, "http://testServer1.com")
	assert.Equal(t, addr.Connections, 5)
}

func TestGetServerWithLeastConnections(t *testing.T) {
	lb := models.NewLeastConnectionLoadBalancer()
	expectedSeq := [12]string{"server4", "server4", "server4", "server2", "server2", "server4", "server4", "server2", "server5", "server5", "server2", "server4"}
	lb.AddServer(10, "server1")
	lb.AddServer(6, "server2")
	lb.AddServer(15, "server3")
	lb.AddServer(4, "server4")
	lb.AddServer(8, "server5")
	for i := 0; i < 12; i++ {
		item := lb.Connections.GetItemMinConnections()
		item.Connections++
		lb.Connections.Update(item, item.Connections)
		assert.Equal(t, item.ServerAddr, expectedSeq[i])
	}
}
