package rbtree

type RBNode struct {
	ls    *RBNode     // left son
	rs    *RBNode     // right son
	p     *RBNode     // father
	color bool        // color, true for red, false for black
	key   Sortable    // a Sortable type val
	val   interface{} // data, a interface{} val
}

const (
	RED   = true
	BLACK = false
)

// Do left rotate operation for the subtree, which is based on node
//        node                        rs
//       /     \                    /    \
//	    ls     rs       --->      node    rr
// 	    /\     / \               /    \
//     ll lr  rl rr            ls      rl
//                            /  \
//                           ll  lr
func (node *RBNode) leftRotate() {
	rs := node.rs
	rs.p, node.p = node.p, rs
	rs.ls, node.rs = node, rs.ls
}

// Do right rotate operation for the subtree, which is based on node
//        node                        ls
//       /     \                    /    \
//	    ls     rs       --->      ll    node
// 	    /\     / \                     /    \
//     ll lr  rl rr                  lr      rs
//                                          /  \
//                                         rl  rr
func (node *RBNode) rightRotate() {
	ls := node.ls
	ls.p, node.p = node.p, ls
	ls.rs, node.ls = node, ls.rs
}
