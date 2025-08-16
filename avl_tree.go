package go_algorithm

type AVLTreeNode struct {
	key   int
	data  any
	left  *AVLTreeNode
	right *AVLTreeNode
}

func GetHeight(root *AVLTreeNode) int {
	if root == nil {
		return 0
	}
	return max(GetHeight(root.left), GetHeight(root.right)) + 1
}

func GetBalance(root *AVLTreeNode) int {
	if root == nil {
		return 0
	}
	return GetHeight(root.left) - GetHeight(root.right)
}

func GetMinNode(root *AVLTreeNode) *AVLTreeNode {
	if root.left == nil {
		return root
	}
	return GetMinNode(root.left)
}

func LeftRotate(root *AVLTreeNode) *AVLTreeNode {
	head := root.right
	tmp := head.left
	head.left = root
	root.left = tmp

	return head
}

func RightRotate(root *AVLTreeNode) *AVLTreeNode {
	head := root.left
	tmp := head.right
	head.right = root
	root.right = tmp

	return head
}

func ReBalance(root *AVLTreeNode) *AVLTreeNode {
	b := GetBalance(root)
	if b > 1 { // 左子树高 右旋
		if GetBalance(root.left) > 0 {
			root.left = RightRotate(root.left)
		}
		return LeftRotate(root)
	}
	if b < -1 {
		if GetBalance(root.right) < 0 {
			root.right = LeftRotate(root.right)
		}
		return RightRotate(root)
	}

	return root
}

func Add(root *AVLTreeNode, key int) *AVLTreeNode {
	if root == nil {
		return &AVLTreeNode{key: key, left: nil, right: nil}
	}
	if root.key > key {
		root.left = Add(root.left, key)
	} else if root.key < key {
		root.right = Add(root.right, root.key)
	}

	return ReBalance(root)
}

func Delete(root *AVLTreeNode, key int) *AVLTreeNode {
	if root == nil {
		return nil
	}
	if root.key > key {
		root.left = Delete(root.left, key)
	} else if root.key < key {
		root.right = Delete(root.right, root.key)
	} else {
		if root.left == nil || root.right == nil {
			temp := root.right
			if root.right == nil {
				temp = root.left
			}
			root = nil
			return temp
		}
		temp := GetMinNode(root.right)
		root.key = temp.key
		root.right = Delete(root.right, temp.key)
	}

	return ReBalance(root)
}

func contain(root *AVLTreeNode, key int) *AVLTreeNode {
	if root == nil {
		return nil
	}

	if root.key == key {
		return root
	} else if root.key > key {
		return contain(root.left, key)
	} else {
		return contain(root.right, key)
	}
}
