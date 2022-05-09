package binaryTree

import (
	"github.com/isdamir/gotype"
	"testing"
)

func TestBinaryTree(t *testing.T) {

	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	root := ArrayToTree(data, 0, len(data)-1)

	gotype.PrintTreeMidOrder(root)
	PrintTreeLayer(root)
}
