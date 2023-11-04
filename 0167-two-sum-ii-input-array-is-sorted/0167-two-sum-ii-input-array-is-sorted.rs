impl Solution {  
    pub fn two_sum(numbers: Vec<i32>, target: i32) -> Vec<i32> {
        let mut i = 0;
        let mut j = numbers.len()-1;

        loop {
            match target.cmp(&(numbers[i] + numbers[j])) {
                std::cmp::Ordering::Less => j -=1,
                std::cmp::Ordering::Equal => break,
                std::cmp::Ordering::Greater => i +=1,
            }
        }

        vec![i as i32 +1, j as i32 +1]
    }
}

