from typing import Optional

# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right
        
    
class Solution:
    def maxDepth(self, root: Optional[TreeNode]) -> int:
        
        if root is None:
            return 0
            # maxLeft = self.maxDepth(root.left)
            # maxRight = self.maxDepth(root.right)
            
            # height = max(maxLeft, maxRight) + 1
        return max(self.maxDepth(root.left), self.maxDepth(root.right)) + 1
    
sol = Solution()

test1 = TreeNode(3,
                 TreeNode(9),
                 TreeNode(20,
                          TreeNode(15),
                          TreeNode(7)))

test2 = TreeNode(1,
                 left=None,
                 right=TreeNode(2))

print(sol.maxDepth(test1))
print(sol.maxDepth(test2))
