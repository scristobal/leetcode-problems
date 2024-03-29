
impl Solution {
    pub fn get_last_moment(n: i32, left: Vec<i32>, right: Vec<i32>) -> i32 {
        left
            .into_iter()
            .chain(right
                .into_iter()
                .map(|l| n-l ))
            .max()
            .unwrap_or(0)
    }
}