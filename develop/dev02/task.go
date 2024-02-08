/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
  - "a4bc2d5e" => "aaaabccddddde"
  - "abcd" => "abcd"
  - "45" => "" (некорректная строка)
  - "" => ""

Дополнительное задание: поддержка escape - последовательностей
  - qwe\4\5 => qwe45 (*)
  - qwe\45 => qwe44444 (*)
  - qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/
package main

import (
	"errors"
	"fmt"
	"strconv"
	"unicode"
)

// Функция преобразования строки в соответствии с требованием.
func unpack(s string) (string, error) {
	// Если строка пустая, то возвращаем пустую строку
	if len(s) == 0 {
		return "", nil
	}

	// Если строка начинается с цифры - ошибка
	if unicode.IsDigit(rune(s[0])) {
		return "", errors.New("string starts with a number")
	}

	// Результирующая строка
	result := ""
	// Преобразуем входную строку в массив рун
	runes := []rune(s)
	for i := 0; i < len(runes); {
		// если текущий символ обратный слэш
		if runes[i] == '\\' {
			// Если следующий символ существует
			if i+1 < len(runes) {
				// Добавляем следующий символ к результату и увеличиваем индекс на 2.
				result += string(runes[i+1])
				i += 2
			} else {
				// Иначе возвращаем ошибку
				return "", errors.New("incorrect string")
			}
			// Если текущий символ - буква
		} else if unicode.IsLetter(runes[i]) {
			// Добавляем его к результату и увеличиваем индекс.
			result += string(runes[i])
			i++
			// Если текущий символ - числовой и предыдущий символ не обратный слэш
		} else if unicode.IsDigit(rune(runes[i])) {
			if i != 0 && runes[i-1] != '\\' {
				// Определяем число повторений и строим строку повторов, чтобы добавить к результату.
				count, _ := strconv.Atoi(string(runes[i]))
				tmp := string(runes[i-1])

				for j := 0; j < count-1; j++ {
					result += tmp
				}
			}
			// увеличиваем индекс
			i++
		} else {
			// Если ни одно из условий не выполнено, возвращаем ошибку.
			return "", errors.New("incorrect string")
		}
	}

	// Возвращаем результат и nil ошибку.
	return result, nil
}

func main() {
	// Тестовые случаи
	testCases := []string{"qwe\\4\\5", "qwe\\45", "qwe\\\\5", "a4bc2d5e"}
	for _, s := range testCases {
		res, err := unpack(s)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
		} else {
			fmt.Println(res)
		}
	}
}
