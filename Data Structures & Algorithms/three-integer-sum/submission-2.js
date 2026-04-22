class Solution {
    /**
     * @param {number[]} nums
     * @return {number[][]}
     */
    threeSum(arr) {
        arr.sort((a, b) => a - b);

        const triplets = [];

        for (let i = 0; i < arr.length; i++) {
            if (i > 0 && arr[i - 1] === arr[i]) continue;
            let left = i + 1, right = arr.length - 1;
            while (left < right) {
                const sum = arr[i] + arr[left] + arr[right];
                if (sum > 0) {
                    right--;
                }
                else if (sum < 0) {
                    left++;
                }
                else {
                    triplets.push([arr[i], arr[left], arr[right]]);
                    left++;
                    while (left < right && arr[left - 1] === arr[left]) {
                        left++;
                    }
                }
            }
        }

        return triplets;
    }
}

