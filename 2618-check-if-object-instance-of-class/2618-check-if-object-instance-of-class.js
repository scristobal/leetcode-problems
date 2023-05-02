/**
 * @param {any} object
 * @param {any} classFunction
 * @return {boolean}
 */
var checkIfInstanceOf = function(obj, classFunction) {
    console.log(obj, classFunction)
   return obj !== null && obj !== undefined && typeof classFunction === 'function'  &&  Object(obj) instanceof classFunction 
};

/**
 * checkIfInstanceOf(new Date(), Date); // true
 */
