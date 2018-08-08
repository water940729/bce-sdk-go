package api

import (
	"fmt"
	"testing"
)

func TestCheckFunctionName(t *testing.T) {
	cases := []struct {
		input string
		err   error
	}{
		{
			input: "brn:bce:cfc:bj:e1799855a16be6b26e2e0afad574dafe:function:function-name",
			err:   nil,
		},
		{
			input: "e1799855a16be6b26e2e0afad574dafe:function-name",
			err:   nil,
		},
		{
			input: "brn:bce:cfc:bj:e1799855a16be6b26e2e0afad574dafe:function:function-name:$LATEST",
			err:   nil,
		},
		{
			input: "$LATS",
			err:   fmt.Errorf(functionNameInvalid, "$LATS"),
		},
	}
	for _, tc := range cases {
		t.Run(tc.input, func(t *testing.T) {
			err := checkFunctionName(tc.input)
			if err == nil && tc.err != nil {
				t.Errorf("Expected err to be %v, but got nil", tc.err)
			} else if err != nil && tc.err == nil {
				t.Errorf("Expected err to be nil, but got %v", err)
			} else if err != nil && tc.err != nil && err.Error() != tc.err.Error() {
				t.Errorf("Expected err to be %v, but got %v", tc.err, err)
			}
		})
	}
}

func TestCheckPtrString(t *testing.T) {
	cases := []struct {
		input  string
		in_min int
		in_max int
		err    error
	}{
		{
			input:  "testkkk",
			in_min: 1,
			in_max: 5,
			err:    fmt.Errorf(strLenIllegal, "testkkk", 1, 5),
		},
		{
			input:  "testkkk",
			in_min: 1,
			in_max: 10,
			err:    nil,
		},
	}
	for _, tc := range cases {
		t.Run(fmt.Sprintf("%s", tc.input), func(t *testing.T) {
			err := checkPtrString(tc.input, tc.in_min, tc.in_max)
			if err == nil && tc.err != nil {
				t.Errorf("Expected err to be %v, but got nil", tc.err)
			} else if err != nil && tc.err == nil {
				t.Errorf("Expected err to be nil, but got %v", err)
			} else if err != nil && tc.err != nil && err.Error() != tc.err.Error() {
				t.Errorf("Expected err to be %v, but got %v", tc.err, err)
			}
		})
	}
}

func TestCheckPtrIntSize(t *testing.T) {
	cases := []struct {
		input  int
		in_min int
		in_max int
		err    error
	}{
		{
			input:  100,
			in_min: 0,
			in_max: 99,
			err:    fmt.Errorf(intLenIllegal, 100, 0, 99),
		},
		{
			input:  100,
			in_min: 1,
			in_max: 100,
			err:    nil,
		},
	}
	for _, tc := range cases {
		t.Run(fmt.Sprintf("%d", tc.input), func(t *testing.T) {
			err := checkPtrIntSize(tc.input, tc.in_min, tc.in_max)
			if err == nil && tc.err != nil {
				t.Errorf("Expected err to be %v, but got nil", tc.err)
			} else if err != nil && tc.err == nil {
				t.Errorf("Expected err to be nil, but got %v", err)
			} else if err != nil && tc.err != nil && err.Error() != tc.err.Error() {
				t.Errorf("Expected err to be %v, but got %v", tc.err, err)
			}
		})
	}
}

func TestCheckMemorySize(t *testing.T) {
	cases := []struct {
		input int
		err   error
	}{
		{
			input: 128,
			err:   nil,
		},
		{
			input: 64,
			err:   fmt.Errorf(intLenIllegal, 64, minMemoryLimit, maxMemoryLimit),
		},
		{
			input: 3007,
			err:   fmt.Errorf(memoryError, 3007),
		},
	}
	for _, tc := range cases {
		t.Run(fmt.Sprintf("%d", tc.input), func(t *testing.T) {
			err := checkMemorySize(tc.input)
			if err == nil && tc.err != nil {
				t.Errorf("Expected err to be %v, but got nil", tc.err)
			} else if err != nil && tc.err == nil {
				t.Errorf("Expected err to be nil, but got %v", err)
			} else if err != nil && tc.err != nil && err.Error() != tc.err.Error() {
				t.Errorf("Expected err to be %v, but got %v", tc.err, err)
			}
		})
	}
}
