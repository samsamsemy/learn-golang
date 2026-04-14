package konstanta

import "fmt"

func Multi() {
	const (
		a int     = 0
		b string  = ""
		c float64 = 0.0
		d bool    = false
	)
	fmt.Println("nice", a, b, c, d)
}

func Multimulti() {
	Multi()
}

func Wkwk() {
	Multi()
	Multimulti()
}

