func longestOddPalindromeAt(s string, n int) string {
    r := string(s[n])
    i := 1;
    for  (n-i)>=0 && (n+i) < len(s){
        if s[n-i] == s[n+i]{
            r = string(s[n-i:n+i+1])
        } else {
            break;
        }
        i++;
    }
    return r
}

func longestEvenPalindromeAt(s string, n int) string {
    r := ""
    i := 0;
    for (n-i) >=0 &&  (n+1+i) < len(s) {
        if s[n-i] == s[n+1+i]{
            r = string(s[n-i:n+1+i+1])
        } else {
            break;
        }
        i++;
    }
    return r
}

func longestPalindromeAt(s string, n int) string{
    r := longestOddPalindromeAt(s, n);
    t := longestEvenPalindromeAt(s, n);
    
    if len(r) > len(t) { return r}
    return t
}


func longestPalindrome(s string) string {
    r := ""
    for i:=0; i< len(s); i++ {
        t := longestPalindromeAt(s, i);
        if len(t) > len(r) { r = t }
    }
    return r
}