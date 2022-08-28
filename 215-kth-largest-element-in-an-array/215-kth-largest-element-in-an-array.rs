
use std::collections::BinaryHeap;


impl Solution {
    pub fn find_kth_largest(nums: Vec<i32>, k: i32) -> i32 {
        let mut heap = BinaryHeap::with_capacity(nums.len());

        for num in nums.iter() {
            heap.push(*num)
        }

        for _ in 1..k {
            heap.pop();
        }

        heap.pop().unwrap()
    }
}
