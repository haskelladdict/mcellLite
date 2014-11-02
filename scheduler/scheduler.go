// Package scheduler is a local event scheduler responsible for managing
// simulation events (molecule diffusion, unimolecular decay, ...)
package scheduler

import "container/heap"

// Events provides access to simulation events managed by the scheduler
type Events struct {
	events *priorityQueue
}

// Create creates a new event scheduler
func Create() Events {
	ev := Events{}
	ev.events = new(priorityQueue)
	heap.Init(ev.events)
	return ev
}

// Add adds a new event to the scheduler
func (e *Events) Add(ev Event) {
	i := item{ev, ev.Time(), -1}
	heap.Push(e.events, &i)
}

// Pop removes an event from the scheduler
func (e *Events) Pop() Event {
	i := heap.Pop(e.events).(*item)
	return i.event
}

// item is a single Event entry in the event scheduler
type item struct {
	event Event   // type of event
	time  float64 // time of event
	index int     // index of event in heap; needed for updating
}

// priorityQueue implements a priority queue via go's heap
type priorityQueue []*item

func (pq priorityQueue) Len() int {
	return len(pq)
}

func (pq priorityQueue) Less(i, j int) bool {
	return pq[i].time > pq[j].time
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *priorityQueue) Push(x interface{}) {
	n := len(*pq)
	i := x.(*item)
	i.index = n
	*pq = append(*pq, i)
}

func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	i := old[n-1]
	i.index = -1
	*pq = old[:n-1]
	return i
}

/*
















*/
