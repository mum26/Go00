package piscine

func Sqrt(nb int) int {
	if nb < 0 {
		return 0
	}
	if nb == 0 || nb == 1 {
		return nb
	}
	start := 1
	end := nb / 2
	var mid int
	var square int
	for start <= end {
		mid = (start + end) / 2
		square = mid * mid;
		if square == nb {
			return mid
		}
		if square < nb {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return 0
}