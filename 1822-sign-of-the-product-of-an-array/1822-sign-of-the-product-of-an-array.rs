impl Solution {
    pub fn array_sign(nums: Vec<i32>) -> i32 {
        nums.iter().fold(1, |acc, &x| {
            match x {
                0 => 0,
                x if x <= 0 =>  -acc,
                _ => acc,
            }
        }) 
    }
}