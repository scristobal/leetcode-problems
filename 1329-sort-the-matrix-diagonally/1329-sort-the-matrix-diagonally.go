func diag(mat [][]int, i,j int) []int {

    if i >= len(mat) { return []int{} }
    
    if j >= len(mat[i]) { return []int{} }
    
    d := make([]int, len(mat) + len(mat[i]))
    k:=0;
    
    for i+k < len(mat) && j+k < len(mat[i+k]) {
        d[k] = mat[i+k][j+k]
        k++
    }

    return d[:k]
}


func replace(mat [][]int, d []int, i,j int)  {
    
    if i >= len(mat) { return }
    
    if j >= len(mat[i]) { return } 
    
    for k:=0; i+k < len(mat) && j+k < len(mat[i+k]); k++ {
        mat[i+k][j+k] = d[k]
    }
}

func diagonalSort(mat [][]int) [][]int {
    
    if len(mat) == 0 { return mat }
    
    for i:=0; i< len(mat[0]); i++ {
        d := diag(mat,0,i)
        sort.Slice(d, func(i,j int) bool { return d[i] < d[j]})
        replace(mat, d, 0, i)
    }
    
    for i:=1; i<len(mat); i++{
        d := diag(mat,i,0)
        sort.Slice(d, func(i,j int) bool { return d[i] < d[j]})
        replace(mat, d, i, 0)
    }
    
    return mat
}