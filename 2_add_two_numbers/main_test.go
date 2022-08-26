package add_two_numbers

import (
	"reflect"
	"strconv"
	"testing"

	"github.com/k0kubun/pp"
)

func toListNode(s string) *ListNode {
	var nodes *ListNode
	var p *ListNode // parent
	for i := range s {
		v, _ := strconv.ParseInt(s[i:i+1], 10, 64)
		nd := &ListNode{
			Val:  int(v),
			Next: nil,
		}
		if nodes == nil {
			nodes = nd
			p = nodes
		} else {
			p.Next = nd
			p = p.Next
		}
	}
	return nodes
}

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{ // 321 + 765 = 1086 => 6801
			args: args{
				l1: toListNode("123"),
				l2: toListNode("567"),
			},
			want: toListNode("6801"),
		},
		{ // 342 + 465 = 807 => 708
			args: args{
				l1: toListNode("243"),
				l2: toListNode("564"),
			},
			want: toListNode("708"),
		},
		{ // 0000000000000000000465 + 1000000000000000000000000000001 = 5650000000000000000000000000001
			args: args{
				l1: toListNode("1000000000000000000000000000001"),
				l2: toListNode("564"),
			},
			want: toListNode("6640000000000000000000000000001"),
		},
		{
			args: args{
				l1: toListNode("5"),
				l2: toListNode("5"),
			},
			want: toListNode("01"),
		},
		{
			args: args{
				l1: toListNode("18"),
				l2: toListNode("0"),
			},
			want: toListNode("18"),
		},
		{
			args: args{
				l1: toListNode("9"),
				l2: toListNode("1999999999"),
			},
			want: toListNode("00000000001"),
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				pp.Println(got)
				pp.Println(tt.want)
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_toS(t *testing.T) {
	type args struct {
		l *ListNode
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			args: args{
				l: toListNode("123"),
			},
			want: "321",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toS(tt.args.l); got != tt.want {
				t.Errorf("toS() = %v, want %v", got, tt.want)
			}
		})
	}
}
