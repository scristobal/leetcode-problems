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
        
        let (mut l, mut h) = (1, n);
        
        let mut g = l + (h-l)/2;
        let mut r = guess(g);
    
        while r != 0 {
            if r == -1 {
                h = g;
            } else {
                if g != l {
                    l = g;
                } else {
                    l = g + 1;
                }
            }
            
            g = l + (h-l)/2;
            r = guess(g);
        }
        
        g
    }
}