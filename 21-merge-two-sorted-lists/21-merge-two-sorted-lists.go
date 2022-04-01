func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	var h *ListNode

	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	var p1, p2 *ListNode

	if l1.Val < l2.Val {
		h = l1
		p1 = l1.Next
		p2 = l2
	} else {
		h = l2
		p2 = l2.Next
		p1 = l1
	}

	p := h

	for p1 != nil && p2 != nil {

		if p1.Val < p2.Val {

			p.Next = p1
			p = p1
			p1 = p1.Next

			continue
		}

		p.Next = p2
		p = p2
		p2 = p2.Next

	}

	if p2 == nil {
		p.Next = p1
	} else {
		p.Next = p2
	}

	return h
}