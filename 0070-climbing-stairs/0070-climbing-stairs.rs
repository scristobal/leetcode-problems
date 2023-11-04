impl Solution {
    pub fn climb_stairs(n: i32) -> i32 {
        
        let n = n as usize; 
     
        let mut steps = vec![0; n+1 ];
        
        steps[0] =1;
        steps[1] =1;
                        
        for i in 2..n+1 {
            steps[i] = steps[i-1] + steps[i-2]
        }
        
        steps[n]
    }
}