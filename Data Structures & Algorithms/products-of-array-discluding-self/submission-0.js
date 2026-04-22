class Solution {
    /**
     * @param {number[]} nums
     * @return {number[]}
     */
    productExceptSelf(nums) {
        const n = nums.length;

        let left = 1, right = 1;
        const prefix = [];
        const suffix = [];

        for (let i = 0; i < n; i++) {
            const j = n - 1 - i;

            prefix[i] = left;
            suffix[j] = right;

            left *= nums[i];
            right *= nums[j];
        }

        const result = [];
        for (let i = 0; i < n; i++) {
            const a = prefix[i];
            const b = suffix[i];
            result.push(a * b);
        }

        return result;
    }
}
