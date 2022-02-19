package array

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	s, e := -1, -1
	for i := 0; i < len(nums); i++ {
		if nums[i] == target && s == -1 {
			s = i
		}
		if s != -1 && nums[i] > target {
			e = i - 1
			break
		} else if s != -1 && nums[i] == target {
			e = i
		}
	}
	if len(nums) == 1 && s == 0 {
		e = s
	}
	return []int{s, e}
}
