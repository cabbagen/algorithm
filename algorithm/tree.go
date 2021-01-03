package algorithm

type Tree struct {
	RootNodeId   int
	NodeBoxes     []TreeBox
}

type TreeBox struct {
	Data         TreeNodeData
	FirstNode    *TreeNode
	ParentNode   *TreeBox
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
		RootNodeId: rootId,
		NodeBoxes: []TreeBox {},
	}
}

func (t *Tree) CreateTree(data []TreeNodeData) {
	var boxes map[int]*TreeBox = make(map[int]*TreeBox)
	var pointer map[int][]*TreeNode = make(map[int][]*TreeNode)

	for i, length := 0, len(data); i < length; i++ {
		node := TreeNode { data[i], nil }

		boxes[data[i].Id] = &TreeBox {data[i], nil, nil }

		if pointer[data[i].ParentId] == nil {
			pointer[data[i].ParentId] = []*TreeNode { &node, &node }
			continue
		}
		pointer[data[i].ParentId][1].Next, pointer[data[i].ParentId] = &node, []*TreeNode { pointer[data[i].ParentId][0], &node }
	}

	for id, treeBox := range boxes {
		if len(pointer[id]) > 0 {
			treeBox.FirstNode = pointer[id][0]
		}
		treeBox.ParentNode = boxes[treeBox.Data.ParentId]
		t.NodeBoxes = append(t.NodeBoxes, *treeBox)
	}
}

func (t *Tree) ClearTree() {
	t.NodeBoxes = []TreeBox {}
	t.RootNodeId = 0
}

func (t *Tree) IsEmptyTree() bool {
	return len(t.NodeBoxes) == 0
}
