/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
package mergeTwoLists

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	last := new(ListNode)
	if list1 != nil || list2 != nil {
		mergeItem(list1, list2, last)
	} else {
		last = nil
	}
	return last
}

func mergeItem(list1 *ListNode, list2 *ListNode, last *ListNode) {
	if list1 != nil && list2 != nil {
		last.Next = new(ListNode)
		if list1.Val > list2.Val {
			last.Val = list2.Val
			mergeItem(list1, list2.Next, last.Next)
		} else {
			last.Val = list1.Val
			mergeItem(list1.Next, list2, last.Next)
		}
	} else if list1 != nil || list2 != nil {
		if list1 != nil {
			last.Val = list1.Val
			if list1.Next != nil {
				last.Next = new(ListNode)
				mergeItem(list1.Next, list2, last.Next)
			}
		} else if list2 != nil {
			last.Val = list2.Val
			if list2.Next != nil {
				last.Next = new(ListNode)
				mergeItem(list1, list2.Next, last.Next)
			}
		}
	}
}
