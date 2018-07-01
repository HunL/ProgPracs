func twoSum(nums []int, target int) []int {
    len := len(nums)
    res := []int{}
    for i := 0;i < len;i++ {
        for j := i+1;j < len;j++ {
            if nums[i]+nums[j] == target {
                res = []int{i, j}
                break
            }
        }
    }
    return res
}