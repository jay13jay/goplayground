package btree

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Value int
}

func InsertNodeToTree(tree *TreeNode, node *TreeNode) {
	if tree == nil {
		tree = node
	}
	if node.Value > tree.Value {
		InsertNodeToTree(tree.Right, node)
	}
	if node.Value < tree.Value {
		InsertNodeToTree(tree.Left, node)
	}
}

func InitTree(values ...int) (root *TreeNode) {
	rootNode := TreeNode{Value: values[0]}
	for _, value := range values {
		node := TreeNode{Value: value}
		InsertNodeToTree(&rootNode, &node)
	}
	return &rootNode
}
