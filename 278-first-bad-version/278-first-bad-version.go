/** 
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad 
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
        
    step:= reduce(n) 
    isPreviousBad := true
    ind := n - step
    fmt.Println("---")
    for {
        
        isCurrentBad := isBadVersion(ind)
        
        
        if ( step ==1) && isPreviousBad != isCurrentBad {
            if isCurrentBad { return ind } 
            if isPreviousBad { return ind + 1}
        }
        
        step= reduce(step) 

        isPreviousBad = isCurrentBad
        
        if isCurrentBad {
            ind -= step
            
            continue
        }
        
        ind += step
    }
    
    return -1
}

func reduce(n int) int {
    
    r:=n/2

    if r == 0 {
        return r+1
        
    }
    
    return r
}