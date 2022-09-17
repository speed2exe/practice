// https://leetcode.com/problems/course-schedule-ii/

package course_schedule_ii

func findOrder(numCourses int, prerequisites [][]int) []int {
	res := make([]int, 0, numCourses)

	// build graph
	// first dereference gets the numCourses
	// second dereference gets the list of dependencies
	graph := make([][]int, numCourses)
	for _, p := range prerequisites {
		key, val := p[0], p[1]
		graphVal := graph[key]
		graph[key] = append(graphVal, val)
	}

	statusByNode := make([]Status, numCourses)

	for i := 0; i < numCourses; i++ {
		if !fillOrder(graph, statusByNode, i, &res) {
			return []int{}
		}
	}

	return res
}

func fillOrder(graph [][]int, statusByNode []Status, node int, res *[]int) bool {
	status := statusByNode[node]
	switch status {
	case VISITING:
		return false
	case VISITED:
		return true
	case UNVISITED:
	default:
		panic("unreachable")
	}
	statusByNode[node] = VISITING

	for _, n := range graph[node] {
		if !fillOrder(graph, statusByNode, n, res) {
			return false
		}
	}

	*res = append(*res, node)
	statusByNode[node] = VISITED
	return true
}

type Status int

const (
	UNVISITED Status = iota
	VISITING
	VISITED
)
