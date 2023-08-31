/**
 * @param {any[]} arr
 * @param {number} depth
 * @return {any[]}
 */
var flat = function (arr, n) {
    let res = arr
    for (let i=0; i<n; i++) {
        res = flatOne(res)
    }
    return res
};

/**
 * @param {any[]} arr
 * @return {any[]}
 */
function flatOne(arr) {
    const res = []
    for (const elem of arr) {
        if (Array.isArray(elem)) {
            for (const nested of elem) {
                res.push(nested)
            }
        } else {
            res.push(elem)
        }
    }
    return res
}