func isIsomorphic(s string, t string) bool {
    st := make(map[byte]byte)
    ts := make(map[byte]byte)
    
    if len(s) != len(t) { return false}
     
    for i := 0; i < len(s); i ++  {
        _, isInST := st[s[i]]
        _, isInTS := ts[t[i]]
        
        if (!isInST && !isInTS) {
            st[s[i]] = t[i]
            ts[t[i]] = s[i]
            continue
        }
        
        if isInST && st[s[i]]!=t[i] { return false }
        if isInTS && ts[t[i]]!=s[i] { return false}
    }
    
    return true
}
