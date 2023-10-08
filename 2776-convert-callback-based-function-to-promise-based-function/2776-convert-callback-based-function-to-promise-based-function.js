/**
 * @param {Function} fn
 * @return {Function<Promise<number>>}
 */
var promisify = function(fn) {
//  	return async function(...args) {
        
//         let res
        
//         const callback = (value, error) => {
//             if (error) throw (error)
//             res = value
//         }
        
//         fn(callback, ...args)
        
//         return res
// 	}
    
    return function(...args){
        return new Promise((resolve, reject) => {
            const callback = (value, error) => {
                if (error) reject(error)
                else resolve(value)
            }          
            fn(callback, ...args)
        })
    }
};

/**
 * const asyncFunc = promisify(callback => callback(42));
 * asyncFunc().then(console.log); // 42
 */