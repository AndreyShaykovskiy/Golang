package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Функция для обрезки строки до 40 символов с добавлением "..."
func truncateString(s string, maxLength int) string {
	if len(s) > maxLength {
		return s[:maxLength] + "..."
	}
	return s
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите выражение (например: \"100\" + \"слов\" или \"Строка\" * 3):")

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	// Регулярное выражение для поиска подстрок, выделенных двойными кавычками
	re := regexp.MustCompile(`"([^"]*)"`)

	// Найти все подстроки, соответствующие шаблону
	matches := re.FindAllStringSubmatch(input, -1)
	if len(matches) < 1 {
		fmt.Println("Ошибка. Ожидается хотя бы одна строка, выделенная двойными кавычками")
		return
	}

	// Получить первую строку из входного выражения
	firstStr := matches[0][1]

	// Проверка длины первой строки
	if len(firstStr) > 10 {
		fmt.Println("Ошибка. Длина первой строки не должна превышать 10 символов")
		return
	}

	// Определить операцию
	var op string
	if strings.Contains(input, "+") {
		op = "+"
	} else if strings.Contains(input, "-") {
		op = "-"
	} else if strings.Contains(input, "*") {
		op = "*"
	} else if strings.Contains(input, "/") {
		op = "/"
	} else {
		fmt.Println("Ошибка. Неизвестная операция")
		return
	}

	// Разбор второго операнда
	var result string
	switch op {
	case "+", "-":
		if len(matches) < 2 {
			fmt.Println("Ошибка. Ожидается две строки, выделенные двойными кавычками для операций + и -")
			return
		}
		secondStr := matches[1][1]
		// Проверка длины второй строки
		if len(secondStr) > 10 {
			fmt.Println("Ошибка. Длина второй строки не должна превышать 10 символов")
			return
		}
		switch op {
		case "+":
			result = firstStr + secondStr
		case "-":
			result = strings.Replace(firstStr, secondStr, "", -1)
		}
	case "*", "/":
		// Извлекаем числовое значение (элемент после операции)
		numStr := strings.TrimSpace(input[strings.LastIndex(input, string(op))+1:])
		num, err := strconv.Atoi(numStr)
		if err != nil {
			fmt.Println("Ошибка")
			return
		}

		// Проверка допустимого диапазона чисел
		if num < 1 || num > 10 {
			fmt.Println("Ошибка. Число должно быть в диапазоне от 1 до 10")
			return
		}

		switch op {
		case "*":
			result = strings.Repeat(firstStr, num)
		case "/":
			// Длина результата равна длине исходной строки, деленной на число
			resultLen := len(firstStr) / num
			result = firstStr[:resultLen]
		}
	}

	// Проверка длины результата
	result = truncateString(result, 40)
	fmt.Println("Ответ:", result)
}
