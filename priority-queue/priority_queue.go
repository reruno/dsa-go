package priority_queue

import "container/heap"

func NewPriorityQueue() *PrioirtyQueue {
	pq := make(pq, 0)
	heap.Init(&pq)
	return &PrioirtyQueue{
		inner: pq,
	}
}

type PrioirtyQueue struct {
	inner pq
}

func (p *PrioirtyQueue) Len() int {
	return len(p.inner)
}

func (p *PrioirtyQueue) Push(value any, priority int) {
	heap.Push(&p.inner, &Item{
		Value:    value,
		Priority: priority,
	})
}

func (p *PrioirtyQueue) Pop() *Item {
	return heap.Pop(&p.inner).(*Item)
}
