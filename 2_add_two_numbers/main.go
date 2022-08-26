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
