package catsort

func minInt(lhs, rhs int) int {
	if lhs < rhs {
		return lhs
	}
	return rhs
}

func compareInt(lhs, rhs int) int {
	if lhs < rhs {
		return -1
	}
	if lhs > rhs {
		return +1
	}
	return 0
}

func compareUint64(lhs, rhs uint64) int {
	if lhs < rhs {
		return -1
	}
	if lhs > rhs {
		return +1
	}
	return 0
}
