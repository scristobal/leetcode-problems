/**
 * @param {string} s
 * @return {string}
 */
var sortVowels = function(s) {

    let vowels = s.split("").map(c => c.charCodeAt(0)).filter(c => isVowel(String.fromCharCode(c)))
    
    vowels.sort((a,b) => a-b);
        
    let res = ""
    j=0
    for (const [i, c] of s.split("").entries() ) {
       
        if (isVowel(c)) {
           
            res +=String.fromCharCode(vowels[j])
            j++
        } else {
            res+=c
        }
    }
    
    return res
};


function isVowel(c) {
    return c == 'a' || c == 'e' || c == 'o'|| c == 'u'|| c == 'i' || 
           c == 'A' || c == 'E' || c == 'O'|| c == 'U'|| c == 'I';
}