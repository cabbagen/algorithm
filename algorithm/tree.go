// 孩子双亲表示法

package algorithm

import "fmt"

type Tree struct {
	NodeBoxs     []TreeBox
	RootNodeId   int
}

type TreeBox struct {
	Data         TreeNodeData
	FirstNode    *TreeNode
}

type TreeNodeData struct {
	Id           int
	ParentId     int
}

type TreeNode struct {
	Data         TreeNodeData
	Next         *TreeNode
}

func NewTree(rootId int) Tree {
	return Tree {
		NodeBoxs: []TreeBox {},
		RootNodeId: rootId,
	}
}

func (t *Tree) CreateTree(datas []TreeNodeData) {	
	var group map[int][]TreeNode = make(map[int][]TreeNode)
	
	for i := 0; i < len(datas); i++ {
		node := TreeNode{ datas[i], nil }
		
		group[datas[i].ParentId] = append(group[datas[i].ParentId], node)
       
		if (len(group[datas[i].ParentId]) > 1) {
			group[datas[i].ParentId][len(group[datas[i].ParentId]) - 2].Next = &node
		}
	}
	
	fmt.Printf("group: %v\n", group)
	
	for j := 0; j < len(datas); j++ {
		var firstNode *TreeNode
		
		if (len(group[datas[j].Id]) > 0) {
			firstNode = &group[datas[j].Id][0]
		}
		
		t.NodeBoxs = append(t.NodeBoxs, TreeBox { datas[j], firstNode })
	}
	
	current := t.NodeBoxs[1].FirstNode
	
	fmt.Printf("current1: %v\n", *current)
	
	fmt.Printf("current2: %v\n", *((*current).Next))
}

func (t *Tree) ClearTree() {
	t.NodeBoxs = []TreeBox{}
	t.RootNodeId = 0
}

func (t *Tree) IsEmptyTree() bool {
	return len(t.NodeBoxs) == 0
}

func (t *Tree) GetTreeNodePaths() {
	// var paths []int
	
	// current := t.NodeBoxs[1].FirstNode
	
	// fmt.Printf("current1: %v\n", *current)
	
	// fmt.Printf("current2: %v\n", *((*current).Next))
	
	// fmt.Printf("current: %v\n", )
	
	
	
	// for current != nil {
		// paths = append(paths, (*current).Data.Id)
		// current = current.Next
	// }
	
	// fmt.Printf("paths: %v\n", paths)
}

func (t *Tree) GetTreeDepth() int {
	var maxDepth int = 0
	
	// for i := 0; i < len(t.NodeBoxs); i++ {
		// if t.NodeBoxs[i].FirstNode == nil {
			// continue
		// }
		// depth, current := 0, t.NodeBoxs[i].FirstNode
		
		// for current != nil {
			// fmt.Printf("id = %d\n", (*current).Data.Id)
			// depth = depth + 1
			// current = current.Next
		// }
		
		// fmt.Printf("current depth: %d\n", depth)
		
		// if maxDepth < depth {
			// maxDepth = depth
		// }
	// }
	
	return maxDepth
}
