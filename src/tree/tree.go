package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

func CreateNode(value int) *Node {
	return &Node{Value: value}
}

func (node Node) Print() {
	fmt.Println(node.Value)
}
func (node *Node) SetValue(value int) {
	if node == nil {
		println("Setting value to nil node. ignore")
		return
	}
	node.Value = value
}
func (node *Node) Traverse() {
	node.TraverseFunc(func(n *Node) {
		n.Print()
	})
	fmt.Println()
}
func (node *Node) TraverseFunc(f func(*Node)) {
	if node == nil {
		return
	}
	node.Left.TraverseFunc(f)
	f(node)
	node.Right.TraverseFunc(f)
}
func (node *Node) TraverseWithChannel() <-chan *Node {
	c := make(chan *Node)
	go func() {
		node.TraverseFunc(func(n *Node) {
			c <- n
		})
		close(c)
	}()
	return c
}

// func main() {
// 	var root Node
// 	root = Node{value: 3}
// 	root.left = &Node{}
// 	root.right = &Node{5, nil, nil}
// 	root.right.left = new(Node)
// 	root.left.right = CreateNode(2)
// 	root.right.left.SetValue(4)
// 	// root.SetValue(100)

// 	// pRoot := &root
// 	// pRoot.Print()
// 	// pRoot.SetValue(200)
// 	// pRoot.Print()
// 	// root.right.right.SetValue(111)
// 	// nodes := []Node{
// 	// 	{value: 3},
// 	// 	{},
// 	// 	{6, nil, nil},
// 	// }
// 	// println(&nodes[0])
// 	// var pRoot *Node
// 	// pRoot = &root
// 	// pRoot.Print()
// 	// pRoot.SetValue(100)
// 	// pRoot.Print()
// 	root.traverse()
// }
