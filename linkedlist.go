package main

type node struct {
	next  *node
	value int
}

type linkedList struct {
	head   *node
	length int
}

func (plist *linkedlist) PushFirst(data int) {
	newNode := node{value: data}
	if plist.head != nil {
		newNode.next = plist.head
		plist.head = &newNode
		plist.length++
	} else {
		plist.head = &newNode
		plist.length++
	}
}
