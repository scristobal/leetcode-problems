use std::collections::HashMap;

impl Solution {
    pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
        
        let mut visited = HashMap::new();
        
        for (i, v) in nums.iter().enumerate() {
            
            let diff = target-v;
            let prev = visited.get_key_value(&diff);
            
            match prev {
                Some((_,&j)) => return vec![j,i as i32],
                _ => visited.insert(v, i as i32),
            };
            
        }
        
        vec![0,1]
    }
}