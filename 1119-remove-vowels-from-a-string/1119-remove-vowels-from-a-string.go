func removeVowels(s string) string {
    vowels := []rune{'a','e','i','o','u'};
    var res string;
    
    for _, v := range s {
        isVowel := false
        
        for _, w := range vowels {
            if v==w { isVowel = true; continue }
        }
        
        if !isVowel { res = res + string(v) }
    }
    
    return res 
}