package piscine

func StrLen(s string) int {
	len := 0
	for range s {
		len++
	}
	return len
}

func BasicAtoi(s string) int {
	runes := []rune(s)
	num := 0
	i := 0
	for runes[i] == '0' {
		i++
	}
	for runes[i] {
		num = num * 10 + (runes[i] - '0')
		i++
	}
	return num
}