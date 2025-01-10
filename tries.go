package main

const ALPHABET = 26 //to limit trie size

type node struct {
	history [rune]*node //rune <-> int 32
	head    bool
}

type trie struct {
	root *node
}

func Create() *trie {
	return &trie{root: &node{history: make(map[rune]*node)}}
}

//insert

//search

//delete

//listWords

//countWords
