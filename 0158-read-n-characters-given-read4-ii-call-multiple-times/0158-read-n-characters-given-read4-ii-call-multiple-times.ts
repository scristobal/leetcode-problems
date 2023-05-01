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
        const chars = []
        
        while (n>0) {
            
            if (rem.length === 0) {
                const c = read4(rem)
                if (c===0) { break }
            }
       
            const e = rem.shift()
            chars.push(e)
            
            n--
        }
        i +=1 
        buf[i] = chars.join('')
        return buf[i].length
    
    };
};