package main

type Stack struct {
	list []int
}

func (pstack *Stack) Push(element int) {
	pstack.list = append(pstack.list, element)
}

func (pstack *Stack) Pop() (int, bool) {
	if len(pstack.list) == 0 {
		return 0, false
	}
	passBack := pstack.list[len(pstack.list)-1]
	pstack.list = pstack.list[:len(pstack.list)-1]
	return passBack, true
}

func (pstack *Stack) Peek() int {
	return pstack.list[len(pstack.list)-1]
}

func (pstack *Stack) IsEmpty() bool {
	return len(pstack.list) == 0
}

func (pstack *Stack) Size() int {
	return len(pstack.list)
}

func (pstack *Stack) Clear() {
	pstack.list = pstack.list[:0]
}

func (pstack *Stack) Contains(element int) (contains bool) {
	contains = false
	for i := 0; i < len(pstack.list); i++ {
		if pstack.list[i] == element {
			contains = true
			break
		}
	}
	return contains
}
