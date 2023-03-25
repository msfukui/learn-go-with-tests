package sum

func Sum(members []int) (sum int) {
	for _, member := range members {
		sum += member
	}
	return sum
}

func SumAll(membersToSum ...[]int) (sums []int) {
	for _, members := range membersToSum {
		sums = append(sums, Sum(members))
	}

	return
}

func SumAllTails(membersToSum ...[]int) (sums []int) {
	for _, members := range membersToSum {
		if len(members) == 0 {
			sums = append(sums, 0)
		} else {
			tail := members[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return
}
