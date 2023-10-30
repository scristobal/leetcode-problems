func sortByBits(arr []int) []int {
   
    sort.Slice(arr,func(i, j int) bool {
        
        a, b := nBits(arr[i]), nBits(arr[j]);
        
        switch {
            case a==b:
                return arr[i]<arr[j]
            default:
                return a<b
            
        }
	})
    
    return arr
}

func nBits(n int) int {
    r := 0;
    
    for (n > 0) {
        r += n & 1;
        n = n >> 1;
    }
    return r
}