


func addDigit(x int, d int) int {
    
    return x*10 + d
}

func removeLastDigit(x int)( y int, d int) {
    d = x % 10;
    y = (x-d) / 10;
    return y, d
}

func reverse(x int) int {
    
    var r int
    
    if x >= 0 {
        r = x
    } else {
        r = -x
    }
    
    d := 0;
    y := 0; 
    
    for r > 0 {
        r, d = removeLastDigit(r)
        
        y = addDigit(y, d)
        
    }
    
    if y <= math.MaxInt32 { 
        if x >= 0 { return y} else {return -y}
    }
    
    return 0
}