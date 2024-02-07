package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

var (
	fieldsFlag    *string = flag.String("f", "", "флаг `fields`")
	delimiterFlag *string = flag.String("d", "\t", "флаг `delimiter`")
	separatedFlag *bool   = flag.Bool("s", false, "флаг `separated`")
)

func main() {
	flag.Parse()

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()

		// Если флаг -s установлен и строка не содержит разделитель, пропустить строку
		if *separatedFlag && !strings.Contains(line, *delimiterFlag) {
			continue
		}

		// Разделите строку на фрагменты
		parts := strings.Split(line, *delimiterFlag)

		// Обработка флага -f (fields)
		if *fieldsFlag != "" {
			fields := strings.Split(*fieldsFlag, ",")
			var outputParts []string
			for _, field := range fields {
				index, err := strconv.Atoi(field)

				if err != nil {
					fmt.Fprintln(os.Stderr, "ошибка при обработке флага -f:", err)
					os.Exit(1)
				}

				if index <= 0 || index > len(parts) {
					fmt.Fprintln(os.Stderr, "ошибка: номер поля вне диапазона")
					os.Exit(1)
				}

				// Добавляем выбранное поле в outputParts
				// Мы вычитаем 1 из index, потому что срезы в Go индексируются с 0, а номера полей начинаются с 1
				outputParts = append(outputParts, parts[index-1])
			}
			// Печатаем выбранные поля
			fmt.Println(strings.Join(outputParts, *delimiterFlag))

		} else {
			// Печатаем все поля
			fmt.Println(strings.Join(parts, *delimiterFlag))
		}

		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "ошибка чтения стандартного ввода:", err)
		}
	}
