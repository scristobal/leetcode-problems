func findArray(pref []int) []int {
    res := make([]int, len(pref))
    
    res[0] = pref[0]
    
    for j:=1; j<len(pref); j++ {
        res[j] = pref[j-1] ^ pref[j]
    }
    
    return res
}




// 0^a = 0 => a=0 --> 0*0=0
// 0^a = 1 => a=1 --> 0*1=1
// 1^a = 0 => a=1 --> 1*0=1
// 1^a = 1 => a=0 --> 1*1=0

//   5 ^   2 = 7         5 ^   7 =   2
// 101 ^ 010 = 111  -> 101 ^ 111 = 010 
