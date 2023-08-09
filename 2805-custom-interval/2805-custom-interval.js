const cleanupFns = new Map(); // id (string) -> clean function

/**
 * @param {Function} fn
 * @param {number} delay
 * @param {number} period
 * @return {number} id
 */
function customInterval(fn, delay, period){
    
    const id = cleanupFns.size;
    let count = 0;
    
    const interval = () => {
        fn()
        const timeoutId = setTimeout(interval, delay + period * ++count)
        cleanupFns.set(id, () => clearInterval(timeoutId))
    }
          
    const timeoutId  = setTimeout(interval, delay + period * count)
    cleanupFns.set(id, () => clearTimeout(timeoutId))
    
    return id
}

/**
 * @param {number} id
 */
function customClearInterval(id) {
    cleanupFns.get(id).call()
}