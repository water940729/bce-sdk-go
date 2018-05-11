package api

import "testing"

func TestValidateFunctionName(t *testing.T) {
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
			input: "brn:bce:cfc:bj:e116be6b26e2e0afad574dafe:function:function-name:1",
			err:   nil,
		},
		{
			input: "e1799855a16be6b26e2e0afad574dafe:function-name",
			err:   nil,
		},
		{
			input: "function-name",
			err:   nil,
		},
	}
	for _, tc := range cases {
		t.Run(tc.input, func(t *testing.T) {
			err := validateFunctionName(tc.input)
			if tc.err != err {
				t.Errorf("Validate %s to parse as %v, but got %v", tc.input, tc.err, err)
			}
		})
	}
}
