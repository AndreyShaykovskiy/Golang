package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Связываем римские числа (строки) с их значениями в арабских числах
var romanToArabicMap = map[string]int{
	"I":    1,
	"II":   2,
	"III":  3,
	"IV":   4,
	"V":    5,
	"VI":   6,
	"VII":  7,
	"VIII": 8,
	"IX":   9,
	"X":    10,
}

// Записываем пары арабских и соответствующих римских цифр для значений от 1 до 100
var arabicToRomanMap = []struct {
	Arabic int
	Roman  string
}{
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

// Функция преобразования римского числа в арабское
func romanToArabic(roman string) (int, error) {
	if value, exists := romanToArabicMap[roman]; exists {
		return value, nil
	}
	return 0, errors.New("Введите римское число от I до X")
}

// Функция преобразования арабского числа в римское
func arabicToRoman(num int) (string, error) {
	roman := ""
	for _, mapping := range arabicToRomanMap {
		for num >= mapping.Arabic {
			roman += mapping.Roman
			num -= mapping.Arabic
		}
	}
	return roman, nil
}

// Выполнение арифметической операции
func calculate(a int, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		return a / b, nil
	default:
		return 0, errors.New("Поддерживаемые арифметические действия: + - * /")
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите выражение через пробел, используя только римские или только арабские числа (Например: 2 + 7 или V * IX):")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	elements := strings.Fields(input)
	if len(elements) != 3 {
		fmt.Println("Ошибка: неверный формат ввода. Введите выражение через пробел.")
		return
	}

	aStr, op, bStr := elements[0], elements[1], elements[2]

	// Проверка ввода только римских или только арабских чисел
	isRoman := false
	if _, err := strconv.Atoi(aStr); err != nil {
		if _, err := romanToArabic(aStr); err == nil {
			isRoman = true
		} else {
			fmt.Println("Ошибка: неверный формат числа")
			return
		}
	}

	if isRoman {
		a, err := romanToArabic(aStr)
		if err != nil {
			fmt.Println("Ошибка:", err)
			return
		}
		if a < 1 || a > 10 {
			fmt.Println("Ошибка: число вне допустимого диапазона (1-10)")
			return
		}

		b, err := romanToArabic(bStr)
		if err != nil {
			fmt.Println("Ошибка:", err)
			return
		}
		if b < 1 || b > 10 {
			fmt.Println("Ошибка: число вне допустимого диапазона (1-10)")
			return
		}

		result, err := calculate(a, b, op)
		if err != nil {
			fmt.Println("Ошибка:", err)
			return
		}

		if result < 1 {
			fmt.Println("Ошибка: результат меньше единицы не может быть представлен римскими цифрами")
			return
		}

		romanResult, err := arabicToRoman(result)
		if err != nil {
			fmt.Println("Ошибка:", err)
			return
		}

		fmt.Println("Ответ:", romanResult)
	} else {
		// Проверка на ведущие нули в арабских числах
		if len(aStr) > 1 && aStr[0] == '0' {
			fmt.Println("Ошибка: числа не должны содержать ведущие нули")
			return
		}
		if len(bStr) > 1 && bStr[0] == '0' {
			fmt.Println("Ошибка: числа не должны содержать ведущие нули")
			return
		}

		a, err := strconv.Atoi(aStr)
		if err != nil {
			fmt.Println("Ошибка:", err)
			return
		}
		if a < 1 || a > 10 {
			fmt.Println("Ошибка: число вне допустимого диапазона (1-10)")
			return
		}

		b, err := strconv.Atoi(bStr)
		if err != nil {
			fmt.Println("Ошибка:", err)
			return
		}
		if b < 1 || b > 10 {
			fmt.Println("Ошибка: число вне допустимого диапазона (1-10)")
			return
		}

		result, err := calculate(a, b, op)
		if err != nil {
			fmt.Println("Ошибка:", err)
			return
		}

		fmt.Println("Ответ:", result)
	}
}
