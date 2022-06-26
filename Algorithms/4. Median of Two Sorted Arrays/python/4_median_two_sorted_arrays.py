from typing import List

from pyparsing import nums
class Solution:
    def findMedianSortedArrays(self, nums1: List[int], nums2: List[int]) -> float:
        i = j = 0
        
        m = len(nums1)
        n = len(nums2)
        
        mid1= mid2 = -1
        if (m+n)%2 == 1:
            for _ in range((m+n)//2+1):
                if i !=m and j!=n:
                    if nums1[i] <= nums2[j]:
                        mid1 = nums1[i]
                        i+=1
                    else:
                        mid1 = nums2[j]
                        j +=1
                elif i<m:
                    mid1 = nums1[i]
                    i += 1
                else:
                    mid1 = nums2[j]
                    j += 1
            return float(mid1)
        else:
            for _ in range((m+n)//2+1):
                # hold previous mid1 value and get next mid1 value
                mid2 = mid1
                if i !=m and j!=n:
                    if nums1[i] <= nums2[j]:
                        mid1 = nums1[i]
                        i+=1
                    else:
                        mid1 = nums2[j]
                        j +=1
                elif i<m:
                    mid1 = nums1[i]
                    i += 1
                else:
                    mid1 = nums2[j]
                    j += 1
                
            return float((mid1+mid2)/2)
        
    
sol = Solution()

test1 = [[1,3],[2]]

test2 = [[1,2],[3,4]]

test3 = [[0,0],[0]]

test4 = [[],[1]]

print(sol.findMedianSortedArrays(test1[0],test1[1]))
print(sol.findMedianSortedArrays(test2[0],test2[1]))
print(sol.findMedianSortedArrays(test3[0],test3[1]))
print(sol.findMedianSortedArrays(test4[0],test4[1]))