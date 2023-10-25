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
        p = succ(p,m)
    }

    return p
};


function succ(n, s /* L=0  R =1 */) {
        return n != s ? 1 : 0
        // 0,0 -> 0
        // 0,1 -> 1
        // 1,0 -> 1
        // 1,1 -> 0
        if (n ===0) { // 0 -> 01
            if (s === 0) return 0
            return 1
        } // 1 -> 10
        if (s === 0) return 1
        return 0
    }