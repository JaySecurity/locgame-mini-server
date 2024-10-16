package network

import "sync"

// Queue holds values in a slice.
type Queue struct {
	values []func()
	mutex  sync.RWMutex
}

// Enqueue adds a value to the end of the queue
func (q *Queue) Enqueue(value func()) {
	q.mutex.Lock()
	q.values = append(q.values, value)
	q.mutex.Unlock()
}

// Dequeue removes first element of the queue and returns it, or nil if queue is empty.
// Second return parameter is true, unless the queue was empty and there was nothing to dequeue.
func (q *Queue) Dequeue() (value func(), ok bool) {
	if q.IsEmpty() {
		return nil, false
	}

	q.mutex.Lock()
	elem := q.values[0]
	q.values = q.values[1:]
	q.mutex.Unlock()

	return elem, true
}

// IsEmpty returns true if queue does not contain any elements.
func (q *Queue) IsEmpty() bool {
	q.mutex.RLock()
	defer q.mutex.RUnlock()
	return len(q.values) == 0
}
