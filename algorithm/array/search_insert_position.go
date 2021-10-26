package array

//searchInsert 搜索插入位置
//eg1:
//输入: nums = [1,3,5,6], target = 5
//输出: 2
//eg2:
//输入: nums = [1,3,5,6], target = 2
//输出: 1
//eg3:
//输入: nums = [1,3,5,6], target = 7
//输出: 4
//eg4:
//输入: nums = [1,3,5,6], target = 0
//输出: 0
//eg5:
//输入: nums = [1], target = 0
//输出: 0
//必须使用时间复杂度为 O(log n) 的算法
//使用二分法把时间复杂度降低到 O(log n)
func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if target < nums[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}
