package main

type Queue struct {
	list []int
}

func (pqueue *Queue) Enqueue(element int) {
	pqueue.list = append(pqueue.list, element)
}

func (pqueue *Queue) Dequeue() (element int, isempty bool) {
	if len(pqueue.list) == 0 {
		return 0, false
	} else {
		passBack := pqueue.list[0]
		pqueue.list = pqueue.list[1:]
		return passBack, true
	}
}

func (pqueue *Queue) Peek() (element int, isempty bool) {
	if len(pqueue.list) == 0 {
		return 0, false
	} else {
		passBack := pqueue.list[0]
		return passBack, true
	}
}

func (pqueue *Queue) IsEmpty() (isempty bool) {
	if len(pqueue.list) == 0 {
		return true
	} else {
		return false
	}
}
