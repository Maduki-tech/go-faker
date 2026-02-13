package parser_test

import (
	"testing"

	"github.com/maduki-tech/go-faker/internal/parser"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		fileName string
		want     map[string]any
		wantErr  bool
	}{
		// TODO: Add test cases.
		{
			name:     "Test case 1: Valid file",
			fileName: "testdata/valid_struct.go",
			want: map[string]any{
				"ID":   "int",
				"Name": "string",
				"Age":  "int",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := parser.Parse(tt.fileName)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("Parse() failed: %v", gotErr)
				}
				return
			}

			if tt.wantErr {
				t.Errorf("Parse() expected to fail but succeeded")
				return
			}

			// TODO: update the condition below to compare got with tt.want.
			if !equal(got, tt.want) {
				t.Errorf("Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}

// equal is a helper function to compare two map[string]any for equality.
func equal(a, b map[string]any) bool {
	if len(a) != len(b) {
		return false
	}
	for key, valueA := range a {
		valueB, ok := b[key]
		if !ok || valueA != valueB {
			return false
		}
	}
	return true
}
