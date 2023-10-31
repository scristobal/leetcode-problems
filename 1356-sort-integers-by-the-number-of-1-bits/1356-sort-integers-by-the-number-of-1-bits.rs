impl Solution {
    pub fn sort_by_bits(arr: Vec<i32>) -> Vec<i32> {
    let mut a = arr.clone();
        
    a.sort_by(|&a, &b| {
        let ones_a = i32::count_ones(a);
        let ones_b = i32::count_ones(b);

        match ones_a.partial_cmp(&ones_b) {
            Some(ord) => match ord {
                std::cmp::Ordering::Less => std::cmp::Ordering::Less,
                std::cmp::Ordering::Equal => a.partial_cmp(&b).unwrap(),
                std::cmp::Ordering::Greater => std::cmp::Ordering::Greater,
            },
            None => unreachable!(),
        }
    });
    a
}
}