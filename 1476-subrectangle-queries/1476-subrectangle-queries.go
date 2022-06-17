type SubrectangleQueries struct {
    values [][]int
}


func Constructor(rectangle [][]int) SubrectangleQueries {
    values := rectangle
    return SubrectangleQueries{values}
}


func (this *SubrectangleQueries) UpdateSubrectangle(row1 int, col1 int, row2 int, col2 int, newValue int)  {
    for r := row1; r<=row2; r++ {
        for c := col1; c <=col2; c++ {
            this.values[r][c] = newValue
        }
    }
}


func (this *SubrectangleQueries) GetValue(row int, col int) int {
    return this.values[row][col]
}


/**
 * Your SubrectangleQueries object will be instantiated and called as such:
 * obj := Constructor(rectangle);
 * obj.UpdateSubrectangle(row1,col1,row2,col2,newValue);
 * param_2 := obj.GetValue(row,col);
 */