package kondisi

import "fmt"
const point int = 5
// var point int = 5

const poinku = point

func Laper() {
	// point = 8
	// point = 8
	if point == 10 {
	fmt.Println("lulus dengan nilai sempurna")
	} else if point > 5 {
	fmt.Println("lulus")
	} else if point == 4 {
	fmt.Println("hampir lulus")
	} else {
	fmt.Printf("tidak lulus. nilai anda %d\n", point)
	}
}