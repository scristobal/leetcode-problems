impl Solution {
    pub fn length_of_last_word(s: String) -> i32 {
         s.split_whitespace()
            .filter(|s| !s.is_empty())
            .last()
            .map(|s| s.len() as i32)
            .unwrap_or(0)
    }
}