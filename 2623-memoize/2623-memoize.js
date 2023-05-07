/**
 * @param {Function} fn
 */
function memoize(fn) {
    
    const cache = new Map()
    
    return function(...args) {
        
        let res = cache.get(JSON.stringify(args))
        
        if (res === undefined) {
            res = fn(...args)
            cache.set(JSON.stringify(args), res)
        } 
        
        return res
    }
}


/** 
 * let callCount = 0;
 * const memoizedFn = memoize(function (a, b) {
 *	 callCount += 1;
 *   return a + b;
 * })
 * memoizedFn(2, 3) // 5
 * memoizedFn(2, 3) // 5
 * console.log(callCount) // 1 
 */