package models

import (
	"container/heap"
	"fmt"
)

type LoadBalancer struct {
	Servers     map[string]QueueItem // map of server address to the Item
	RRNext      int                  // will tell which is the next server to return in a round robin fashion
	Connections PriorityQueue        // to return the server wiht the least connections in a dynamic load balancer
}

func NewLoadBalancer() *LoadBalancer {
	lb := &LoadBalancer{
		Servers:     make(map[string]QueueItem),
		RRNext:      -1,
		Connections: make(PriorityQueue, 0),
	}
	heap.Init(&lb.Connections)
	return lb
}

func (lb *LoadBalancer) AddServer(connections int, serverAdd string) error {
	_, contains := lb.Servers[serverAdd]
	if !contains {
		item := NewQueueItem(connections, serverAdd)
		lb.Servers[serverAdd] = *item
		heap.Push(&lb.Connections, item)
		return nil
	}
	return fmt.Errorf("error: This server already exists")
}

func (lb *LoadBalancer) GetServerWithMinimumServer() *QueueItem {
	item := lb.Connections.GetItemMinConnections()
	delete(lb.Servers, item.ServerAddr)
	return item
}
