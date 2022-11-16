/** 
 * Forward declaration of guess API.
 * @param {number} num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * var guess = function(num) {}
 */


function guessNumber(n: number): number {
    let l = 1;
    let h = n;
    
    let g = Math.floor(l+(h-l)/2)
    let r = guess(g);
    
    while ( r!== 0){
       
        if (r === -1 ){ 
            h = g;
        } else {
            if (g !== l) {
                l = g;
            } else {
                l = g + 1;
            }
            
        }
        
        g = Math.floor(l+(h-l)/2);
        r = guess(g)
    }
    
    return g;
};