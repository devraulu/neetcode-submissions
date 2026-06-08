/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
	arr := []*ListNode{}

	curr := head
	for curr != nil {
		arr = append(arr, curr)
		curr = curr.Next
	}

	n := len(arr)
	i, j := 0, n - 1
	for i < j {
		arr[i].Next = arr[j]
		i++
		if i >= j {
			break
		}
		arr[j].Next = arr[i]
		j--
	}
	
	arr[i].Next = nil
}

func (ln *ListNode) String() string {
	if ln == nil {
		return "<nil>"
	}

	return fmt.Sprintf("{Val: %d, Next: %v}", ln.Val, ln.Next)

}