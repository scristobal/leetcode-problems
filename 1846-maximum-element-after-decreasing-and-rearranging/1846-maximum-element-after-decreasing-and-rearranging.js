/**
 * @param {number[]} arr
 * @return {number}
 */
var maximumElementAfterDecrementingAndRearranging = function(nums) {
    nums.sort((a,b) => a - b);
    
    let res = 0;
    
    for (const num of nums){
        if (num >= (res + 1)) res++
    }
    
    return res;
};