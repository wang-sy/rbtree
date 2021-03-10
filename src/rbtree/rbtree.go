package rbtree

// RBTree struct
// To create a rbtree ::
// rb := rbtree.RBTree{}
// rb.Set() \\ rb.Get()
type RBTree struct {
	root    *RBNode // tree root
	nilNode *RBNode // nil point
	count   uint    // tree node counter
}

// Set function
// if key is exist, update data
// if not exist, insert k-v pair to rbtree
func (tree *RBTree) Set(key Sortable, val interface{}) {
	if tree.count == 0 { // if tree is empty, set k-v as tree root
		tree.nilNode = &RBNode{
			color: BLACK,
		}
		tree.nilNode.p = tree.nilNode
		tree.nilNode.ls = tree.nilNode
		tree.nilNode.rs = tree.nilNode
		tree.root = &RBNode{
			p:     tree.nilNode,
			val:   val,
			key:   key,
			color: BLACK,
			ls:    tree.nilNode,
			rs:    tree.nilNode,
		}
		tree.count++
		return
	}
	insertPos := tree.getNode(key)
	if insertPos != tree.nilNode { // key is already exist
		insertPos.val = val
		return
	}
	insertPos = insertPos.p
	// not exist
	insertNode := &RBNode{
		p:     insertPos,
		val:   val,
		key:   key,
		color: RED,
		ls:    tree.nilNode,
		rs:    tree.nilNode,
	}
	if insertNode.key.Less(insertPos.key) {
		insertPos.ls = insertNode
	} else {
		insertPos.rs = insertNode
	}
	tree.count++
	tree.doBalance(insertNode)
}

// Get function
// Get the data corresponding to the given key
// if it's not exist, return nil
func (tree *RBTree) Get(key Sortable) interface{} {
	if tree.count == 0 {
		return nil
	}
	cur := tree.getNode(key)
	if cur == tree.nilNode {
		return nil
	}
	return cur.val
}

// search a node in RBTree, if that node is not exist, return tree.nilNode
func (tree *RBTree) getNode(key Sortable) *RBNode {
	cur := tree.root
	for cur != tree.nilNode {
		tree.nilNode.p = cur
		if key.Less(cur.key) {
			cur = cur.ls // goal is smaller than cur point, cur = cur.ls
		} else if cur.key.Less(key) {
			cur = cur.rs // goal is larger than cur point, cur = cur.rs
		} else { // goal is similar as cur point, return
			break
		}
	}
	return cur
}

// reverse sub tree's color, dep = 3
func reverseSubTreeColor(node *RBNode) {
	node.color = !node.color
	node.ls.color = !node.ls.color
	node.rs.color = !node.rs.color
	node.ls.ls.color = !node.ls.ls.color
	node.ls.rs.color = !node.ls.rs.color
	node.rs.ls.color = !node.rs.ls.color
	node.rs.rs.color = !node.rs.rs.color
}

// do the r&b balance
func (tree *RBTree) doBalance(node *RBNode) {
	if node.p.p != tree.nilNode && ((node.p.p.ls == node.p && node.p.p.rs != tree.nilNode) ||
		(node.p.p.rs == node.p && node.p.p.ls != tree.nilNode)) { // node's parent has brother
		reverseSubTreeColor(node)
		tree.doBalance(node.p.p)
		return
	}

	if node.p.p.ls == node.p {
		if node == node.p.ls {
			if node.p.p == tree.root {
				tree.root = node.p
			}
			node.p.color = BLACK
			node.p.p.color = RED
			node.p.p.rightRotate()
		} else {
			pp := node.p.p
			if pp == tree.root {
				tree.root = node
			}
			node.p.leftRotate()
			node.color = BLACK
			pp.color = RED
			pp.rightRotate()
		}
	} else { // node.p is right son of node.p.p
		if node == node.p.rs {
			if node.p.p == tree.root {
				tree.root = node.p
			}
			node.p.color = BLACK
			node.p.p.color = RED
			node.p.p.leftRotate()
		} else {
			pp := node.p.p
			if pp == tree.root {
				tree.root = node
			}
			node.p.rightRotate()
			node.color = BLACK
			pp.color = RED
			pp.leftRotate()
		}
	}

}

/*
// Delete RBNode which has key
func (tree *RBTree) Delete(key Sortable) {
	node := tree.getNode(key)
	if node == tree.nilNode {
		return
	}
	// get replaceNode && do replace
	replaceNode := tree.getDeleteNode(node)
	if replaceNode == tree.nilNode {
		if node == node.p.ls {
			node.p.ls = replaceNode
		} else {
			node.p.rs = replaceNode
		}
	}
	replaceNodeColor := replaceNode.color
	replaceNode.ls, replaceNode.rs, replaceNode.p, replaceNode.color = node.ls, node.rs, node.p, node.color
	replaceNode.ls.p, replaceNode.rs.p = replaceNode, replaceNode
	if replaceNodeColor != RED {
		replaceNode = tree.deleteBalance(replaceNode)
	}
	if node == tree.root {
		tree.root = replaceNode
	}
	tree.count--
}


// make the tree balance again
func (tree *RBTree)deleteBalance (node *RBNode) *RBNode {
	var pb *RBNode
	if node == node.p.ls {
		pb = node.p.rs
	} else {
		pb = node.p.ls
	}
	if pb.color == BLACK {
		if pb.ls.color == BLACK && pb.rs.color == BLACK {
			tree.deleteBalance(node.p)
		} else if node.rs.color
	}
}

func (tree *RBTree)getDeleteNode (node *RBNode) *RBNode {
	var resNode *RBNode
	if node.ls != tree.nilNode || node.rs != tree.nilNode {
		if node.ls == tree.nilNode { // have only right son, find min key in right subtree
			resNode = tree.min(node.rs)
		} else if node.ls == tree.nilNode {
			resNode = tree.max(node.ls)
		} else { // have two sons : swap(node, resNode), delete node
			resNode = tree.max(node.rs)
		}
	} else { // node have no son
		resNode = tree.nilNode
	}
	return resNode

}

*/

// find RBNode which have max key in a sub_tree which has node as root
func (tree *RBTree) max(node *RBNode) *RBNode {
	for node.rs != tree.nilNode {
		node = node.rs
	}
	return node
}

// find RBNode which have min key in a sub_tree which has node as root
func (tree *RBTree) min(node *RBNode) *RBNode {
	for node.ls != tree.nilNode {
		node = node.ls
	}
	return node
}
