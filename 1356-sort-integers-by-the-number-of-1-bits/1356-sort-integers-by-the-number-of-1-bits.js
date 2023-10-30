/**
 * @param {number[]} arr
 * @return {number[]}
 */
var sortByBits = function(arr) {
  
    
    return arr.sort((a,b) => {
        const bA = numBits(a);
        const bB = numBits(b);
        
        const d = bA - bB;
        
        if (d!==0) return d
        else return a-b
    })
};


function numBits(n){
    if (n===0) return 0;
    
  
    
    const h = Math.floor(n/2);

    if ( (n % 2) === 1) {
        const b = numBits(h) + 1;
      
        return b
    }
    
    const b = numBits(h);
    
    return b
    
};

