func newLocator(rows int) func(int) int {
    
    if rows == 1 { 
        return func(ind int) int { 
            return 0 
        } 
    }
    
    last:=-1
    inc:=true
    return func(ind int) int {
        if inc {
            if last < rows-1 {
                last++
            } else {
                inc=false
                last--
            }
        } else {
            if last>0 {
                last--
            } else {
                inc=true
                last++
            }
        }
        
        return last
    }
}

func convert(s string, numRows int) string {
    
    loc := newLocator(numRows)

    rows := make([]string, numRows)
    
    for j, c := range s {
        rows[loc(j)] += string(c)
    }

    var r string
   
    for i:=0; i<numRows; i++ {
        r += rows[i]
    }
    
    return string(r)
}


