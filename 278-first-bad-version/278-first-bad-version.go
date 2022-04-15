/** 
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad 
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */    


  func firstBadVersion(n int) int {
          
    lo := 1
    
    for hi := n; lo < hi; {
   
      if mid := lo + (hi-lo)/2; isBadVersion(mid)    {
        hi = mid                                                                   
      } else {                                                                     
        lo = mid + 1                                                                  
      }            
    
        //fmt.Printf("(%v, %v)\n", lo, hi)
    }                                                                              
    
    // fmt.Printf("(%v, %v)\n", lo, hi)                                                           
    return lo                                                                      
  }    