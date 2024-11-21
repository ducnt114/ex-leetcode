package main

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	l1 := mergeTwoLists(
		&ListNode{1, &ListNode{2, &ListNode{4, nil}}},
		&ListNode{1, &ListNode{3, &ListNode{4, nil}}},
	)
	l1.print()

	fmt.Println("***********")

	l2 := mergeTwoLists(nil, nil)
	l2.print()

	fmt.Println("***********")

	l3 := mergeTwoLists(nil, &ListNode{0, nil})
	l3.print()

}

//func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
//	if list1 == nil {
//		return list2
//	}
//	if list2 == nil {
//		return list1
//	}
//	if list1.Val < list2.Val {
//		list1.Next = mergeTwoLists(list1.Next, list2)
//		return list1
//	}
//	list2.Next = mergeTwoLists(list1, list2.Next)
//	return list2
//}

func (l *ListNode) print() {
	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	var first *ListNode
	var last *ListNode
	var current *ListNode
	for {
		if list1 == nil && list2 == nil {
			break
		}
		if list1 == nil {
			current = list2
			list2 = list2.Next
		} else if list2 == nil {
			current = list1
			list1 = list1.Next
		} else {
			if list1.Val < list2.Val {
				current = list1
				list1 = list1.Next
			} else {
				current = list2
				list2 = list2.Next
			}
		}
		if first == nil {
			first = current
			last = first
		} else {
			last.Next = current
			last = current
		}
	}
	return first
}
