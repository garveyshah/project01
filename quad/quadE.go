package piscine
import "github.com/01-edu/z01"

func QuadE(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}else {
		for i := 1; i <= y; i++ {
			for j := 1; j <= x; j++ {
				if i == 1 && j == 1 {
					z01.PrintRune('A')
				} else if i == 1 && j == x {
					z01.PrintRune('C')
				} else if i == y && j == 1 {
					z01.PrintRune('C')
				} else if i == y && j == x {
					z01.PrintRune('A')
				} else if i == 1 || i == y || j == 1 || j == x {
					z01.PrintRune('B')
				} else {
					z01.PrintRune(' ')
				}
			}
			z01.PrintRune('\n')
		}
   }
}

