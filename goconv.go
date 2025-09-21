package goconverting

import (
	"strings"
)

func ToRussian(input string) string {
	layout := map[rune]string{
		'q': "й", 'w': "ц", 'e': "у", 'r': "к", 't': "е",
		'y': "н", 'u': "г", 'i': "ш", 'o': "щ", 'p': "з",
		'a': "ф", 's': "ы", 'd': "в", 'f': "а", 'g': "п",
		'h': "р", 'j': "о", 'k': "л", 'l': "д",
		'z': "я", 'x': "ч", 'c': "с", 'v': "м", 'b': "и",
		'n': "т", 'm': "ь",
		'{': "х", '}': "ъ", ':': "ж", '"': "э",
		'<': "б", '>': "ю",
	}
	var builder strings.Builder
	for _, ch := range input {
		if val, ok := layout[ch]; ok {
			builder.WriteString(val)
		} else {
			builder.WriteRune(ch)
		}
	}

	return builder.String()
}

func ToEnglish(input string) string {
	layout := map[rune]string{
		'й': "q", 'ц': "w", 'у': "e", 'к': "r", 'е': "t",
		'н': "y", 'г': "u", 'ш': "i", 'щ': "o", 'з': "p",
		'ф': "a", 'ы': "s", 'в': "d", 'а': "f", 'п': "g",
		'р': "h", 'о': "j", 'л': "k", 'д': "l",
		'я': "z", 'ч': "x", 'с': "c", 'м': "v", 'и': "b",
		'т': "n", 'ь': "m",
		'х': "{", 'ъ': "}", 'ж': ":", 'э': "\"",
		'б': "<", 'ю': ">",
	}

	var builder strings.Builder
	for _, ch := range input {
		if val, ok := layout[ch]; ok {
			builder.WriteString(val)
		} else {
			builder.WriteRune(ch)
		}
	}
	return builder.String()
}
