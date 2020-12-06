package top

//简单粗暴的想法
//func twoSum(nums []int, target int) []int {
//	for i, v:=range nums {
//		for j, k:=range nums {
//			if v+k == target && i!= j {
//				return []int{i, j}
//			}
//		}
//	}
//	return nil
//}

//wrong
//没考虑到无序的情况
//func twoSum(nums []int, target int) []int {
//	for i,j := 0, len(nums) - 1;i < j; {
//		if nums[i] + nums[j] == target {
//			return []int{i, j}
//		}
//		if nums[i] + nums[j] < target {
//			i++
//		} else {
//			j--
//		}
//	}
//	return nil
//}

//牺牲空间换时间
func twoSum(nums []int, target int) []int {
	tmp := make(map[int]int)
	for i, v := range nums {
		tmp[v] = i
	}
	for i, v := range nums {
		if j, ok := tmp[target-v]; ok && j != i {
			return []int{i, j}
		}
	}
	return nil
}
