package rbtree

// The key inserted into the red black tree needs to implement sortable interface
// Realize the less function, which has the same purpose as the less than symbol
// if self < Sortable return true, else false
type Sortable interface {
	Less(Sortable) bool
}
