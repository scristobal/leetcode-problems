func computeArea(ax1 int, ay1 int, ax2 int, ay2 int, bx1 int, by1 int, bx2 int, by2 int) int {
    a :=area(ax1, ax2, ay1, ay2);
    b :=area(bx1, bx2, by1, by2);
    
    cx1 := max(ax1, bx1);
    cx2 := min(ax2, bx2);

    if cx1 > cx2 {return a + b;}
    
    cy1 := max(ay1, by1);
    cy2 := min(ay2, by2);
    
    if cy1 > cy2 {return a + b;}
    
    c :=area(cx1, cx2, cy1, cy2);
        
    return a + b -c;

}


func max(a, b int) int{
    if a>b {return a}
    return b
}

func min(a,b int) int {
    if a < b {return a}
    return b
}

func area(x1,x2,y1,y2 int) int {
    return (x2-x1) * (y2-y1);
}