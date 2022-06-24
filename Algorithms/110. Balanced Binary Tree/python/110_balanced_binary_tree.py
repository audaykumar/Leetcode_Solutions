from tabnanny import check
from typing import Optional
# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right
        
class Solution:
    def isBalanced(self, root: Optional[TreeNode]) -> bool:
        if not root:
            return True
        
        left_height = self.checkHeight(root.left)
        right_height = self.checkHeight(root.right)
        
        if abs(left_height - right_height) <=1:
            return self.isBalanced(root.left) and self.isBalanced(root.right)
        else:
            return False
        
    
    def checkHeight(self, root):
        if not root:
            return 0
        left_height = self.checkHeight(root.left)
        right_height = self.checkHeight(root.right)
        
        return 1 + max(left_height, right_height)
        
    
sol = Solution()

test1 = TreeNode(3,
                 TreeNode(9),
                 TreeNode(20,
                          TreeNode(15),
                          TreeNode(7)))

print(sol.isBalanced(test1))

