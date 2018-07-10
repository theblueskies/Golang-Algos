package bst

// AllPaths keeps a list of all the paths
type AllPaths struct {
	paths [][]int
}

// SearchPaths builds a tree with list of values, then searches for all paths which lead to a sum of "num"
func SearchPaths(v []int, num int) [][]int {
	bst := BinarySearchTree{}
	a := AllPaths{}
	p := []int{}

	bst.BuildTree(v)

	if bst.root == nil {
		return a.paths
	}

	searchAllPathsHelper(bst.root, &a, p, num)

	return a.paths
}

func sumArr(v []int) int {
	s := 0
	for _, n := range v {
		s += n
	}
	return s
}

func searchAllPathsHelper(node *Node, a *AllPaths, path []int, num int) {
	if node != nil {
		searchAllPaths(node, a, path, num)
		searchAllPathsHelper(node.left, a, path, num)
		searchAllPathsHelper(node.right, a, path, num)
	}

}

func searchAllPaths(node *Node, a *AllPaths, path []int, num int) {
	if node != nil {

		// fmt.Println(node.value)
		path = append(path, node.key)
		sum := sumArr(path)
		if sum == num {
			a.paths = append(a.paths, path)
		} else if sum > num {
			path = path[:len(path)-1]
			return
		}
		searchAllPaths(node.left, a, path, num)
		searchAllPaths(node.right, a, path, num)
	}
}
