package models

import "container/heap"

type LoadBalancer struct {
	Servers     []string      // list of the servers
	RRNext      int           // will tell which is the next server to return in a round robin fashion
	Connections PriorityQueue // to return the server wiht the least connections in a dynamic load balancer
}

func NewLoadBalancer() *LoadBalancer {
	lb := &LoadBalancer{
		Servers:     make([]string, 0),
		RRNext:      -1,
		Connections: make(PriorityQueue, 0),
	}
	heap.Init(&lb.Connections)
	return lb
}

func (lb *LoadBalancer) AddServer(connections int, serverAdd string) {
	lb.Servers = append(lb.Servers, serverAdd)
	item := NewQueueItem(connections, len(lb.Servers)-1)
	heap.Push(&lb.Connections, item)
}
