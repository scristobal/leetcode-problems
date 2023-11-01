// Definition for a binary tree node.
// #[derive(Debug, PartialEq, Eq)]
// pub struct TreeNode {
//   pub val: i32,
//   pub left: Option<Rc<RefCell<TreeNode>>>,
//   pub right: Option<Rc<RefCell<TreeNode>>>,
// }
// 
// impl TreeNode {
//   #[inline]
//   pub fn new(val: i32) -> Self {
//     TreeNode {
//       val,
//       left: None,
//       right: None
//     }
//   }
// }
use std::rc::Rc;
use std::cell::RefCell;
use std::collections::HashMap;
impl Solution {
    pub fn find_mode(root: Option<Rc<RefCell<TreeNode>>>) -> Vec<i32> {
         let mut count = HashMap::new();
        let mut max = 0;
        let mut stack = vec![];

        stack.push(root);

        while let Some(node) = stack.pop() {
            match node {
                Some(node) => {
                    let node = node.borrow();

                    count.entry(node.val).and_modify(|e| *e += 1).or_insert(1);

                    if count[&node.val] > max {
                        max = count[&node.val];
                    }

                    stack.push(node.left.clone());
                    stack.push(node.right.clone());
                }
                None => continue,
            }
        }

        count
            .into_iter()
            .filter(|(_, v)| *v == max)
            .map(|(k, _)| k)
            .collect()
    }
}