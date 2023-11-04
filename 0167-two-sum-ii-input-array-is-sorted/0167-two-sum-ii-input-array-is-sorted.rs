impl Solution {  
    pub fn two_sum(numbers: Vec<i32>, target: i32) -> Vec<i32> {
        let mut i = 0;
        let mut j = numbers.len()-1;

        let mut r = numbers[i] + numbers[j];

        while r != target {
            
            match r.cmp(&target) {
                std::cmp::Ordering::Less => i +=1,
                std::cmp::Ordering::Equal => unreachable!(),
                std::cmp::Ordering::Greater => j -=1,
            }
           
            r = numbers[i] + numbers[j];
        }

        vec![i as i32 +1, j as i32 +1]
    }
}

