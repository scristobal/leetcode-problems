func mySqrt(x int) int {
    
    if x == 0 {
		return 0
	}

    est := x

	for est*est > x {
		
		est = (est + (x / est)) / 2
	}

	return est
}