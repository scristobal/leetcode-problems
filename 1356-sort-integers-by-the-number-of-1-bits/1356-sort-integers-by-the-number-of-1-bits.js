/**
 * @param {number[]} arr
 * @return {number[]}
 */
var sortByBits = (arr) => arr.sort((a,b) => nBits(a) - nBits(b) || a-b)


function nBits(n){
   let r = 0
   while(n) {
       r += n & 1
       n = n >> 1
   }
   return r
};

