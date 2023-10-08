/**
 * @param {Array<Function>} functions
 * @return {Promise<Array>}
 */
var promiseAllSettled = function(funcs) {
    
    return new Promise((accept, reject) => {
        
        const res = new Array(funcs.length);
        let completed = 0;
    
        for (const [i, func] of funcs.entries()) {
            func()
                .then(v => {
                    res[i] = { status: "fulfilled", value: v }
                    completed++
                })
                .catch(e => {
                    res[i] = { status: "rejected", reason: e}
                    completed++
                })
                .finally(() => {
                    if (completed === funcs.length) accept(res)
                })
        }     
       
    })
};


/**
 * const functions = [
 *    () => new Promise(resolve => setTimeout(() => resolve(15), 100))
 * ]
 * const time = performance.now()
 *
 * const promise = promiseAllSettled(functions);
 *              
 * promise.then(res => {
 *     const out = {t: Math.floor(performance.now() - time), values: res}
 *     console.log(out) // {"t":100,"values":[{"status":"fulfilled","value":15}]}
 * })
 */