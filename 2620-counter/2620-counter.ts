function createCounter(n: number): () => number {

    let state = n;
    
    return function() {
        return state++
    };
}


/** 
 * const counter = createCounter(10)
 * counter() // 10
 * counter() // 11
 * counter() // 12
 */