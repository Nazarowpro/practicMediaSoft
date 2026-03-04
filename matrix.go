package main

import (
	"fmt"
	"math/rand"
	"time"
)

// UniqueMatrix представляет двумерный массив с уникальными числами
type UniqueMatrix struct {
	Rows    int
	Cols    int
	Matrix  [][]int
	Numbers map[int]bool // для отслеживания уникальности
}

// NewUniqueMatrix создает новый двумерный массив указанного размера
func NewUniqueMatrix(rows, cols int) *UniqueMatrix {
	matrix := make([][]int, rows)
	for i := range matrix {
		matrix[i] = make([]int, cols)
	}
	
	return &UniqueMatrix{
		Rows:    rows,
		Cols:    cols,
		Matrix:  matrix,
		Numbers: make(map[int]bool),
	}
}

// Generate заполняет матрицу случайными уникальными числами
func (um *UniqueMatrix) Generate(min, max int) error {
	// Проверяем, достаточно ли чисел в диапазоне
	totalCells := um.Rows * um.Cols
	if max-min+1 < totalCells {
		return fmt.Errorf("недостаточно уникальных чисел в диапазоне [%d, %d]. Нужно %d чисел", 
			min, max, totalCells)
	}
	
	// Очищаем предыдущие значения
	um.Numbers = make(map[int]bool)
	
	// Создаем срез со всеми возможными числами
	available := make([]int, 0, max-min+1)
	for i := min; i <= max; i++ {
		available = append(available, i)
	}
	
	// Перемешиваем срез
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(available), func(i, j int) {
		available[i], available[j] = available[j], available[i]
	})
	
	// Заполняем матрицу
	index := 0
	for i := 0; i < um.Rows; i++ {
		for j := 0; j < um.Cols; j++ {
			um.Matrix[i][j] = available[index]
			um.Numbers[available[index]] = true
			index++
		}
	}
	
	return nil
}

// GenerateRandomRange заполняет матрицу случайными уникальными числами с автоматическим выбором диапазона
func (um *UniqueMatrix) GenerateRandomRange() {
	totalCells := um.Rows * um.Cols
	// Выбираем диапазон в 2 раза больше необходимого для разнообразия
	maxNum := totalCells * 2
	minNum := 1
	
	// Пытаемся сгенерировать, пока не получится
	for {
		err := um.Generate(minNum, maxNum)
		if err == nil {
			break
		}
		// Увеличиваем диапазон если не хватило
		maxNum += totalCells
	}
}

// Print выводит матрицу в консоль
func (um *UniqueMatrix) Print() {
	fmt.Printf("Матрица %dx%d с уникальными числами:\n", um.Rows, um.Cols)
	
	// Находим максимальную длину числа для форматирования
	maxValue := 0
	for i := 0; i < um.Rows; i++ {
		for j := 0; j < um.Cols; j++ {
			if um.Matrix[i][j] > maxValue {
				maxValue = um.Matrix[i][j]
			}
		}
	}
	
	// Определяем ширину колонки
	width := 1
	for maxValue >= 10 {
		width++
		maxValue /= 10
	}
	width += 2 // добавляем отступы
	
	// Выводим матрицу
	for i := 0; i < um.Rows; i++ {
		for j := 0; j < um.Cols; j++ {
			fmt.Printf("%*d", width, um.Matrix[i][j])
		}
		fmt.Println()
	}
}

// ValidateUnique проверяет уникальность всех чисел в матрице
func (um *UniqueMatrix) ValidateUnique() bool {
	seen := make(map[int]bool)
	
	for i := 0; i < um.Rows; i++ {
		for j := 0; j < um.Cols; j++ {
			val := um.Matrix[i][j]
			if seen[val] {
				return false
			}
			seen[val] = true
		}
	}
	
	return len(seen) == um.Rows*um.Cols
}

