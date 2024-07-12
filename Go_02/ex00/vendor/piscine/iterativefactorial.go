package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 {
		return 0
	}
	if nb == 0 {
		return 1
	}
	result := 1
	for 0 < nb {
		result *= nb
		nb--
		if result < 0 {
			return 0
		}
	}
	return result
}