function isPalindrome(s: string): boolean {
    const t = s.toLowerCase().replace(/[^0-9a-z]/g,'');
    
    const r = t.split("").reverse().join("")
    
    return r===t
};