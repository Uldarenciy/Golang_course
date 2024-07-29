package main

import (
	"fmt"

)
func abs(n int) int {
    if n < 0 {
        return -n
    }
    return n
}
func main() {
	var a int
	fmt.Scanln(&a)
	if a < 1946 {
		//fmt.Println("Ваша дата не входит в нужный промежуток")
		panic("Ваша дата не входит в нужный промежуток")
		
	}
	switch {
		case abs(a - 1946) <= 18:
			fmt.Println("Привет, бумер!")
		case abs(a - 1965) <= 15:	
			fmt.Println("Привет, представитель X!")
		case abs(a - 1965) <= 15:	
			fmt.Println("Привет, представитель X!")
		case abs(a - 1981) <= 15:	
			fmt.Println("Привет, миллениал!")
		case abs(a - 1997) <= 15:	
			fmt.Println("Привет, зумер!")
			fallthrough
		default:
			fmt.Println("Привет, альфа!")
		}
}