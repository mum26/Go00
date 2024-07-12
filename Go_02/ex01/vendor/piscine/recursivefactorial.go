package piscine

func RecursiveFactorial(nb int) int {
	result := 1
	if nb < 0 {
		return 0
	}
	if nb == 0 {
		return 1
	}
	if 1 < nb {
		result = nb * RecursiveFactorial(nb - 1)
	}
	if result < 0 {
		return 0
	}
	return result
}