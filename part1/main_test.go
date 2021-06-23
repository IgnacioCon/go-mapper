package main

import "testing"

func TestCapitalizeEveryThirdAlphanumericChar(t *testing.T) {
	tests := []struct {
		name       string
		testString string
		expected   string
	}{
		{name: "string with dot", testString: "Aspiration", expected: "asPirAtiOn"},
		{name: "string with numbers", testString: "11abcdf", expected: "11AbcDf"},
		{name: "string with unicode char", testString: "ABCdr€b", expected: "abCdr€B"},
		{name: "string with spaces", testString: "I am Home.", expected: "i aM hoMe."},
		{name: "empty string", testString: "", expected: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := CapitalizeEveryThirdAlphanumericChar(tt.testString)
			if tt.expected != actual {
				t.Errorf("Expected: %s to equal Actual: %s", tt.expected, actual)
			}
		})
	}

}
