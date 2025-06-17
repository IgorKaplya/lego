package propbased

import "strings"

type romanNumeral struct {
	Value  int
	Symbol string
}

var allRomanNumerals = []romanNumeral{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ConvertToRoman(value int) string {
	var builder strings.Builder

	for _, numeral := range allRomanNumerals {
		for value >= numeral.Value {
			builder.WriteString(numeral.Symbol)
			value -= numeral.Value
		}
	}

	return builder.String()
}

func ConvertToArabic(value string) (result int) {
	for _, numeral := range allRomanNumerals {
		for strings.HasPrefix(value, numeral.Symbol) {
			result += numeral.Value
			value = strings.TrimPrefix(value, numeral.Symbol)
		}
	}
	return
}
