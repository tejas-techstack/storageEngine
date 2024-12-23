package main

import (
  "log"
)

func insertKey(head **node, key int){
	if *head == nil {
    log.Fatal("Invalid head")
		return
	}

  if (*head).key == -1{
    *head = createNode(key);
  }

	current := *head
	for {
		if key == current.key {
      log.Println("key already present : ", key)
			return
		}
		if key < current.key {
			if current.left == nil {
				current.left = createNode(key)
				return
			}
			current = current.left
		} else {
			if current.right == nil {
				current.right = createNode(key)
				return
			}
			current = current.right
		}
	}
}


