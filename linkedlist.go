package main

import "fmt"

type node struct {
	next  *node
	value int
}

type linkedList struct {
	head   *node
	length int
}

// basically checks if the head of the list is nil and if so the new node will become a head
// else the new node will become the node in front of head.
func (plist *linkedList) Push(data int) {
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

func (plist *linkedList) PrintList() {
	currentNode := plist.head
	for currentNode != nil {
		currentNode := currentNode.next
		fmt.Printf("%d", currentNode.value)
	}
	return
}

func (plist *linkedList) Length() int {
	return plist.length
}

func (plist *linkedList) Delete(data int) {
	if plist.head == nil {
		return
	}

	if plist.head.value == data {
		plist.length--
		plist.head = plist.head.next
	}

	current := plist.head
	next := plist.head.next

	for next != nil {
		if next.value == data {
			current.next = next.next
			plist.length--
			return
		}
		current = next
		next = next.next
	}
}
