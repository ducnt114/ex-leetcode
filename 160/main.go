/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    pA := headA
    pB := headB
    
    lenA := 1
    for pA.Next != nil {
        pA = pA.Next
        lenA++
    }
    lenB := 1
    for pB.Next != nil {
        pB = pB.Next
        lenB++
    }
    if pA.Val != pB.Val {
        return nil
    }
    pA = headA
    pB = headB
    if lenA > lenB {
        for i:=0; i< lenA - lenB; i++ {
            pA = pA.Next
        }
    } else if lenB > lenA {
        for i:=0; i< lenB - lenA; i++ {
            pB = pB.Next
        }
    }
    for pA != nil && pB != nil {
        if pA == pB {
            return pA
        }
        pA = pA.Next
        pB = pB.Next
    }
    return nil
}
