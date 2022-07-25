func addBinary(a string, b string) string {   
    
    one, zero := byte('1'), byte('0')
    
    if len(a) < len(b) { a, b = b, a}

    r := make([]byte, len(a) + 1)
    k:=len(a)
    
    i, c := len(a)-1, false
    for j:= len(b)-1; j >=0; j-- {
        
        p := a[i] == one
        q := b[j] == one
        
        if (p && q && c) || (p && !q && !c) || (!p && q && !c) || (!p && !q && c) {
            r[k] = one
        } else {
            r[k] = zero
        }
        
        c = (p && q && !c) || (p && !q && c) || (!p && q && c) || (p && q && c)
        
        k--
        i--
    }
    
    for i>=0 {
        p := a[i] == one
    
        if (p && !c) || (!p && c) {
            r[k] = one
        } else {
            r[k] = zero
        }
        
        c = (p && c)
        
        k--
        i--
    }
    
    if c {
        r[k] = one 
        k--
    }

    return string(r[k+1:])
}

