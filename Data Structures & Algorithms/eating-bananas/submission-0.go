func minEatingSpeed(piles []int, h int) int {
	for i := 1; i <= maxSlice(piles); i++ {
		var total int
		for _, p := range piles {
			if p % i != 0 {
				total += p/i + 1
				continue
			}
			total += p / i
		}
		if total <= h {
			return i 
		}
	}
	return 1
}


func maxSlice(input []int) int {
	var max = input[0]
	for _, i := range input {
		if i > max {
			max = i
		}
	}
	return max
}
// eating rate of koko, let find the bounary
// maximin eat rate can be: max of piles => 4 b / h 
// mininum eat rate can be min of piles => 1 b /h 
// [1,2,3,4]
