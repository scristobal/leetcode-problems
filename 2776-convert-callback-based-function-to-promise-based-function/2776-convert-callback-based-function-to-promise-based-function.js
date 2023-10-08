/**
 * @param {Function} fn
 * @return {Function<Promise<number>>}
 */
var promisify = function(fn) {
 	return async function(...args) {
        
        let res
        
        const callback = (value, error) => {
            if (error) throw (error)
            res = value
        }
        
        fn(callback, ...args)
        
        return res
	}
};

/**
 * const asyncFunc = promisify(callback => callback(42));
 * asyncFunc().then(console.log); // 42
 */