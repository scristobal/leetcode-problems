function hash(d: number[]): number {
    if (d[0] > d[1]) return (d[0]*10) +  d[1]
    return (d[1]*10)+d[0]
}

function numEquivDominoPairs(ds: number[][]): number {
    
    const acc = new Map<number, number>() 
    
    for (let d of ds) {
        
        const h = hash(d)
        
        const cur = acc.get(h) ?? 0
        
        acc.set(h, cur+1)
        
    }
    
    let res = 0;
    for (let c of Array.from(acc.values())) {
        if (c > 1) {
            res += c*(c - 1)/2
        }
    }
    return res
};

