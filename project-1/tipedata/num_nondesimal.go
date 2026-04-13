package tipedata

import "fmt"

func DataType() {

	// Kalau lebih dari batas ketentuan overflows
	// Unsigned integer (tidak ada negatif)
	var u8 uint8 = 255
	var u16 uint16 = 12456
	var u32 uint32 = 4294967295
	var u64 uint64 = 18446744073709551615

	fmt.Println("uint8  :", u8)
	fmt.Println("uint16 :", u16)
	fmt.Println("uint32 :", u32)
	fmt.Println("uint64 :", u64)

	// int (bisa negatif)
	var i8 int8 = 127
	var i16 int16 = 32767
	var i32 int32 = -147483647
	var i64 int64 = -9223372036854775807

	fmt.Println("int8  :", i8)
	fmt.Println("int16 :", i16)
	fmt.Println("int32 :", i32)
	fmt.Println("int64 :", i64)

	// int & uint (otomatis tergantung OS)
	var angkaInt int = 100
	var angkaUint uint = 200

	fmt.Println("int  :", angkaInt)
	fmt.Println("uint :", angkaUint)

	// byte = uint8
	var b byte = 255
	fmt.Println("byte :", b)

	// rune = int32 (untuk karakter)
	var r rune = 'H'
	fmt.Println("rune :", r)
	fmt.Println("rune (char) :", string(r), "\n")
}