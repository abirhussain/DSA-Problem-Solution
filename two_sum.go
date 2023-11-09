1func twoSum(nums []int, target int) []int {
    m := make(map[int]int)
    n := len(nums)
    for i := 0; i < n; i++ {
        requiredNumber := target - nums[i]
        if _, found := m[requiredNumber]; found {
            return []int{m[requiredNumber], i}
        }
        m[nums[i]] = i
    }
    return nil
}
