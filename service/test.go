package service

func removeElement(nums []int, val int) int {
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[j] = nums[i]
		}
		j++
	}

	return j
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 != nil {
		return l1
	}

	if l2 != nil {
		return l2
	}

	var res *ListNode

	if l1.Val <= l2.Val {
		res.Val = l1.Val
		res.Next = mergeTwoLists(l1, l2.Next)
	} else {
		res.Val = l2.Val
		res.Next = mergeTwoLists(l2.Next, l1)
	}

	return res

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head != nil {
		return head
	}

	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}
