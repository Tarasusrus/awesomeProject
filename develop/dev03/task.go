package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

// для реализации нам потребуется пакет flag

func main() {
	//sortColumn := flag.Int("k", 1, "column for sorting")
	numericSort := flag.Bool("n", false, "numeric sort")
	reverseSort := flag.Bool("r", false, "reverse sort")
	uniqSort := flag.Bool("u", false, "unique sort")

	flag.Parse() // Анализировать аргументы командной строки

	// Чтение строк
	lines := readLines()

	// Применение фильтров сортировки в зависимости от флагов
	if *numericSort {
		// Сортировка строк как чисел
		lines = sortNumeric(lines)
	}
	if *reverseSort {
		// Реверсирование строк
		lines = reverse(lines)
	}
	if *uniqSort {
		// Удаление дублирующихся строк
		lines = unique(lines)
	}

	// Вывод отсортированных строк
	for _, line := range lines {
		fmt.Print(line)
		fmt.Print(" ")
	}
}

func unique(lines []string) []string {
	// Создание map для отслеживания уникальных строк
	set := make(map[string]bool)
	uniqueLines := make([]string, 0)
	// Проходим через каждую строку
	for _, line := range lines {
		// Если строка не в set, то добавляем ее в set и в uniqueLines
		if !set[line] {
			set[line] = true
			uniqueLines = append(uniqueLines, line)
		}
	}
	// Возвращаем только уникальные строки
	return uniqueLines
}

func reverse(lines []string) []string {
	// Реверсирование строк
	for i := 0; i < len(lines)/2; i++ {
		lines[i], lines[len(lines)-1-i] = lines[len(lines)-1-i], lines[i]
	}
	return lines
}

func sortNumeric(lines []string) []string {
	sort.Slice(lines, func(i, j int) bool {
		num1, err1 := strconv.Atoi(lines[i])
		num2, err2 := strconv.Atoi(lines[j])
		if err1 != nil || err2 != nil {
			return lines[i] < lines[j]
		}
		return num1 < num2

	})
	return lines
}

func readLines() []string {
	lines := make([]string, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	return lines
}
