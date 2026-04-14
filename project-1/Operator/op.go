package operator

import "fmt"

func Op() {
	var left = false
	var right = true

	const kiri = true
	const kanan = false

	fmt.Println(kiri, kanan)

	var leftandright = left && right
	fmt.Printf("left && right \t(%t)\n", leftandright)

	var leftorright = left || right
	fmt.Printf("left || right \t(%t)\n", leftorright)

	var leftReverse = !left
	fmt.Printf("!left \t\t(%t)\n", leftReverse)
}