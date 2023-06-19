class Solution {
public:
    vector<int> twoSum(vector<int>& nums, int target) {
	
        unordered_map<int, int> seen; 
        
	    for (int i = 0; i < nums.size(); i++) 
            {
            int num = target - nums[i];

            if (seen.find(num) != seen.end()) 
                return  {seen[num] ,i };

            seen[nums[i]] = i;
        }
        return {};
    }
};