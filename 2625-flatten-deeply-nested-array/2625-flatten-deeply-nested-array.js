/**
 * @param {any[]} arr
 * @param {number} depth
 * @return {any[]}
 */
var flat = function (arr, n) {
    
    const input = JSON.stringify(arr);
    let output = "["
    let depth = -1
    for (const char of input) {
        switch (char) {
            case '[' : {
                depth++
                if (depth > n){
                    output+=char
                }
                continue;
            }
            case ']': {
                if (depth > n){
                    output+=char
                }
                depth--
                continue;
            }
           
            default:
                output+=char
        }
    }
    
    output+="]"
    
    output=output.replaceAll(",,","").replaceAll("[,", "[").replaceAll(",]","]")

    return JSON.parse(output)
   
};
