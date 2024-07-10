package piscine

import "ft"

func PrintComb() {
	for i := '0'; i <= '7'; i++ {
		for j := i + 1; j <= '8'; j++ {
			for l := j + 1; l <= '9'; l++ {
				ft.PrintRune(i)
				ft.PrintRune(j)
				ft.PrintRune(l)
				if i != '7' || j != '8' || l != '9' {
					ft.PrintRune(',')
					ft.PrintRune(' ')
				}
			}
		}
	} 
}