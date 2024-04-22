package main

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

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

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	var res strings.Builder
	runes := []rune(s)

	for i := 0; i < len(runes); i++ {
		ch := runes[i]

		if unicode.IsDigit(ch) {
			return "", ErrInvalidString
		}

		repeat := 1
		if i+1 < len(runes) && unicode.IsDigit(runes[i+1]) {
			repeat, _ = strconv.Atoi(string(runes[i+1]))
			i++
		}

		res.WriteString(strings.Repeat(string(ch), repeat))
	}

	return res.String(), nil
}
