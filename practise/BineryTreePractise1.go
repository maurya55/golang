package main

import (
	"fmt"
	"io"
	"os"
)

type Child struct {
	data  int
	right *Child
	left  *Child
}

type BineryTree struct {
	root *Child
}

func (t *BineryTree) insert(val int) *BineryTree {

	if t.root == nil {
		t.root = &Child{data: val, left: nil, right: nil}
	} else {
		t.root.insertData(val)
	}
	return t
}

func (ch *Child) insertData(val int) *Child {

	if val > ch.data {
		if ch.right == nil {
			ch.right = &Child{data: val, left: nil, right: nil}
		} else {
			ch.right.insertData(val)
		}
	} else {
		if ch.left == nil {
			ch.left = &Child{data: val, left: nil, right: nil}
		} else {
			ch.left.insertData(val)
		}
	}

	// fmt.Println(ch.left)

	return ch
}

func print(w io.Writer, node *Child, ns int, ch rune) {
	if node == nil {
		return
	}

	for i := 0; i < ns; i++ {
		fmt.Fprint(w, " ")
	}
	fmt.Fprintf(w, "%c:%v\n", ch, node.data)
	print(w, node.left, ns+2, 'L')
	print(w, node.right, ns+2, 'R')
}
func main() {

	tree := &BineryTree{}
	tree.insert(10).insert(100).insert(5).insert(95)

	fmt.Println(tree)
	print(os.Stdout, tree.root, 0, 'M')

}
