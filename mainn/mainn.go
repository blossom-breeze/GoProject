package mainn

func Deduplicate(nums []int) []int {
	seen := make(map[int]bool)
	result := make([]int, 0)
	for _, num := range nums {
		if !seen[num] {
			seen[num] = true
			result = append(result, num)
		}
	}
	return result
}
