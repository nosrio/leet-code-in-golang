package add_two_numbers

import (
	"strconv"
	"testing"
)

func compareList(l1 *ListNode, l2 *ListNode) bool {
	var v1, v2 int
	c1 := l1
	c2 := l2
	for c1 != nil || c2 != nil {
		v1 = 0
		v2 = 0
		if c1 != nil {
			v1 = c1.Val
			c1 = c1.Next
		}
		if c2 != nil {
			v2 = c2.Val
			c2 = c2.Next
		}
		if v1 != v2 {
			return false
		}
	}
	return true
}

func printList(l *ListNode) string {
	var str string
	for cur := l; cur != nil; cur = cur.Next {
		str += strconv.Itoa(cur.Val)
	}
	return str
}

func TestExample1(t *testing.T) {
	var l1, l2, output = &List{}, &List{}, &List{}
	var num int
	for _, num = range []int{2, 4, 3} {
		l1.add(num)
	}
	for _, num = range []int{5, 6, 4} {
		l2.add(num)
	}
	for _, num = range []int{7, 0, 8} {
		output.add(num)
	}
	got := addTwoNumbers(l1.head, l2.head)
	if !compareList(output.head, got) {
		t.Errorf("addTwoNumbers(%v,%v) = %v; want %v", printList(l1.head), printList(l2.head), printList(got), printList(output.head))
	}
}

func TestExample2(t *testing.T) {
	var l1, l2, output = &List{}, &List{}, &List{}
	var num int
	for _, num = range []int{0} {
		l1.add(num)
	}
	for _, num = range []int{0} {
		l2.add(num)
	}
	for _, num = range []int{0} {
		output.add(num)
	}
	got := addTwoNumbers(l1.head, l2.head)
	if !compareList(output.head, got) {
		t.Errorf("addTwoNumbers(%v,%v) = %v; want %v", printList(l1.head), printList(l2.head), printList(got), printList(output.head))
	}
}

func TestExample3(t *testing.T) {
	var l1, l2, output = &List{}, &List{}, &List{}
	var num int
	for _, num = range []int{9, 9, 9, 9, 9, 9, 9} {
		l1.add(num)
	}
	for _, num = range []int{9, 9, 9, 9} {
		l2.add(num)
	}
	for _, num = range []int{8, 9, 9, 9, 0, 0, 0, 1} {
		output.add(num)
	}
	got := addTwoNumbers(l1.head, l2.head)
	if !compareList(output.head, got) {
		t.Errorf("addTwoNumbers(%v,%v) = %v; want %v", printList(l1.head), printList(l2.head), printList(got), printList(output.head))
	}
}

func TestExample4(t *testing.T) {
	var l1, l2, output = &List{}, &List{}, &List{}
	var num int
	for _, num = range []int{5, 6} {
		l1.add(num)
	}
	for _, num = range []int{5, 4, 9} {
		l2.add(num)
	}
	for _, num = range []int{0, 1, 0, 1} {
		output.add(num)
	}
	got := addTwoNumbers(l1.head, l2.head)
	if !compareList(output.head, got) {
		t.Errorf("addTwoNumbers(%v,%v) = %v; want %v", printList(l1.head), printList(l2.head), printList(got), printList(output.head))
	}
}

func TestExample5(t *testing.T) {
	var l1, l2, output = &List{}, &List{}, &List{}
	var num int
	for _, num = range []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1} {
		l1.add(num)
	}
	for _, num = range []int{5, 6, 4} {
		l2.add(num)
	}
	for _, num = range []int{6, 6, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1} {
		output.add(num)
	}
	got := addTwoNumbers(l1.head, l2.head)
	if !compareList(output.head, got) {
		t.Errorf("addTwoNumbers(%v,%v) = %v; want %v", printList(l1.head), printList(l2.head), printList(got), printList(output.head))
	}
}
