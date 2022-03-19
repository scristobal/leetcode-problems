const dict = {
    I: 1,
    V: 5,
    X: 10, 
    L: 50, 
    C: 100, 
    D: 500, 
    M: 1000
}

function sub(act: string, prev: string): boolean{
    return dict[act] > dict[prev]
}


function romanToInt(s: string): number {
    
    const symbols = Array.from(s)
    
    return symbols.reduce(function(acc, el, ind, arr) {
        
        const greedy = acc + dict[el]
                
        if ((ind>0) && sub(el, arr[ind-1]) ){
            
            return greedy - 2*dict[arr[ind-1]]
        }
                   
        return greedy

    }, 0)
    
};