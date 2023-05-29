package add_two_numbers

type ListNode struct {
	Val  int
	Next *ListNode
}

type List struct {
	head *ListNode
}

func (l *List) add(val int) {

	node := &ListNode{Val: val}
	if l.head == nil {
		l.head = node
		return
	}
	cur := l.head
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = node
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var v1, v2, output, acum int
	r := &List{}

	c1 := l1
	c2 := l2
	acum = 0
	for c1 != nil || c2 != nil || acum != 0 {
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
		output = v1 + v2 + acum
		r.add(output % 10)
		acum = output / 10
	}

	return r.head
}
