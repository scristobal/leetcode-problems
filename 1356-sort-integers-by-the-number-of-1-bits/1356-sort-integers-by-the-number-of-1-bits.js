/**
 * @param {number[]} arr
 * @return {number[]}
 */
var sortByBits = function(arr) {
    const cache = new Map();
    
    // arr.forEach((n) => {
    //     console.log("n",n)
    //     console.log("bits", numBits(n))
    // })
    
    return arr.sort((a,b) => {
        const bA = numBits(a, cache);
        const bB = numBits(b,cache);
        
        const d = bA - bB;
        
        if (d!==0) return d
        else return a-b
    })
};


function numBits(n, cache){
    if (n===0) return 0;
    
    const c = cache.get(n)
    
    if (c) return c
    
    const h = Math.floor(n/2);

    if ( (n % 2) === 1) {
        const b = numBits(h, cache) + 1;
        cache.set(n, b)
        return b
    }
    
    const b = numBits(h,cache);
    cache.set(n,b);
    return b
    
};