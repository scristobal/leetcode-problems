# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def sumNumbers(self, root: Optional[TreeNode]) -> int:
        return self.sumSubtree(root, 0)
        
    def sumSubtree(self, node: Optional[TreeNode], acc: int) -> int:
        
        if node == None:
            return 0
        
        sum = acc*10 + node.val
        
        if (node.left == None) and (node.right == None):
            return sum
        
        return self.sumSubtree(node.left, sum) + self.sumSubtree(node.right, sum)