package sum

func Sum(members [5]int) (sum int) {
	for _, member := range members {
		sum += member
	}
	return sum
}
