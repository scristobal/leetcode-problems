use std::collections::HashMap;

impl Solution {
    pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
        
        let mut visited  = HashMap::<i32,i32>::new();
        
        for (i, v) in nums.iter().enumerate() {
            
            let diff = target-v;
            
            match visited.get(&diff){
                Some(j) => return vec![*j,i as i32],
                None => visited.insert(*v, i as i32),
            };
        }
        
        vec![]
    }
}