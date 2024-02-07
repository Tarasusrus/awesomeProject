package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

var (
	after     *int    = flag.Int("A", 0, "флаг `after`")
	before    *int    = flag.Int("B", 0, "флаг `before`")
	context   *int    = flag.Int("C", 0, "флаг `context`")
	count     *bool   = flag.Bool("c", false, "флаг `count`")
	ignore    *bool   = flag.Bool("i", false, "флаг `ignore-case`")
	invert    *bool   = flag.Bool("v", false, "флаг `invert`")
	fixed     *bool   = flag.Bool("F", false, "флаг `fixed`")
	lineNum   *bool   = flag.Bool("n", false, "флаг `line num`")
	pattern    *string = flag.String("e", "", "флаг `pattern`")
)

func main() {
	flag.Parse()

	// Открываем файл
	file, err := os.Open(flag.Arg(0))
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		os.Exit(1)
	}
	defer file.Close()

	// Создаем сканер для чтения файла
	scanner := bufio.NewScanner(file)

	// Выполняем поиск
	for lineNum := 1; scanner.Scan(); lineNum++ {
		lineText := scanner.Text()

		// Обрабатываем флаг ignore-case
		if *ignore {
			lineText = strings.ToLower(lineText)
			*pattern = strings.ToLower(*pattern)
		}


	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка при чтении файла:", err)
	}
}
