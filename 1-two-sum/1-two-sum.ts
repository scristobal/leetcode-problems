function twoSum(nums: number[], target: number): number[] {

    const cache = new Map()

    for (let i = 0; i < nums.length; i++) {
  
        let m = target - nums[i]
        
        if (cache.has(m)){
            return  [i, cache.get(m)]
        }
        
        cache.set(nums[i], i)
    }
   
};
