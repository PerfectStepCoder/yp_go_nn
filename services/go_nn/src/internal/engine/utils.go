package engine

func FindMaxIndices(labels [][]float32) []int {
	// Массив для хранения индексов максимальных элементов
	maxIndices := make([]int, len(labels))

	// Перебираем каждый ряд массива
	for i, row := range labels {
		maxIndex := 0
		maxValue := row[0]

		// Найти индекс максимального элемента в текущем ряду
		for j, value := range row {
			if value > maxValue {
				maxValue = value
				maxIndex = j
			}
		}

		// Добавляем индекс в результат
		maxIndices[i] = maxIndex
	}

	return maxIndices
}

// Функция для сравнения двух массивов
func CompareArrays(arr1, arr2 []int) bool {
	// Проверяем, что массивы одинаковой длины
	if len(arr1) != len(arr2) {
		return false
	}

	// Сравниваем каждый элемент
	for i := range arr1 {
		if arr1[i] != arr2[i] {
			return false
		}
	}

	return true
}

// Функция для подсчета процента совпадений
func CalculateMatchPercentage(arr1, arr2 []int) float64 {
	// Проверяем, что массивы одинаковой длины
	if len(arr1) != len(arr2) {
		return 0 // Если массивы разной длины, процент совпадений невозможен
	}

	// Считаем количество совпавших элементов
	matches := 0
	for i := range arr1 {
		if arr1[i] == arr2[i] {
			matches++
		}
	}

	// Рассчитываем процент совпадений
	percentage := (float64(matches) / float64(len(arr1))) * 100
	return percentage
}
