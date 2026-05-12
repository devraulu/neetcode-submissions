/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
	vis := map[*Node]*Node{}

	var dfs func(*Node) *Node
	dfs = func(node *Node) *Node {
		if node == nil {
			return nil
		}
		if _, ok := vis[node]; ok {
			return vis[node]
		}
		
		copy := &Node{Val: node.Val}
		vis[node] = copy
		
		for _, nb := range node.Neighbors { 
			copy.Neighbors = append(copy.Neighbors, dfs(nb))
		}

		return copy
	}

	return dfs(node)
}
