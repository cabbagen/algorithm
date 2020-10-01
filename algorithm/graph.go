package algorithm

type MyGraphNode struct {
	ID         int
	Title      string
}

type MyGraphEdge struct {
	Start      int
	End        int
	Weight     int
}

type MyGraph struct {
	Graph       [][]int
	Vertex      []MyGraphNode
}

func NewMyGraph(vertex []MyGraphNode, edges []MyGraphEdge) *MyGraph {
	var graph [][]int

	for i := 0; i < len(vertex); i++ {
		graph = append(graph, make([]int, len(vertex)))

		for j := 0; j < len(vertex); j++ {
			graph[i][j] = -1

			for k := 0; k < len(edges); k++ {
				if edges[k].Start == vertex[i].ID && edges[k].End == vertex[j].ID {
					graph[i][j] = edges[k].Weight
					break
				}
			}
		}
	}
	return &MyGraph { graph, vertex }
}

func (mg *MyGraph) DestroyGraph() {
	mg.Graph = [][]int{}
	mg.Vertex = []MyGraphNode{}
}

func (mg *MyGraph) GetVex(id int) (MyGraphNode, bool) {
	for _, value := range mg.Vertex {
		if value.ID == id {
			return value, true
		}
	}
	return MyGraphNode{}, false
}

func (mg *MyGraph) PutVex(id int, title string) {
	for _, value := range mg.Vertex {
		if value.ID == id {
			value.Title = title
		}
	}
}

func (mg MyGraph) FirstAdjVex(id int) []MyGraphNode {
	var vertexIndex int
	var nodes []MyGraphNode

	for i := 0; i < len(mg.Vertex); i++ {
		if mg.Vertex[i].ID == id {
			vertexIndex = i
			break
		}
	}

	for j := 0; j < len(mg.Graph); j++ {
		for k := 0; k < len(mg.Graph[j]); k++ {
			if j == vertexIndex && mg.Graph[j][k] > -1 {
				nodes = append(nodes, mg.Vertex[k])
				continue
			}
			if k == vertexIndex && mg.Graph[j][k] > -1 {
				nodes = append(nodes, mg.Vertex[j])
				continue
			}
		}
	}
	return nodes
}

func (mg *MyGraph) InsertVex(graphNode MyGraphNode) {
	var needAppendRow []int = make([]int, len(mg.Vertex))

	for i := 0; i < len(needAppendRow); i++ {
		needAppendRow[i] = -1
	}

	mg.Graph = append(mg.Graph, needAppendRow)

	for i := 0; i < len(mg.Graph); i++ {
		mg.Graph[i] = append(mg.Graph[i], -1)
	}

	mg.Vertex = append(mg.Vertex, graphNode)
}

func (mg *MyGraph) DeleteVex(id int) {
	var deletedIndex int

	for i := 0; i < len(mg.Vertex); i++ {
		if id == mg.Vertex[i].ID {
			deletedIndex = i
			mg.Vertex = append(mg.Vertex[0:i], mg.Vertex[i + 1:]...)
			break
		}
	}

	mg.Graph = append(mg.Graph[0:deletedIndex], mg.Graph[deletedIndex + 1:]...)

	for i := 0; i < len(mg.Graph); i++ {
		mg.Graph[i] = append(mg.Graph[i][0:deletedIndex], mg.Graph[i][deletedIndex + 1:]...)
	}
}

func (mg *MyGraph) InsertArc(startId, endId, weight int) {
	var startIndex, endIndex int

	for i := 0; i < len(mg.Vertex); i++ {
		if startId == mg.Vertex[i].ID {
			startIndex = i
			continue
		}
		if endId == mg.Vertex[i].ID {
			endIndex = i
			continue
		}
	}

	mg.Graph[startIndex][endIndex] = weight
}

func (mg *MyGraph) DeleteArc(startId, endId int) {
	var startIndex, endIndex int

	for i := 0; i < len(mg.Vertex); i++ {
		if startId == mg.Vertex[i].ID {
			startIndex = i
			continue
		}
		if endId == mg.Vertex[i].ID {
			endIndex = i
			continue
		}
	}

	mg.Graph[startIndex][endIndex] = -1
}

func (mg *MyGraph) DFSTraverse() {
	var visited []int = make([]int, len(mg.Vertex))

	for i := 0; i < len(mg.Vertex); i++ {
		if visited[i] == 1 {
			continue
		}
		mg.DFS(&visited, i)
	}
}

func (mg *MyGraph) DFS(visited *[]int, index int) {
	(*visited)[index] = 1

	for j := 0; j < len(mg.Vertex); j++ {
		if mg.Graph[index][j] > -1 && (*visited)[j] == 0 {
			mg.DFS(visited, j)
		}
	}
}

func (mg *MyGraph) HFSTraverse() {
	var queue []int
	var visited []int = make([]int, len(mg.Vertex))

	for i := 0; i < len(mg.Vertex); i++ {
		if visited[i] == 1 {
			continue
		}
		visited[i] = 1
		queue = append(queue, i)

		for len(queue) > 0 {
			i, queue = queue[0:1][0], queue[1:]

			for j := 0; j < len(mg.Vertex); j++ {
				if mg.Graph[i][j] > -1 && visited[j] == 0 {
					visited[j], queue = 1, append(queue, j)
				}
			}
		}
	}
}
