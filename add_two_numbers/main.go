package add_two_numbers

import (
	"fmt"
	"math/big"
	"strconv"
)

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func toS(l *ListNode) string {
	if l.Next == nil {
		return fmt.Sprint(l.Val)
	}
	return toS(l.Next) + fmt.Sprint(l.Val)
}

func toBigInt(l *ListNode) *big.Int {
	f := big.NewInt(0)
	f.SetString(toS(l), 10)
	return f
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	n1 := toBigInt(l1)
	n2 := toBigInt(l2)
	n3 := big.NewInt(0).Add(n1, n2).String()
	var l *ListNode
	for i := 0; i < len(n3); i++ {
		ch := n3[i : i+1]
		v, _ := strconv.ParseInt(ch, 10, 64)
		l = &ListNode{
			Val:  int(v),
			Next: l,
		}
	}
	return l
}
//func add(l1 *ListNode, l2 *ListNode) *ListNode {
//	n1 := l1.Next
//	n2 := l2.Next
//	if n1 == nil && n2 == nil {
//		return &ListNode{
//			Val:  l1.Val + l2.Val,
//			Next: nil,
//		}
//	}
//	if n1 == nil {
//		n1 = &ListNode{0, nil}
//	}
//	if n2 == nil {
//		n2 = &ListNode{0, nil}
//	}
//	return &ListNode{
//		Val:  l1.Val + l2.Val,
//		Next: add(n1, n2),
//	}
//}
//func carry(l *ListNode) int {
//	if l.Next == nil {
//		p10, p1 := (l.Val / 10), (l.Val % 10)
//		l.Val = p1
//		return p10
//	}
//	l.Val += carry(l.Next)
//	p10, p1 := (l.Val / 10), (l.Val % 10)
//	l.Val = p1
//	return p10
//
//	//if l.Val >= 10 {
//	//	p10, p1 := (l.Val/10), (l.Val%10)
//	//	l.Val = p1
//	//	// 繰り上がりがある
//	//	if p10 > 0 {
//	//		if l.Next == nil {
//	//			l.Next = &ListNode{
//	//				Val: p10,
//	//				Next: nil,
//	//			}
//	//		} else {
//	//			l.Next.Val += p10
//	//		}
//	//	}
//	//}
//	//for l != nil {
//	//	l = l.Next
//	//}
//	//l = head
//}
