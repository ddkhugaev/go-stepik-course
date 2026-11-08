package main

import "fmt"

func main() {
	// Полное объявление переменных
	// var имя тип = значение

	// Шаблоны форматирования
	// %v - универсальный вывод
	var x int = 10
	fmt.Printf("%v\n", x)
	// %T - тип переменной
	fmt.Printf("%T\n", x)
	// %d - целочисленный тип
	fmt.Printf("%d\n", x)
	// Число с плавающей точкой
	var floatVar = 3.1415
	fmt.Printf("%f\n", floatVar)
	// .Nf контролирует количество знаков после запятой
	fmt.Printf("%.2f\n", floatVar)
	// Логический тип
	var boolVar bool = true
	fmt.Printf("%t\n", boolVar)
	// Строковый тип
	var s string = "привет"
	fmt.Printf("%s\n", s)
	// %c - символ из rune
	// Rune-литералы записываются в одинарных кавычках, а строки в двойных
	var r rune = 'A'
	fmt.Printf("%c\n", r)
	// X - шестнадцатеричное представление
	fmt.Printf("U+%04X\n", r)

	// Максимальные значения
	// int8: 127
	// uint8: 255
	// int16: 32767
	// uint16: 65535
	// И так далее
}
