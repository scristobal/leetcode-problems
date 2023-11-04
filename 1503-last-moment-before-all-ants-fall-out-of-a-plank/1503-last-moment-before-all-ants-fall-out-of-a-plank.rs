
impl Solution {
    pub fn get_last_moment(n: i32, left: Vec<i32>, right: Vec<i32>) -> i32 {
        
        std::cmp::max(
            *left.iter().max().or(Some(&0)).unwrap(), 
            right.iter().map(|l| n-l ).max().or(Some(0)).unwrap()
        )
        
       
    }
}