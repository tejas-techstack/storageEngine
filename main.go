package main

import (
  "log"
)

func main() {
  head := createNewTree();
  insertArray := [5]int{12, 9, 1, 5, 6};

  var err error

  for _, v := range insertArray{
    insertKey(&head, v);
  }

  err = printTree(head)
  if err != nil{
    log.Println(err)
  }
}

