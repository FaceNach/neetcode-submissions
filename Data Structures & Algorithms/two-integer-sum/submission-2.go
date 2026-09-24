func twoSum(nums []int, target int) []int {
	m := make(map[int]int,0)
	out := make([]int,0)

    for i, n := range nums {
		r  := target - n

		_, ok:= m[r]
		if ok {
			return []int{m[r], i}
		}

		_, ok = m[n]
		if !ok{
			m[n] = i
		}
	}


	return out
}
