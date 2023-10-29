/**
 * @param {number} n
 * @param {number} k
 * @return {number}
 */
var kthGrammar = function(n, k) {
    n--
    k--

    const seq = new Array(n)
    
    for (let i =0; i<n; i++) {
        seq[i] = k%2;
        k = Math.floor(k/2)
    }

    let p = 0;
    for (const m of seq){
        p = p!=m // succesor m is 0=L, 1=R
    }

    return p
};
