package kondisi

import "fmt"


func Fallth(){
	fmt.Println("\n")
	var point = 23
	switch {
		case point == 8:
		fmt.Println("muantap")

		case point > 10 && point < 35:
		fmt.Println("awesome")
		fallthrough

		case point >= 70:
		fmt.Println("kegedean")

		default:
		fmt.Println("belajar lagi")
	}
}

func Fallth1(){
	fmt.Println("\n1")
	var nasi = 72
	switch {
		case nasi == 8:
		fmt.Println("muantap")

		case nasi > 10 && nasi < 35:
		fmt.Println("awesome")

		case nasi >= 70:
		fmt.Println("kegedean")

		default:
		fmt.Println("belajar lagi")
	}
}

func Fallth2(){
	fmt.Println("\n2")
	var ikan = 28
	switch {
		case ikan == 8:
		fmt.Println("muantap")
		fallthrough

		case ikan > 10 && point < 35:
		fmt.Println("awesome")
		fallthrough
		
		case ikan >= 70:
		fmt.Println("kegedean")
		fallthrough

		default:
		fmt.Println("belajar lagi")
	}
}