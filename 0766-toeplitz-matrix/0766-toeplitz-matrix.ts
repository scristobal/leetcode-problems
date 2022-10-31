function isToeplitzMatrix(matrix: number[][]): boolean {
    
   

    for (let i=0; (i+1) <matrix.length; i++){
    
        let row = matrix[i].slice(0, matrix[i].length-1);
        let next = matrix[i+1].slice(1, matrix[i+1].length)
        
        
        for (let j=0; j<next.length;j++){
            if (row[j] !== next[j] ) return false 
        }
      
    }
    
    return true
    
};