// GetStats возвращает статистику по матрице
func (um *UniqueMatrix) GetStats() map[string]interface{} {
	min := um.Matrix[0][0]
	max := um.Matrix[0][0]
	sum := 0
	
	for i := 0; i < um.Rows; i++ {
		for j := 0; j < um.Cols; j++ {
			val := um.Matrix[i][j]
			if val < min {
				min = val
			}
			if val > max {
				max = val
			}
			sum += val
		}
	}
	
	avg := float64(sum) / float64(um.Rows*um.Cols)
	
	return map[string]interface{}{
		"min":        min,
		"max":        max,
		"sum":        sum,
		"average":    avg,
		"unique":     um.ValidateUnique(),
		"totalCells": um.Rows * um.Cols,
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	
	fmt.Println("=== ГЕНЕРАТОР ДВУМЕРНОГО МАССИВА С УНИКАЛЬНЫМИ ЧИСЛАМИ ===\n")
	
	// Пример 1: Матрица 3x3 с числами от 1 до 9
	fmt.Println("Пример 1: Матрица 3x3 с числами от 1 до 9")
	matrix1 := NewUniqueMatrix(3, 3)
	err := matrix1.Generate(1, 9)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
	} else {
		matrix1.Print()
		stats := matrix1.GetStats()
		fmt.Printf("Статистика: мин=%d, макс=%d, сумма=%d, среднее=%.2f, уникальность=%v\n",
			stats["min"], stats["max"], stats["sum"], stats["average"], stats["unique"])
	}
	fmt.Println()
	
	// Пример 2: Матрица 4x5 с автоматическим диапазоном
	fmt.Println("Пример 2: Матрица 4x5 со случайными уникальными числами")
	matrix2 := NewUniqueMatrix(4, 5)
	matrix2.GenerateRandomRange()
	matrix2.Print()
	stats2 := matrix2.GetStats()
	fmt.Printf("Статистика: мин=%d, макс=%d, сумма=%d, среднее=%.2f, уникальность=%v\n",
		stats2["min"], stats2["max"], stats2["sum"], stats2["average"], stats2["unique"])
	fmt.Println()
	
	// Пример 3: Матрица 5x5 с числами от 100 до 200
	fmt.Println("Пример 3: Матрица 5x5 с числами от 100 до 200")
	matrix3 := NewUniqueMatrix(5, 5)
	err = matrix3.Generate(100, 200)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
		fmt.Println("Попробуем с большим диапазоном...")
		err = matrix3.Generate(100, 224) // 5x5 = 25 чисел, диапазон 100-224 = 125 чисел
		if err == nil {
			matrix3.Print()
			stats3 := matrix3.GetStats()
			fmt.Printf("Статистика: мин=%d, макс=%d, сумма=%d, среднее=%.2f, уникальность=%v\n",
				stats3["min"], stats3["max"], stats3["sum"], stats3["average"], stats3["unique"])
		}
	} else {
		matrix3.Print()
		stats3 := matrix3.GetStats()
		fmt.Printf("Статистика: мин=%d, макс=%d, сумма=%d, среднее=%.2f, уникальность=%v\n",
			stats3["min"], stats3["max"], stats3["sum"], stats3["average"], stats3["unique"])
	}
	fmt.Println()
	
	// Демонстрация проверки уникальности
	fmt.Println("=== ПРОВЕРКА УНИКАЛЬНОСТИ ===")
	testMatrix := NewUniqueMatrix(2, 2)
	testMatrix.Matrix = [][]int{
		{1, 2},
		{3, 4},
	}
	fmt.Println("Матрица с уникальными числами:")
	testMatrix.Print()
	fmt.Printf("Уникальна: %v\n", testMatrix.ValidateUnique())
	
	testMatrix.Matrix = [][]int{
		{1, 2},
		{2, 3},
	}
	fmt.Println("\nМатрица с повторяющимся числом 2:")
	testMatrix.Print()
	fmt.Printf("Уникальна: %v\n", testMatrix.ValidateUnique())
}