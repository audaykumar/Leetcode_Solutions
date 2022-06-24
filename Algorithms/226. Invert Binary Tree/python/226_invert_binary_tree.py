from operator import invert
from typing import Optional, List
# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right
        
class Solution:
    def invertTree(self, root: Optional[TreeNode]) -> Optional[TreeNode]:
        
        if root:
            # self.invertTree(root.left)
            # self.invertTree(root.right)
            
            # temp = root.left
            # root.left = root.right
            # root.right = temp
            
            root.left, root.right = self.invertTree(root.right), self.invertTree(root.left)
            return root
    
    
sol = Solution()

test1 = TreeNode(4,
                 TreeNode(2,
                          TreeNode(1),
                          TreeNode(3)),
                 TreeNode(7,
                          TreeNode(6),
                          TreeNode(9)))

x = sol.invertTree(test1)

print(x.right.val)