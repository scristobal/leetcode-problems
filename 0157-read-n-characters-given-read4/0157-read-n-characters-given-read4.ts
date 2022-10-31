/**
 * Definition for read4()
 * read4 = function(buf4: string[]): number {
 *     ...
 * };
 */

var solution = function(read4: any) {

    return function(buf: string[], n: number): number {
     
        let p = 0;
    
        while (p < n) {
            let b = ["","","",""]
            const c = read4(b)
            
            if (c===0) return p
            
            for (let i=0; i<c && p<n; i++) {
                buf[p] = b[i]
                p +=1
            }
        }
        
        return p
        
    };
};