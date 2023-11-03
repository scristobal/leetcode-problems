
static PUSH: &str = "Push";
static POP: &str = "Pop";

impl Solution {
 pub fn build_array(target: Vec<i32>, n: i32) -> Vec<String> {
    let mut operations: Vec<String> = vec![];
     
    operations.reserve(target.len() * 2);

    let mut next: i32 = 1;

    for i in target.into_iter() {
        operations.push(PUSH.to_string());

        for _ in next..i {
            operations.push(POP.to_string());
            operations.push(PUSH.to_string());
        }

        next = i + 1;
    }

    operations
}
}


