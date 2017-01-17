package http

import (
	"testing"
)

type testNumberPair struct {
	input    string
	expected int64
	err      bool
}

func TestValidateNumber(t *testing.T) {
	var testCases = []testNumberPair{
		{"64", 64, false},
		{"-10", -10, false},
		{"0", 0, false},
		{"+0", 0, false},
		{"-0", 0, false},
		{"f", 0, true},
		{"", 0, true},
	}

	for _, test := range testCases {
		output, err := ValidateNumber(test.input)
		if err != nil && !test.err {
			t.Errorf("Expected no error, got error on test case %+v", test)
		}
		if err == nil && output != test.expected {
			t.Errorf("Expected %d, got %d on test case %+v", test.expected, output, test)
		}
	}
}

type testPositiveNumberPair struct {
	input     string
	expected  int64
	allowZero bool
	err       bool
}

func TestValidatePositiveNumber(t *testing.T) {
	var testCases = []testPositiveNumberPair{
		{"64", 64, true, false},
		{"64", 64, false, false},
		{"-10", -10, true, true},
		{"-10", -10, false, true},
		{"0", 0, true, false},
		{"0", 0, false, true},
		{"f", 0, true, true},
		{"f", 0, false, true},
		{"", 0, true, true},
		{"", 0, false, true},
	}

	for _, test := range testCases {
		output, err := ValidatePositiveNumber(test.input, test.allowZero)
		if err != nil && !test.err {
			t.Errorf("Expected no error, got error on test case %+v", test)
		}
		if err == nil && test.err {
			t.Errorf("Expected error, got no error on test case %+v", test)
		}
		if err == nil && output != test.expected {
			t.Errorf("Expected %d, got %d on test case %+v", test.expected, output, test)
		}
	}
}

func TestValidateNegativeNumber(t *testing.T) {
	var testCases = []testNumberPair{
		{"64", 64, true},
		{"-10", -10, false},
		{"0", 0, true},
		{"+0", 0, true},
		{"-0", 0, true},
		{"f", 0, true},
		{"", 0, true},
	}

	for _, test := range testCases {
		output, err := ValidateNegativeNumber(test.input)
		if err != nil && !test.err {
			t.Errorf("Expected no error, got error on test case %+v", test)
		}
		if err == nil && output != test.expected {
			t.Errorf("Expected %d, got %d on test case %+v", test.expected, output, test)
		}
	}
}

type testDate struct {
	input  string
	layout string
	err    bool
}

func TestValidateDateFormat(t *testing.T) {
	var testCases = []testDate{
		{"17 October 2016", "2 January 2006", false},
		{"x17 October 2016x", "2 January 2006", true},
		{"32 October 2016", "2 January 2006", true},
		{"October 17 2016", "2 January 2006", true},
		{"17/10/2016", "2 January 2006", true},
	}

	for _, test := range testCases {
		_, err := ValidateDateFormat(test.input, test.layout)
		if err == nil && test.err {
			t.Errorf("Expected error, got no error on test case %+v", test)
		}
		if err != nil && !test.err {
			t.Errorf("Expected no error, got error on test case %+v", test)
		}
	}
}
