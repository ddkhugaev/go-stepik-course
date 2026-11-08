package main

import "fmt"

func main() {
	// Целые числа без знака (Unsigned Integers)
	var x1 uint8 = 8
	var x2 uint16 = 16
	var x3 uint32 = 32
	var x4 uint64 = 64
	fmt.Println(x1, x2, x3, x4)

	// Целые числа со знаком (Signed Integers)
	var x5 int8 = -8
	var x6 int16 = -16
	var x7 int32 = -32
	var x8 int64 = -64
	fmt.Println(x5, x6, x7, x8)

	// Числа с плавающей точкой (Floating-Point)
	var x9 float32 = 32.32
	var x10 float64 = 64.64
	fmt.Println(x9, x10)

	// Комплексные числа
	// complex64, complex128

	// Алиасы (псевдонимы типов)
	// uint8
	var x11 byte = 8
	// int32
	var x12 rune = -32
	fmt.Println(x11, x12)

	// Архитектурно-зависимые типы
	// 32 бита на 32-разрядной системе, 64 бита на 64-разрядной системе
	var x13 int = -64
	// аналогично для беззнакового
	var x14 uint = 64
	fmt.Println(x13, x14)
}
