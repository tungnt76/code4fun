package hashmap

func twoSum(nums []int, target int) []int {
	indexes := make(map[int]int)
	for i, num := range nums {
		if index, ok := indexes[target-num]; ok {
			return []int{index, i}
		}

		indexes[num] = i
	}

	return nil
}
