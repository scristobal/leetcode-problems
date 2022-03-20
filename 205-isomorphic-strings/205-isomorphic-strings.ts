function isIsomorphic(s: string, t: string): boolean {
    
    const st = new Map()
    const ts = new Map()
    
    let i
    
    for ( i = 0; i < s.length; i++ ){
        
        if (!st.has(s[i]) && !ts.has(t[i])) {
            st.set(s[i], t[i])
            ts.set(t[i], s[i])
            continue
        }
        
        if (st.has(s[i]) && st.get(s[i]) !== t[i]) {break}
        
        if (ts.has(t[i]) && ts.get(t[i]) !== s[i]) {break}
    }
    
    //console.log({f: f.get(s[i]), t: t[i], i, fe: f.entries()})
    
    return (st.get(s[i]) === t[i]) && (ts.get(t[i]) === s[i])
};