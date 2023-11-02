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
impl Solution {

    pub fn average_of_subtree(root: Option<Rc<RefCell<TreeNode>>>) -> i32 {
        let mut ans = 0;
        dfs(&mut ans, root);
        ans
    }
}


 fn dfs(ans: &mut i32, node: Option<Rc<RefCell<TreeNode>>>) -> (i32, i32) {
    if let Some(node) = node {
        let node = node.borrow();

        let (left_sum, left_count) = dfs(ans, node.left.clone());
        let (right_sum, right_count) = dfs(ans, node.right.clone());
        let sum = left_sum + right_sum + node.val;
        let count = left_count + right_count + 1;

        if (sum / count) == node.val {
            *ans += 1;
        }
        (sum, count)
    } else {
        (0, 0)
    }
}
