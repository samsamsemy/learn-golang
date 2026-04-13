package variabel 

import "fmt"

func Pointer() {
	fmt.Println("\nini pointer")
	name := new(string)
	fmt.Println(name)
	fmt.Println(*name)

	*name = "Samuel"

}