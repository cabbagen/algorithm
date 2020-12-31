package main

import (
	// "fmt"
	"algorithm/algorithm"
)

func main() {
	var treeNodeDatas []algorithm.TreeNodeData = []algorithm.TreeNodeData {
		algorithm.TreeNodeData{0, -1},
		algorithm.TreeNodeData{1, 0},
		algorithm.TreeNodeData{2, 0},
		algorithm.TreeNodeData{3, 1},
		algorithm.TreeNodeData{4, 1},
		algorithm.TreeNodeData{5, 1},
		algorithm.TreeNodeData{6, 2},
		algorithm.TreeNodeData{7, 2},
		algorithm.TreeNodeData{8, 2},
		algorithm.TreeNodeData{9, 7},
		algorithm.TreeNodeData{10, 7},
		algorithm.TreeNodeData{11, 10},
		algorithm.TreeNodeData{12, 11},
	}
	
	tree := algorithm.NewTree(0)
	
	tree.CreateTree(treeNodeDatas)
	
	// depth := tree.GetTreeDepth()
	
	tree.GetTreeNodePaths()
	
	// fmt.Printf("result: %d\n", depth)
}
