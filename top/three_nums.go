package top

import "sort"

//去重太恶心了
//func threeSum(nums []int) [][]int {
//	if len(nums) < 3 {
//		return nil
//	}
//	var (
//		ans [][]int
//	)
//	sort.Ints(nums)
//	for i:=0; i < len(nums)-2; i++ {
//		for j:=i+1; j < len(nums)-1; j++ {
//			for k:=j+1; k < len(nums); k++ {
//				if nums[i] + nums[j] + nums[k] == 0 {
//					if len(ans) == 0 {
//						ans = append(ans, []int{nums[i], nums[j], nums[k]})
//						continue
//					} else {
//						去重
//					}
//				}
//
//			}
//		}
//	}
//	return ans
//}

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}
	var (
		ans [][]int
	)
	sort.Ints(nums)
	for i, v := range nums {
		//防止首个数重复 case1
		if i > 0 && nums[i-1] == v {
			continue
		}
		j := i + 1
		k := len(nums) - 1
		for i < j && j < k {
			sum := v + nums[j] + nums[k]
			if sum == 0 {
				ans = append(ans, []int{v, nums[j], nums[k]})
				j++
				//防止第二个数重复 case4
				for nums[j] == nums[j-1] && j < k {
					j++
				}
			} else if sum > 0 {
				k--
			} else {
				j++
			}
		}
	}
	return ans
}
