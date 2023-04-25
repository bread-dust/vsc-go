/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertionSortList(head *ListNode) *ListNode {
if head == nil && head.Next == nil {
		return nil
	}

	dummyNode := new(ListNode)
	dummyNode.Next = head

	need := head.Next
	pre := head
	tempHead := dummyNode
=
	for need!=nil{
		if pre.Next!=nil&&pre.Val<=need.Val{
			need=need.Next
			pre=pre.Next
			continue
		}

        tempHead = dummyNode
		for tempHead.Next.Val <= need.Val{
			tempHead=tempHead.Next
		}
 
		pre.Next =need.Next
		need.Next=tempHead.Next
		tempHead.Next = need
		need=pre.Next

	}
	return dummyNode.Next
}