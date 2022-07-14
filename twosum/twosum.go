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

func TwoSum_HM(nums []int, target int) []int {
	indexMap := make(map[int]int)
	for currIndex, currNum := range nums {
		fmt.Printf("Current map: %v\n", indexMap)
		fmt.Printf("Current number: %d\n", currNum)
		if requiredIdx, isPresent := indexMap[target-currNum]; isPresent {
			return []int{requiredIdx, currIndex}
		}
		indexMap[currNum] = currIndex
	}
	return []int{}
}
