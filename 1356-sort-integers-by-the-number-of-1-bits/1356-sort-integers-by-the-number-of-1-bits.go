func sortByBits(arr []int) []int {
   
    sort.Slice(arr,func(i, j int) bool {
        
        a, b := bits.OnesCount(uint(arr[i])), bits.OnesCount(uint(arr[j]));
        
        switch {
            case a==b:
                return arr[i]<arr[j]
            default:
                return a<b
            
        }
	})
    
    return arr
}

// func nBits(n int) int {
//     r := 0;
    
//     for (n > 0) {
//         r += n & 1;
//         n = n >> 1;
//     }
//     return r
// }