func findOcurrences(text string, first string, second string) []string {
    words := strings.Split(text, " ")
    
    res := make([]string, len(words))
    j := 0
    
    for i:=0; i < len(words); i++ {
        if i+2 < len(words) && words[i] == first && words[i+1] == second {
            res[j] = words[i+2]
            j++
        }        
    }
    
    return res[:j]
}