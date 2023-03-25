package sum

func Sum(members []int) (sum int) {
	for _, member := range members {
		sum += member
	}
	return sum
}
