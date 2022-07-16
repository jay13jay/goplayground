package twosum

import (
	"fmt"
)

func TwoSum_BF(nums []int, target int) []int {
	var ret []int

	for i := 0; i < len(nums); i++ {
		for x := 0; x < len(nums); x++ {
			if x != i && len(ret) == 0 {
				fmt.Printf("i: %v:%v\tx: %v:%v\n", i, nums[i], x, nums[x])
				fmt.Printf("i+x=%v\ttarget: %v\n\n", nums[i]+nums[x], target)

				if nums[i]+nums[x] == target {
					ret = append(ret, i)
					ret = append(ret, x)
				}
			}
		}
	}
	return ret
}

func TwoSum_HM(nums []int, target int) [2]int {
	indexMap := make(map[int]int)
	for currIndex, currNum := range nums {
		if requiredIdx, isPresent := indexMap[target-currNum]; isPresent {
			return [2]int{requiredIdx, currIndex}
		}
		indexMap[currNum] = currIndex
	}
	return [2]int{}
}
