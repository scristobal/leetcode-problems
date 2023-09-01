/**
 * @param {any[]} arr
 * @param {number} depth
 * @return {any[]}
 */
var flat = function (arr, n) {

    const res = [];

    function diveIn(input, level) {
        for(const item of input) {
            if(Number.isInteger(item)) {
                res.push(item);
                continue;
            }  
            
            if(level >= n) {
                res.push(item);
                continue;
            } 
                
            diveIn(item, level+1);
            
        }
    }

    diveIn(arr, 0);

    return res;
};
   
