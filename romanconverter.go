package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// RomanToArabic конвертирует римское число в арабское
func RomanToArabic(roman string) (int, error) {
	// Словарь соответствия римских цифр арабским
	romanMap := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	// Проверка на допустимые символы
	for _, char := range roman {
		if _, exists := romanMap[char]; !exists {
			return 0, fmt.Errorf("недопустимый символ: %c", char)
		}
	}

	total := 0
	prevValue := 0

	// Проходим по строке справа налево
	for i := len(roman) - 1; i >= 0; i-- {
		currentValue := romanMap[rune(roman[i])]
		
		if currentValue < prevValue {
			// Если текущее значение меньше предыдущего, вычитаем его
			total -= currentValue
		} else {
			// Иначе прибавляем
			total += currentValue
		}
		prevValue = currentValue
	}

	return total, nil
}

// ArabicToRoman конвертирует арабское число в римское (дополнительная функция)
func ArabicToRoman(num int) (string, error) {
	if num <= 0 || num > 3999 {
		return "", fmt.Errorf("число должно быть в диапазоне 1-3999")
	}

	// Словарь соответствия арабских чисел римским
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	result := ""
	remaining := num

	for i := 0; i < len(values); i++ {
		for remaining >= values[i] {
			result += symbols[i]
			remaining -= values[i]
		}
	}

	return result, nil
}

// ValidateRoman проверяет корректность римского числа
func ValidateRoman(roman string) bool {
	// Проверка на допустимые символы
	validChars := map[rune]bool{
		'I': true, 'V': true, 'X': true, 
		'L': true, 'C': true, 'D': true, 'M': true,
	}
	
	for _, char := range roman {
		if !validChars[char] {
			return false
		}
	}
	
	// Проверка на недопустимые комбинации
	invalidPatterns := []string{
		"IIII", "VV", "XXXX", "LL", "CCCC", "DD", "MMMM",
		"IVI", "IXI", "XLX", "XCX", "CDC", "CMC",
	}
	
	for _, pattern := range invalidPatterns {
		if strings.Contains(roman, pattern) {
			return false
		}
	}
	
	return true
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	
	fmt.Println("=== КОНВЕРТЕР РИМСКИХ ЦИФР В АРАБСКИЕ ===")
	fmt.Println("Доступные команды:")
	fmt.Println("  roman2arabic <римское число> - конвертировать римское в арабское")
	fmt.Println("  arabic2roman <арабское число> - конвертировать арабское в римское")
	fmt.Println("  examples - показать примеры")
	fmt.Println("  exit - выход")
	fmt.Println()

	for {
		fmt.Print("\nВведите команду: ")
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())
		
		if input == "exit" {
			fmt.Println("Программа завершена")
			break
		}
		
		if input == "examples" {
			fmt.Println("\nПримеры конвертации:")
			examples := []string{"I", "IV", "IX", "XII", "XIV", "XIX", "XXIV", "XL", "XLV", "XC", "XCV", "CD", "D", "CM", "MCMXCV", "MMXXIV"}
			for _, roman := range examples {
				if arabic, err := RomanToArabic(roman); err == nil {
					fmt.Printf("  %s = %d\n", roman, arabic)
				}
			}
			continue
		}
		
		parts := strings.Fields(input)
		if len(parts) != 2 {
			fmt.Println("Ошибка: неверный формат команды")
			fmt.Println("Используйте: roman2arabic <число> или arabic2roman <число>")
			continue
		}
		
		command := parts[0]
		value := parts[1]
		
		switch command {
		case "roman2arabic":
			value = strings.ToUpper(value)
			if !ValidateRoman(value) {
				fmt.Printf("Ошибка: '%s' не является корректным римским числом\n", value)
				continue
			}
			
			arabic, err := RomanToArabic(value)
			if err != nil {
				fmt.Printf("Ошибка: %v\n", err)
			} else {
				fmt.Printf("Результат: %s = %d\n", value, arabic)
			}
			
		case "arabic2roman":
			var num int
			_, err := fmt.Sscanf(value, "%d", &num)
			if err != nil {
				fmt.Println("Ошибка: введите целое число")
				continue
			}
			
			roman, err := ArabicToRoman(num)
			if err != nil {
				fmt.Printf("Ошибка: %v\n", err)
			} else {
				fmt.Printf("Результат: %d = %s\n", num, roman)
			}
			
		default:
			fmt.Println("Неизвестная команда. Доступные команды: roman2arabic, arabic2roman, examples, exit")
		}
	}
}