package main

import "fmt"

type Tree struct {
	data  int
	left  *Tree
	right *Tree
}

func new(data int) *Tree {
	return &Tree{
		data:  data,
		left:  nil,
		right: nil,
	}
}

func main() {

	// tree structure
	//
	//        10 -> root
	//       /  \
	//      9    8 -> leave
	//     / \
	//    3   5
	//
	//

	tree := new(10)
	fmt.Println(tree)

	tree.left = new(9)
	tree.right = new(8)
	tree.left.left = new(3)
	tree.left.right = new(5)

	copyTree := tree

	for copyTree != nil {
		fmt.Println(copyTree.data)
	}

	fmt.Println(tree)

}
