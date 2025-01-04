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
