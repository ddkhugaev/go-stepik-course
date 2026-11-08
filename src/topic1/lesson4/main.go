package main

import "fmt"

func main() {
	// Логический тип (bool)
	var f1 bool = false
	var f2 bool = true
	fmt.Println(f1, f2)

	// Строковый тип (string)
	var s1 string = "привет"
	fmt.Println(s1)
	// В Go len() возвращает количество байтов, а не символов
	fmt.Println(len(s1))
	// Латинские символы (ASCII) занимают 1 байт
	var s2 string = "hello"
	fmt.Println(len(s2))
	// Кириллица, китайские иероглифы, эмодзи занимают 2-4 байта
	var s3 string = "🚀"
	fmt.Println(s3)
	// Строковые литералы
	// Двойные кавычки - интерпретируемые строки
	var s4 string = "hello\nworld"
	// Обратные кавычки - сырые строки (raw string)
	var s5 string = `hello\nworld`
	fmt.Println(s4, s5)

	// Значения по умолчанию
	// bool: false
	// string: "" (пустая строка)
	// int: 0
	// float64: 0.0
}
