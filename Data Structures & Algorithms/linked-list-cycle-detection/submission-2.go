/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
	visited := map[*ListNode]struct{}{}
	cursor := head
	for cursor != nil {
		if _, ok := visited[cursor]; ok {
			return true
		}
		visited[cursor] = struct{}{}
		cursor = cursor.Next
	}
	return false
}
