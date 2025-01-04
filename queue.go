package main

type Queue struct {
	list []int
}

func (pqueue *Queue) Enqueue(element int) {
	pqueue.list = append(pqueue.list, element)
}

func (pqueue *Queue) Dequeue() (int, bool) {
	if len(pqueue.list) == 0 {
		return 0, false
	}

	passBack := pqueue.list[0]
	pqueue.list = pqueue.list[1:]
	return passBack, true

}

func (pqueue *Queue) Peek() (int, bool) {
	if len(pqueue.list) == 0 {
		return 0, false
	}
	return pqueue.list[0], true
}

func (pqueue *Queue) IsEmpty() bool {
	if len(pqueue.list) == 0 {
		return true
	}
	return false

}

func (pqueue *Queue) Size() int {
	return len(pqueue.list)
}

func (pqueue *Queue) Clear() {
	pqueue.list = pqueue.list[:0]
}

func (pqueue *Queue) Contains(element int) (contains bool) {
	contains = false
	for i := 0; i < len(pqueue.list); i++ {
		if pqueue.list[i] == element {
			contains = true
			break
		}
	}
	return contains
}
