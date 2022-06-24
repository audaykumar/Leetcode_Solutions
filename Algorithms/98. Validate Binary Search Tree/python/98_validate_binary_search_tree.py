from typing import Optional
# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

# Brute Force approach to find all nodes from LHS  and RHS and comparing it with current Node
# Improvements can be to just take the min and max values from each side of the Tree
class Solution:
    # def isValidBST(self, root: Optional[TreeNode]) -> bool:
        
    #     # check = self.checkNodes(root)
    #     check = self.improvedCheckNodes(root)
    #     return check[0]
    
    # def checkNodes(self, root):
    #     if root == None:
    #         return [True, []]
        
    #     leftarr = self.checkNodes(root.left)
    #     rightarr = self.checkNodes(root.right)
    #     if leftarr[0] and rightarr[0]:
    #         for l in leftarr[1]:
    #             if root.val <= l:
    #                 return [False, []]
            
    #         for r in rightarr[1]:
    #             if root.val >= r:
    #                 return [False, []]
            
    #         mergedarr = leftarr[1] + rightarr[1]
    #         mergedarr.append(root.val)
    #         return [True, mergedarr]
    #     else:
    #         return [False, []]
        
    # def improvedCheckNodes(self, root, left, right):
    #     if root == None:
    #         return 
    #     pass
    
    def isValidBST(self, root: Optional[TreeNode]) -> bool:
        def validBST(root, left, right):
            if root:
                print(root.val, left, right)
                if left>=root.val or root.val>=right: return False
                return validBST(root.left, left, root.val) and validBST(root.right, root.val, right)      
            return True
    
        return validBST(root, -float('inf'), float('inf')) 
    
sol = Solution()


test1 = TreeNode(2,
                 TreeNode(1),
                 TreeNode(3))

test2 = TreeNode(8,
                 TreeNode(4),
                 TreeNode(20,
                          TreeNode(15),
                          TreeNode(7)))

# [5,4,6,null,null,3,7]

test3 = TreeNode(10,
                 TreeNode(4),
                 TreeNode(20,
                          TreeNode(15,
                                   TreeNode(12,
                                            TreeNode(11)),
                                   TreeNode(16)),
                          TreeNode(25)))

test4 = TreeNode(2,
                 TreeNode(2),
                 TreeNode(2))


print(sol.isValidBST(test3))