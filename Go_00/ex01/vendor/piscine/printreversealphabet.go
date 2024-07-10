package piscine

import "ft"

func PrintReverseAlphabet() {
	for i := 'z'; 'a' <= i; i-- {
		ft.PrintRune(i)
		} 
		ft.PrintRune('\n')
}