type priorityQueue []float64

func (p priorityQueue) Len() int {
	return len(p)
}

func (p priorityQueue) Less(i, j int) bool {
	return p[i] > p[j]
}

func (p priorityQueue) Swap(i, j int) {

	p[i], p[j] = p[j], p[i]
}

func (p *priorityQueue) Push(x interface{}) {
	*p = append(*p, x.(float64))
}

func (p *priorityQueue) Pop() interface{} {
	t := *p
	n := len(t)
	x := t[n-1]
	*p = (t)[:n-1]
	return x
}

func halveArray(nums []int) int {

	var pq = make(priorityQueue, len(nums))

	var sum float64 = 0

	for i, v := range nums {
		f := float64(v)
		pq[i] = f
		sum += f
	}

	heap.Init(&pq)

	r := 0
	s := sum

	for s > sum/2 {
		p := heap.Pop(&pq).(float64)
		h := p / 2
		s -= h
		heap.Push(&pq, h)
		r++
	}

	return r
}
