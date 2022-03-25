function firstUniqChar(s: string): number {
    const charCount = new Map()
    
    const arr = Array.from(s)
    
    let ind
    
    for (ind = 0; ind < arr.length; ind++){
        if (charCount.has(arr[ind])){
            const current = charCount.get(arr[ind])
            charCount.set(arr[ind],[current[0], current[1] + 1])
        } else {
            charCount.set(arr[ind], [ind, 1])
        }
    }
    
    let pairs
    
    for (pairs of charCount.values()){
        console.log(pairs)
        if (pairs[1] === 1) break
    }
    
    return pairs[1]===1? pairs[0]:-1
};