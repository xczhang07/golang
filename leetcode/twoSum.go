func twoSum(nums []int, target int) []int {
    ret := []int{}
    if len(nums) == 0 {
        return ret
    }
    m := make(map[int]int)
    for idx, val := range nums {
        num2 := target - val
        if p_idx, ok := m[num2]; ok {
            ret = append(ret, p_idx, idx)
            break;
        }
        m[val] = idx
    }
    return ret
}
