/** 
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
    l, h := 1, n;
    
    g  := l + (h -l) / 2;
    r := guess(g);
    
    for r != 0 {
        if r == -1 {
            h = g;
        } else {
            if l != g {
                l = g;
            } else { 
                l = g+1;
            }
        }
        
        g = l + (h -l) / 2;
        r = guess(g);
    }
    
    return g
}