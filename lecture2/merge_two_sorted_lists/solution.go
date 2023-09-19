package merge_two_sorted_lists

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var ans *ListNode
	var temp *ListNode
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	}
	if list1.Val > list2.Val {
		temp = list2
		ans = temp
		list2 = list2.Next
	} else {
		temp = list1
		ans = temp
		list1 = list1.Next
	}
	for list1 != nil && list2 != nil {
		fmt.Println(*temp)
		if list1.Val > list2.Val {
			temp.Next = list2
			temp = list2
			list2 = list2.Next
		} else {
			temp.Next = list1
			temp = list1
			list1 = list1.Next
		}
	}
	if list1 == nil {
		temp.Next = list2
	} else {
		temp.Next = list1
	}
	return ans
}
