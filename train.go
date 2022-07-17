package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	if a == 1 {
		fmt.Printf("%d korova", a)
	} else if 2 <= a && a <= 4 {
		fmt.Printf("%d korovy", a)
	} else if 4 <= a && a <= 20 {
		fmt.Printf("%d korov", a)
	} else if a%10 == 0 {
		fmt.Printf("%d korov", a)
	} else if a%10 == 1 {
		fmt.Printf("%d korova", a)
	} else if a%10 == 2 || a%10 == 3 || a%10 == 4 {
		fmt.Printf("%d korovy", a)
	} else {
		fmt.Printf("%d korov", a)
	}
}
