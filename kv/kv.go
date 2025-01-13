package kv

import (
  "os"
  "fmt"
  "math"
  "bytes"
)

// set default order of tree.
const defaultOrder = 500

// pager constants
const maxPageSize = 4096 //change to math.MaxUint16 later on.
const minPageSize = 32

const maxOrder = math.MaxUint16

type BPTree struct {
  order int

  storage *storage

  metadata *treeMetaData

  // minKey = ceil(order/2) - 1
  minKeyNum int
}

type treeMetaData struct {
  order uint16
  rootID uint32
  pageSize uint16
}

// Open either opens a new tree or loads a pre existing tree.
func Open(path string) (*BPTree, error) {
  // replace defaultOrder with user selected order
  // replace pageSize with user selected pageSize

  // use to set page size
  pageSize := os.Getpagesize()

  storage, err := newStorage(path, pageSize)
  if err != nil {
    return nil, fmt.Errorf("failed to init the storage: %w", err)
  }

  metadata, err := storage.loadMetadata()
  if err != nil {
    return nil, fmt.Errorf("failed to init the metadata: %w", err)
  } 

  if metadata != nil && metadata.order != defaultOrder {
    return nil, fmt.Errorf("Tried to open a tree with order %w, but has order %w", metadata.order, defaultOrder)
  }

  minKeyNum := calcMinOrder(defaultOrder)

  return &BPTree{order : defaultOrder, storage : storage, metadata : metadata, minKeyNum : minKeyNum}, nil
}

type node struct {
  id uint32

  isLeaf bool
  parentId uint32

  key [][]byte
  // keyNum represents number of keys present in the node.
  keyNum int
  
  // pointer can either be a value or children
  // based on if the node is a leaf or not
  pointers []*pointer
}

type pointer struct {
  value interface{}
}

func (p *pointer) isValue() bool {
  _, ok := p.value.([]byte)
  return ok
}

func (p *pointer) isNodeId() bool {
  _, ok := p.value.(uint32)
  return ok
}

func (p *pointer) asValue() []byte {
  return p.value.([]byte)
}

func (p *pointer) asNodeId() uint32 {
  return p.value.(uint32)
}

// returns (value, err)
func (t *BPTree) Get(key []byte) ([]byte, error) {
  if t.metadata == nil {
    return nil, fmt.Errorf("Not initialized")
  }

  leaf, err := findLeaf(key)
  if err != nil {
    return nil, fmt.Errorf("Could not find leaf : %w", err)
  }

  for i := 0; i < leaf.keyNum ; i++ {
    if compare(key, leaf.key[i]) == 0 {
      return leaf.pointers[i].asValue(), nil
    }
  }

  return nil, fmt.Errorf("Key not found error")
}

func compare(byteA ,byteB []byte) int {
  return bytes.Compare(byteA, byteB)
}

func calcMinOrder(order uint16) {}
