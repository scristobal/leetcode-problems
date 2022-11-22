function numSquares(n: number): number {
    
    const res = Array.from({length: n+1}, (_, i) => i)
    
    for (let i=0; i<=n; i++){
        
        for (let j=1; i-j*j>=0; j++){
            res[i] = Math.min(res[i], res[i-j*j] + 1)
        }
    
    }
    
    return res[n]
};