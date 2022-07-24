func findOcurrences(text string, first string, second string) []string {
    words := strings.Split(text, " ")
    
    res := make([]string, len(words))
    j := 0
    
    for i, w := range(words) {
        if i+2 < len(words) && w == first && words[i+1] == second {
            res[j] = words[i+2]
            j++
        }        
    }
    
    return res[:j]
}