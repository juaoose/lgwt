package roman

import (
	"testing"
	"testing/quick"
)

func TestRomanNumerals(t *testing.T) {

	cases := []struct {
		Description string
		Arabic      uint16
		Want        string
	}{
		{"1 gets I", 1, "I"},
		{"2 gets II", 2, "II"},
		{"3 gets III", 3, "III"},
		{"4 gets IV", 4, "IV"},
		{"5 gets V", 5, "V"},
		{"3990 gets MMMCMXCIX", 3999, "MMMCMXCIX"},
	}

	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			got := ConvertToRoman(test.Arabic)
			if got != test.Want {
				t.Errorf("got %q want %q", got, test.Want)
			}
		})
	}
}

func TestArabicNumbers(t *testing.T) {
	cases := []struct {
		Description string
		Roman       string
		Want        uint16
	}{
		{"I gets 1", "I", 1},
		{"II gets 2", "II", 2},
		{"III gets 3", "III", 3},
		{"MMMCMXCIX gets 3999", "MMMCMXCIX", 3999},
	}

	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			got := ConvertToArabic(test.Roman)
			if got != test.Want {
				t.Errorf("got %q want %q", got, test.Want)
			}
		})
	}
}

func TestPropertiesOfConversion(t *testing.T) {
	assertion := func(arabic uint16) bool {

		if arabic > 3999 {
			return true
		}

		roman := ConvertToRoman(arabic)
		fromRoman := ConvertToArabic(roman)
		return fromRoman == arabic
	}

	if err := quick.Check(assertion, &quick.Config{
		MaxCount: 1000,
	}); err != nil {
		t.Error("failed checks", err)
	}

}
