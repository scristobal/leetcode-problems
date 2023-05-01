function average(salary: number[]): number {
    const max = Math.max(...salary)
    const min = Math.min(...salary)
    
    const sum = salary.reduce((e, acc) => acc + e, -max -min) ;
    
    return sum/ (salary.length -2 )
};