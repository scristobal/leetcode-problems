func waysToMakeFair(nums []int) int {
    var headEvenSum, headOddSum, tailEvenSum, tailOddSum int
    
    for i, n := range nums {
        if (i % 2) == 0 {
            tailEvenSum += n
            continue
        }
        tailOddSum += n
    }
    
    count := 0
    
    for i, n := range nums {
        
        if (i % 2) == 0 {
            tailEvenSum -= n
        } else {
            tailOddSum -=n
        }
        
        if (headEvenSum + tailOddSum) == (headOddSum + tailEvenSum) { 
            count +=1
        }
        
        if (i % 2) == 0{
            headEvenSum +=n
        } else {
            tailOddSum -= n
        }
        
    }
    
    return count 
}