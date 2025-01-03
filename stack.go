package main

type Stack struct {
	list []int
}

func (pstack *Stack) Push(element int) {
	pstack.list = append(pstack.list, element)
}

func (pstack *Stack) Pop() (element int, isempty bool) {
	if len(pstack.list) == 0 {
		return 0, false
	} else {
		passBack := pstack.list[len(pstack.list)-1]
		pstack.list = pstack.list[:len(pstack.list)-1]
		return passBack, true
	}
}
