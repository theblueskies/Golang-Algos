package dp

// KnapSackItem denotes the model for an individual item in the knapsack
type KnapSackItem struct {
	Weight int
	value  int
}

// GetValue calculates the value of a list of items
func GetValue(r []KnapSackItem) (value int) {
	sum := 0
	for _, i := range r {
		sum += i.value
	}
	return sum
}

// GetMostValuableItems computes which list of items are the most valuable to
// be put in the knapsack
func GetMostValuableItems(arr []KnapSackItem, tmp []KnapSackItem, max int, rem int) (result []KnapSackItem) {
	if len(arr) == 0 || rem == 0 {
		return tmp
	}
	// THIS IS USED TO AVOID MEMORY LEAKS WITH SLICING
	opt1 := make([]KnapSackItem, len(arr)-1)
	for i, item := range arr {
		// opt1 := append(arr[:i], arr[i+1:]...)
		// Above append statement on Line 31 causes MEMORY LEAKS WITH SLICES
		// Solution to memory leaks is to create a fresh copy of the
		// elements from arr that are desired
		opt1 = append(opt1[:], arr[:i]...)
		copy(opt1[:], arr[i+1:])

		r1 := []KnapSackItem{}
		if item.Weight <= rem {
			// If Weight is <= remaining capacity, then add it to the recursive call
			r1 = GetMostValuableItems(opt1, append(tmp, item), max, rem-item.Weight)
		}
		r2 := GetMostValuableItems(opt1, tmp, max, rem)

		if GetValue(r1) > GetValue(r2) {
			return r1
		}
		return r2
	}
	return tmp
}
