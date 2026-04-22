class Solution {
    /**
     * @param {number[]} numbers
     * @param {number} target
     * @return {number[]}
     */
    twoSum(arr, target) {
        let left = 0, right = arr.length - 1;
        while(left < right){
            const sum = arr[left] + arr[right];
            if(sum == target) return [left + 1, right + 1];
            if(sum > target) right--;
            else left++;
        }
        return [];
    }
}
