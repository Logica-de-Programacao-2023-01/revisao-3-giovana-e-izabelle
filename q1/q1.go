package q1

func TwoSum(nums []int, targe int) []int{
	Map := make(map[int]int)

	for i, num := range nums{
		complemento := target - num
		if index, ok := Map[complemento]; ok{
			return []int{index, i}
		}
		Map[num] = i
	}
	return nil
}	
