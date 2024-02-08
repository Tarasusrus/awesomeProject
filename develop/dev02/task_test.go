package main

import (
	"testing"
)

func TestUnpack(t *testing.T) {
	testCases := []struct {
		input   string
		result  string
		wantErr bool
	}{
		{input: "a4bc2d5e", result: "aaaabccddddde", wantErr: false},
		{input: "abcd", result: "abcd", wantErr: false},
		{input: "45", result: "", wantErr: true},
		{input: "", result: "", wantErr: false},
		{input: "qwe\\45", result: "qwe44444", wantErr: false},
		{input: "qwe\\\\5", result: "qwe\\\\\\\\\\", wantErr: false},
	}

	for _, tc := range testCases {
		got, err := unpack(tc.input)
		if (err != nil) != tc.wantErr {
			t.Errorf("unpack(%q) unexpected error: %v", tc.input, err)
			continue
		}
		if got != tc.result {
			t.Errorf("unpack(%q) = %q, want %q ", tc.input, got, tc.result)
		}
	}
}
