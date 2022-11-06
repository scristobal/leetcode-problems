function getRow(ind: number, rows: number){
    
    if (rows <= 2) {
        return ind % rows
    }
    
    const mod =  ind % (rows + rows -2 )
    
    if (mod < rows ) return mod
    
    return rows + rows -2 - mod
}

function convert(str: string, numRows: number): string {
        
    const chars = Array.from(str)
    
    let rows: string[] = Array.from({length: numRows}, () => '')
    
    for (let i =0; i < str.length; i++ ){
        
        rows[getRow(i, numRows)]=rows[getRow(i, numRows)].concat(chars[i])
    }
    
    return rows.join('')
    
};