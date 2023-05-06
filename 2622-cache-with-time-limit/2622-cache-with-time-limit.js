
var TimeLimitedCache = function() {    
    this.values = new Map();
    this.timeouts = new Map();

};

/** 
 * @param {number} key
 * @param {number} value
 * @param {number} time until expiration in ms
 * @return {boolean} if un-expired key already existed
 */
TimeLimitedCache.prototype.set = function(key, value, duration) {

    const stored = this.values.get(key)
    clearTimeout(this.timeouts.get(key))

    const timeout = setTimeout(() => this.values.delete(key), duration)

    this.values.set(key, value)
    this.timeouts.set(key, timeout)

    return stored !== undefined
    
};

/** 
 * @param {number} key
 * @return {number} value associated with key
 */
TimeLimitedCache.prototype.get = function(key) {
     return this.values.get(key) ??  -1 ;
};

/** 
 * @return {number} count of non-expired keys
 */
TimeLimitedCache.prototype.count = function() {
     return this.values.size
};

/**
 * Your TimeLimitedCache object will be instantiated and called as such:
 * var obj = new TimeLimitedCache()
 * obj.set(1, 42, 1000); // false
 * obj.get(1) // 42
 * obj.count() // 1
 */