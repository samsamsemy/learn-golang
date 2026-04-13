package variabel

import "fmt"

func Underscore() {
	fmt.Println("\nini bagian underscore")
	name, _ := getName()
	fmt.Println(name)
}

func getName() (string, int) {
	return "samuel", 22
}