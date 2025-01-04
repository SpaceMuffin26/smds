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
