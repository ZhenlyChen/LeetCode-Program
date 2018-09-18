package Algorithms

func Test23() {

}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	count := len(lists)
	if count == 0 {
		return nil
	} else if count == 1 {
		return lists[0]
	} else if count == 2 {
		res := new(ListNode)
		current := res
		la := lists[0]
		lb := lists[1]
		if la == nil && lb == nil {
			return nil
		}
		for la != nil && lb != nil {
			if current.Next != nil {
				current = current.Next
			}
			if la.Val < lb.Val {
				current.Val = la.Val
				la = la.Next
			} else {
				current.Val = lb.Val
				lb = lb.Next
			}
			current.Next = new(ListNode)
		}
		for la != nil {
			if current.Next != nil {
				current = current.Next
			}
			current.Val = la.Val
			current.Next = new(ListNode)
			la = la.Next
		}
		for lb != nil {
			if current.Next != nil {
				current = current.Next
			}
			current.Val = lb.Val
			current.Next = new(ListNode)
			lb = lb.Next
		}
		current.Next = nil
		return res
	} else {
		return mergeKLists([]*ListNode{mergeKLists(lists[:count/2]), mergeKLists(lists[count/2:])})
	}
}