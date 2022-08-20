class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        
        visited : dict[int, int] = {}
        
        for i, v in enumerate(nums):
            
            j = visited.get(target-v)
            
            if j!=None:
                return [j,i]
            else:
                visited[v]=i
                
        return []