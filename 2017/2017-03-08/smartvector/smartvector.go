package smartvector

import "fmt"

type SmartVector struct {
	capacity int
	data     *node
}

var _ Vector = (*SmartVector)(nil)

func (v *SmartVector) String() string {
	return fmt.Sprintf("{cap: %d, data: %v}", v.capacity, v.data)
}

type node struct {
	value       int
	first, last int
	next        *node
}

func (n *node) String() string {
	return fmt.Sprintf("(%v [%d,%d] %v)", n.value, n.first, n.last, n.next)
}

func (n *node) Set(i, value int) {

}

func (n *node) Get(i int) int {
	if n.first <= i && i <= n.last {
		return n.value
	}
	if n.next == nil {
		return 0
	}
	return n.next.Get(i)
}

func NewSmartVector(n int) Vector {
	return &SmartVector{capacity: n}
}

func (v *SmartVector) Set(i, value int) {
	v.checkIndex(i)

	// if v.data == nil || i < v.data.first-1 {
	// 	v.data = &node{
	// 		value: value,
	// 		first: i,
	// 		last:  i,
	// 		next:  v.data,
	// 	}
	// 	return
	// }

	// find node where to store this value
	n := v.data
loop:
	for n != nil {
		switch {
		case i < n.first-1:
			n = &node{
				value: value,
				first: i,
				last:  i,
				next:  n,
			}
			break loop
		case n.first-1 <= i && i <= n.last+1:
			switch {
			case value == n.value:
				break loop
			case i == n.first-1:
			}

		}
		n = n.next
	}

	// append node to the end of the linked list
	if n == nil {
		n = &node{
			value: value,
			first: i,
			last:  i,
		}
		return
	}

	// either update the index...
	if value == n.value {
		n.first = min(n.first, i)
		n.last = max(n.last, i)
		return
	}

	// ... or split...
	if n.first < i && i < n.last {
		// ... into 3 nodes
		n.next = &node{
			value: value,
			first: i,
			last:  i,
			next: &node{
				value: n.value,
				first: i + 1,
				last:  n.last,
			},
		}
		n.last = i - 1
		return
	}
	// ... into 2 nodes
	if i >= n.last {
		n = &node{}
	}

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (v *SmartVector) Get(i int) int {
	v.checkIndex(i)
	if v.data == nil {
		// mimic a []int, all valid indices are implicitly initialized to the
		// zero value.
		return 0
	}
	return v.data.Get(i)
}

func (v *SmartVector) checkIndex(i int) {
	if i < 0 || i >= v.capacity {
		panic(fmt.Sprintf("invalid index %d: not in range [0, %d)", i, v.capacity))
	}
}
