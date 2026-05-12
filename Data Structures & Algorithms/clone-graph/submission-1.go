/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	
	vis := map[*Node]*Node{}
	vis[node] = &Node{ Val:node.Val, Neighbors: make([]*Node, 0)}
	q := []*Node{}
	q = append(q, node)
		
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		for _, nb := range cur.Neighbors {
			if _, ok := vis[nb]; !ok {
				vis[nb] = &Node{Val: nb.Val, Neighbors: make([]*Node, 0)}
				q = append(q, nb) 
			}
			vis[cur].Neighbors = append(vis[cur].Neighbors, vis[nb])
		}
	}

	return vis[node]
}
