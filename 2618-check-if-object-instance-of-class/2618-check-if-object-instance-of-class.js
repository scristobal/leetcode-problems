/**
 * @param {any} object
 * @param {any} classFunction
 * @return {boolean}
 */
var checkIfInstanceOf = function(obj, classFunction) {
    
    if (isNullable(obj) || isNullable(classFunction)) { return false }
    
    let proto = Object.getPrototypeOf(obj) 
    
    while (proto){
        if (proto === classFunction.prototype){ return true }
        
        proto = Object.getPrototypeOf(proto) 
    }
    return false
};

/**
 * checkIfInstanceOf(new Date(), Date); // true
 */

const isNullable = function(obj) {
    return (obj ?? true) === true
}