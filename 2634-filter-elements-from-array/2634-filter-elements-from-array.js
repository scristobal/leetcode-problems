/**
 * @param {number[]} arr
 * @param {Function} fn
 * @return {number[]}
 */
var filter = function(arr, fn) {
    const res = []
    for (const [i, num] of arr.entries()){
        if (fn(num, i)) res.push(num)
    }
    return res
};