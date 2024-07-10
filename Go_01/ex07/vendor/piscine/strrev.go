package piscine

func StrLen(s string) int {
	len := 0
	for range s {
		len++
	}
	return len
}

func StrRev(s string) string {
	runes := []rune(s)
	len := StrLen(s)
	for i := 0; i < len / 2; i++ {
		runes[i], runes[len - 1 - i] = runes[len - 1 - i], runes[i]
	}
	return string(runes)
}