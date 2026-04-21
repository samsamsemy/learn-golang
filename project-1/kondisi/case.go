package kondisi

import "fmt"

func Case(){
	var point = 5
	switch point {
		case 8:
		fmt.Println("perfect")
		case 7, 6:
		fmt.Println("awesome")
		case 5, 4:
		fmt.Println("bigger")
		default:
		fmt.Println("not bad")
	}
}

func Default(){
	var point = 10
	switch point {
		case 8:
		fmt.Println("perfect")
		case 7, 6:
		fmt.Println("awesome")
		case 5, 4:
		fmt.Println("bigger")
		default:
		fmt.Println("not bad")
	}
}
func Casestyleif(){
	var point = 35
	
	switch {
	case point == 50:
		fmt.Println("perfect")
		
	case point > 1 && point < 10:
		fmt.Println(point, " perfect")
		
	case point > 10 && point < 35:
		fmt.Println("awesome")
	case point >= 35 || point != 35:
		fmt.Println(point, " bigger")
	default:
		fmt.Println("not bad")
	}
}

