class Solution {
    /**
     * @param {number[]} nums
     * @param {number} target
     * @return {number[]}
     */
    twoSum(nums, target) {
        const map = new Map();

        for(const [i, num] of nums.entries()){
            const diff = target - num;
            if(map.has(diff)) return [map.get(diff), i];
            map.set(num, i);
        }
    }
}
