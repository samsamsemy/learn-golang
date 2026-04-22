package kondisi

import "fmt"

func Sarang() {
	nilai := 7
	// Nilaiditulis := "enam belas"

			if nilai < 10 {
				switch nilai {
					case 2:
					fmt.Println("dapet 2")
					case 5:
					fmt.Println("dapet 5")
					case 7:
					fmt.Println("dapet 7")
					// fallthrough buat ngecek
					case 9:
					fmt.Println("dapet 9")
				default:
					fmt.Println("yang penting dapet")
			}} else {
				if nilai == 10 {
					fmt.Println("penting dapet 10")
				} else if nilai == 13 {
					fmt.Println("km dapet 13")
				} else {
					fmt.Println("hhe")
			}
		}
	}
