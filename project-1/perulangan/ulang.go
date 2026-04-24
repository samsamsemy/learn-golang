package perulangan

import "fmt"

func Ulang (){
	for i := 0; i < 5; i++ {
		fmt.Println("Angka", i)
	}
}

func Nais (){
	var i = 0
	for i < 5 {
		fmt.Println("Angka", i, "\n")
		i++
	}
}

func Nasigoreng(){
	var i = 0
	for {
		fmt.Println("Angka", i)
		i++
		if i == 5 {
			break
		}
	}

}