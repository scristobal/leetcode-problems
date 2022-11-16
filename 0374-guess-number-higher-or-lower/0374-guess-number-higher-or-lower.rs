/** 
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * unsafe fn guess(num: i32) -> i32 {}
 */

impl Solution {
    unsafe fn guessNumber(n: i32) -> i32 {
        
        let (mut low, mut high) = (1, n + 1);
    
        loop {
            
            let mut mid = low + (high-low)/2;
            
            match guess(mid) {
                -1 => high = mid,
                1 => low = mid + 1,
                0 => break mid,
                _ => break -1,
            }    
        }
        
    }
}