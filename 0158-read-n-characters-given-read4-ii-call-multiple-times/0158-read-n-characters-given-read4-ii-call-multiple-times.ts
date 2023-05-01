/**
 * Definition for read4()
 * read4 = function(buf4: string[]): number {
 *     ...
 * };
 */

var solution = function(read4: any) {

    const rem : string [] = []
    let i = -1
    
    return function(buf: string[], n: number): number {
        const chars = Array(4)
        
        while (n>0) {
            
            if (rem.length === 0) {
                read4(rem)
            }
       
            chars.push(rem.shift())
        
            n--
        }
        
        
        buf[++i] = chars.join('')
        return buf[i].length
    
    };
};