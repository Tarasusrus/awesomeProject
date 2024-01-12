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

func unpack(s string) (string, error) {
	if len(s) == 0 {
		return "", nil
	}

	if unicode.IsDigit(rune(s[0])) {
		return "", errors.New("string starts with a number")
	}

	result := ""
	runes := []rune(s)
	for i := 0; i < len(runes); {
		if runes[i] == '\\' {
			if i+1 < len(runes) {
				result += string(runes[i+1])
				i += 2
			} else {
				return "", errors.New("incorrect string")
			}
		} else if unicode.IsLetter(runes[i]) {
			result += string(runes[i])
			i++
		} else if unicode.IsDigit(rune(runes[i])) {
			if i != 0 && runes[i-1] != '\\' {
				count, _ := strconv.Atoi(string(runes[i]))
				tmp := string(runes[i-1])

				for j := 0; j < count-1; j++ {
					result += tmp
				}
			}
			i++
		} else {
			return "", errors.New("incorrect string")
		}
	}

	return result, nil
}

func main() {
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
