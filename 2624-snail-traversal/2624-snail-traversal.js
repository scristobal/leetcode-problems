/**
 * @param {number} rowsCount
 * @param {number} colsCount
 * @return {Array<Array<number>>}
 */
Array.prototype.snail = function(rowsCount, colsCount) {
    
    if (rowsCount * colsCount !== this.length) {
        return []
    }

    const res = new Array(rowsCount)
    
    let L = 2*rowsCount;
    
    const indices = new Uint16Array(colsCount)
    
    let offset = 0;
    let factor = 0;
    
    for (let i=0; i<colsCount; i++) {
        
        indices[i] = factor*L + offset;
        
        const delta = 1+offset
        factor += delta;
        offset = -delta;
    }
    
    //   0                            9  10                                19
    //   |                            |   |                                 |
    //   v                            v   v                                 v
    // [19, 10, 3, 7, 9, 8, 5, 2, 1, 17, 16, 14, 12, 18, 6, 13, 11, 20, 4, 15]
    
    const row = new Uint16Array(colsCount);
          
    for (let r=0; r<rowsCount; r++){
        
        for (let c=0; c<colsCount; c++){
            row[c] = this[indices[c]]
        }
        
        res[r] = Array.from(row)
        
        // update indices ...
        // r=0 -> [0*L,     1*L-1,     1*L,     2*L-1,   2*L,   ... ];
        // r=1 -> [0*L+1,   1*L-1-1,   1*L+1,   2*L-1-1, 2*L+1, ... ]
        // r=2 -> [0*L+1+1, 1*L-1-1-1, 1*L+1+1, ...]
        let offset = +1
        for (let i=0; i<colsCount; i++){
            indices[i] = indices[i] + offset
            offset*=-1
        }
    }
    
    return res
}

/**
 * const arr = [1,2,3,4];
 * arr.snail(1,4); // [[1,2,3,4]]
 */