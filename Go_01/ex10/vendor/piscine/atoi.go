package piscine

func StrLen(s string) int {
	len := 0
	for range s {
		len++
	}
	return len
}

func Atoi(s string) int {
	len := StrLen(s)
	isNagative := false
	num := 0
	i := 0

	if i < len && s[i] == '-' || s[i] == '+' {
		if s[i] == '-' {
			isNagative = true
		}
		i++
	}
	for i < len && '0' <= s[i] && s[i] <= '9' {
		num = num * 10 + int(s[i] - '0')
		i++
	}
	if i < len {
		return 0
	}
	if isNagative {
		num = num * -1
	}
	return num
}