package mapper_test

import (
	"mapper/mapper"
	"testing"
)

func TestMapString(t *testing.T) {
	tests := []struct {
		name       string
		expected   string
		testString string
		skips      int
	}{
		{name: "string with dot", expected: "asPirAtiOn", testString: "Aspiration", skips: 3},
		{name: "string with 1 skip", expected: "ASPIRATION", testString: "Aspiration", skips: 1},
		{name: "string with negative skips", expected: "Aspiration", testString: "Aspiration", skips: -2},
		{name: "string with length 1 and 1 skip", expected: "A", testString: "a", skips: 1},
		{name: "string with numbers", expected: "11AbcDf", testString: "11abcdf", skips: 3},
		{name: "string with unicode char", expected: "abCdr€B", testString: "ABCdr€B", skips: 3},
		{name: "string with spaces", expected: "i aM hoMe.", testString: "I am Home.", skips: 3},
		{name: "empty string", expected: "", testString: "", skips: 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			s := mapper.NewSkipString(tt.skips, tt.testString)
			mapper.MapString(&s)
			actual := s.String()

			if tt.expected != actual {
				t.Errorf("Expected: %s to equal Actual: %s", tt.expected, actual)
			}
		})
	}

}
