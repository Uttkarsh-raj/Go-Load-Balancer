package models

import (
	"container/heap"
)

type QueueItem struct {
	Connections int
	ServerIndex int
}

func NewQueueItem(connections, index int) *QueueItem {
	return &QueueItem{
		Connections: connections,
		ServerIndex: index,
	}
}

type PriorityQueue []*QueueItem

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*QueueItem)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Len() int {
	return len(*pq)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Connections < pq[j].Connections
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// Update modifies the number of connections of a specific QueueItem.
func (pq *PriorityQueue) Update(item *QueueItem, connections int) {
	item.Connections = connections
	heap.Fix(pq, item.ServerIndex)
}

// Will implement a Min Heap so must return the item with the least connections
func (pq *PriorityQueue) GetServerWithMinConnections() *QueueItem {
	if pq.Len() == 0 {
		return nil
	}
	item := heap.Pop(pq).(*QueueItem)
	return item
}
