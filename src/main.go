package main

import (
	"fmt"
	"rbtree/src/rbtree"
)

// Int is a Struct which have Sortable interface
type Int struct {
	val int
	rbtree.Sortable
}

// Less is the implemention of Sortable interface
func (a Int) Less(than rbtree.Sortable) bool {
	return a.val < than.(Int).val
}

func main() {
	tree := rbtree.RBTree{}
	A := Int{val: 10}
	B := Int{val: 20}
	C := Int{val: 30}
	tree.Set(A, 10)
	tree.Set(B, 20)
	tree.Set(C, 30)
	fmt.Println(tree.Get(A))
	fmt.Println(tree.Get(B))
	fmt.Println(tree.Get(C))
}
