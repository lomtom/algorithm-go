package offer35

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	m := make(map[*Node]*Node)
	var newHead = &Node{
		Val: head.Val,
	}
	var cur = newHead
	m[head] = cur
	for head != nil {
		if head.Next != nil {
			if v, ok := m[head.Next]; ok {
				cur.Next = v
			} else {
				cur.Next = &Node{
					Val: head.Next.Val,
				}
				m[head.Next] = cur.Next
			}
		}
		if head.Random != nil {
			if v, ok := m[head.Random]; ok {
				cur.Random = v
			} else {
				cur.Random = &Node{
					Val: head.Random.Val,
				}
				m[head.Random] = cur.Random
			}
		}
		cur = cur.Next
		head = head.Next
	}
	return newHead
}
