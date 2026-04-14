package konstanta

import "fmt"

func Multi() {
	const (
		a	= 1
		b	= true
		c 	= 3.4
		d	= "hello world"
	)

	fmt.Println("nice", a, b, c, d)
}