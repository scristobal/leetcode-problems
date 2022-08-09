func mySqrt(x int) int {
    
    if x == 0 {
		return 0
	}

	N := float64(x)

    prev := float64(1)
	est := N

	for math.Abs(est-prev) > 0.5 {
		prev = est
		est = (prev + (N / prev)) / 2
	}

	return int(est)
}