/**
 * @param {Function} fn
 * @return {Object}
 */
Array.prototype.groupBy = function(fn) {
    
    const o = {};
    
    for (let i=0; i<this.length; i++) {
        const r = fn(this[i]);
        
        o[r]? o[r].push(this[i]): o[r] = [this[i]];
    }
    
    return o
};

/**
 * [1,2,3].groupBy(String) // {"1":[1],"2":[2],"3":[3]}
 */