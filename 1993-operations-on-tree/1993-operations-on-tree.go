

type LockingTree struct {
	usr nodes   // users locking, 0 = free
	d   []nodes // list of descendants
	p   nodes   // parents

}

type nodes []int

func Constructor(parent []int) LockingTree {

	size := len(parent)

	tree := LockingTree{make(nodes, size), make([]nodes, size), make(nodes, size)}

	for i, p := range parent {
		tree.p[i] = p

	}

	for i := 0; i < len(parent); i++ {

		for p := parent[i]; p != -1; p = parent[p] {
			tree.d[p] = append(tree.d[p], i)
		}

	}

	fmt.Println(tree)
	return tree
}

func (t *LockingTree) Lock(n, user int) bool {

	if t.usr[n] == 0 {
		t.usr[n] = user

		return true
	}
	return false
}

func (t *LockingTree) Unlock(n, user int) bool {

	if t.usr[n] == user {
		t.usr[n] = 0
		return true
	}
	return false
}

func (t *LockingTree) Upgrade(n, user int) bool {

	if t.usr[n] != 0 {
		//fmt.Println("node  locked")
		return false
	}

	for m := t.p[n]; m != -1; m = t.p[m] {
		if t.usr[m] != 0 {
			//fmt.Println("ancestor locked")
			return false
		}
	}

	l := false
	for _, d := range t.d[n] {
		if t.usr[d] != 0 {
			l = true
			break
		}
	}

	if !l {
		//fmt.Println("no descendant is locked")
		return false
	}

	t.Lock(n, user)

	for _, d := range t.d[n] {
		t.Unlock(d, t.usr[d])
	}

	return true

}
