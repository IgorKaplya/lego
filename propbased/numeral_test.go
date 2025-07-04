package propbased

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
	"testing/quick"
)

var (
	cases = []struct {
		Arabic int
		Roman  string
	}{
		{Arabic: 1, Roman: "I"},
		{Arabic: 2, Roman: "II"},
		{Arabic: 3, Roman: "III"},
		{Arabic: 4, Roman: "IV"},
		{Arabic: 5, Roman: "V"},
		{Arabic: 6, Roman: "VI"},
		{Arabic: 7, Roman: "VII"},
		{Arabic: 8, Roman: "VIII"},
		{Arabic: 9, Roman: "IX"},
		{Arabic: 10, Roman: "X"},
		{Arabic: 14, Roman: "XIV"},
		{Arabic: 18, Roman: "XVIII"},
		{Arabic: 20, Roman: "XX"},
		{Arabic: 39, Roman: "XXXIX"},
		{Arabic: 40, Roman: "XL"},
		{Arabic: 47, Roman: "XLVII"},
		{Arabic: 49, Roman: "XLIX"},
		{Arabic: 50, Roman: "L"},
		{Arabic: 100, Roman: "C"},
		{Arabic: 90, Roman: "XC"},
		{Arabic: 400, Roman: "CD"},
		{Arabic: 500, Roman: "D"},
		{Arabic: 900, Roman: "CM"},
		{Arabic: 1000, Roman: "M"},
		{Arabic: 1984, Roman: "MCMLXXXIV"},
		{Arabic: 3999, Roman: "MMMCMXCIX"},
		{Arabic: 2014, Roman: "MMXIV"},
		{Arabic: 1006, Roman: "MVI"},
		{Arabic: 798, Roman: "DCCXCVIII"},
	}
)

func TestConvertToRoman(t *testing.T) {
	for _, test := range cases {
		t.Run(fmt.Sprintf("%d", test.Arabic), func(t *testing.T) {
			var got string = ConvertToRoman(test.Arabic)

			if got != test.Roman {
				t.Errorf("got %q, want %q given %d", got, test.Roman, test.Arabic)
			}
		})
	}
}
func TestConvertToArabic(t *testing.T) {
	for _, test := range cases {
		t.Run(fmt.Sprintf("%q", test.Roman), func(t *testing.T) {
			var got int = ConvertToArabic(test.Roman)

			if got != test.Arabic {
				t.Errorf("got %d, want %d given %q", got, test.Arabic, test.Roman)
			}
		})
	}
}

func TestConversionProperties(t *testing.T) {
	var assertion = func(value int) bool {
		var romanValue = ConvertToRoman(value)
		var arabicValue = ConvertToArabic(romanValue)
		return value == arabicValue
	}

	var cfg = &quick.Config{
		Values: func(args []reflect.Value, rand *rand.Rand) {
			args[0] = reflect.ValueOf(rand.Intn(3999) + 1)
		},
	}

	var err = quick.Check(assertion, cfg)
	if err != nil {
		t.Error("Prop based assertion failed", err)
	}
}
