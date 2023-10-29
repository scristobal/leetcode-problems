/**
 * @param {Function} fn
 * @return {Object}
 */
Array.prototype.groupBy = function(fn) {
    
    const o = new Object();
    
    for (const i of this) {
        const r = fn(i);
        if (o[r] !== undefined) {
            o[r].push(i)
        } else {
            o[r] = [i]
        }
    }
    
    return o
};

/**
 * [1,2,3].groupBy(String) // {"1":[1],"2":[2],"3":[3]}
 */