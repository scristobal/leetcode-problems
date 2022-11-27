

function numberOfArithmeticSlices(nums: number[]): number {
    const seqs = []
    let total = 0
    
    for (let i=0; i<nums.length; i++) {
        
        for (let j=0; j<i; j++){
            const d = nums[i] - nums[j];
            
            const prev = (seqs[i] && seqs[i][d]) || 0
            
            if (seqs[i] === undefined) seqs[i] = []
            seqs[i][d] = prev + 1
            
            if (seqs[j] !== undefined && seqs[j][d] !== undefined) {
                if (seqs[i] === undefined) seqs[i] = []
                seqs[i][d] += seqs[j][d]
                total += seqs[j][d]
            }
            
        }
    }
    
    return total 
};