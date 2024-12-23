package main

import (
  "fmt"
  "errors"
)

func createNewTree() *node{
  return &node{
    key: -1,
    index : -1,
    left : nil,
    right: nil,
  }
}

func createNode(key int) *node {
	return &node{
		key:   key,
		index: -1,
		left:  nil,
		right: nil,
	}
}

func inorder(head *node){
	// Inorder print
	if head != nil {
		inorder(head.left)
		fmt.Printf("%d ", head.key)
		inorder(head.right)
	}
}

func printTree(head *node) (error){
  if head == nil {
    return errors.New("Empty tree")
  }
  inorder(head);
  fmt.Printf("\n");
  return nil;
}

