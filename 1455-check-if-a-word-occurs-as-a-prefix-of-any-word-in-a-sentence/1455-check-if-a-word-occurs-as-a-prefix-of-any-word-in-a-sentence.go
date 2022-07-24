func isPrefixOfWord(sentence string, searchWord string) int {
    words := strings.Split(sentence, " ")
    
    for i, w := range(words) {   
        if strings.HasPrefix(w, searchWord) { return i+1 }
    }
    
    return -1
}