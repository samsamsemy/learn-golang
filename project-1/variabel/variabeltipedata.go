package variabel

import "fmt"

func VariabelHello() {
	var firstname string 
	firstname = "samuel"

	var lastname string 
	lastname = "wick"

	var lat string
	lat = "variabel"

	// tanpa deklarasi
	idk := "halo dunia"

	// hoo = "hidup solo!"

	// sebagai note %s untuk string %d interger
	fmt.Printf("halo %s %s %s %s!\n", firstname, lastname, lat, idk)
}