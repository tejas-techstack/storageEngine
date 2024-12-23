package main

type arrayIndex int

type node struct {
	key   int
	index arrayIndex
	left  *node
	right *node
}


