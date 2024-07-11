package piscine

func ArrayLen[T any](array []T) int {
	arr_len := 0
	for range array {
		arr_len++
	}
	return arr_len
}

func Swap(a *int, b *int) {
	tmp := *a
	*a = *b
	*b = tmp
}

func SortIntegerTable(table []int) {
	len := ArrayLen(table)
	
	for i := 0; i < len - 1; i++ {
		for j := i + 1; j < len; j++ {
			if table[j] < table[i] {
				Swap(&table[i], &table[j])
			}
		}
	}
}