package golibIO

import (
	"bytes"
	"testing"
)

// go test ./io
func TestNSingle(t *testing.T) {
	// モック入力を準備
	tests := []struct {
		input 	string
		expect 	int
	}{
		{"42\n", 42},
		{"-1\n", -1},
		{"0\n", 0},
		{"12345\n", 12345},
	}
	
	// テスト実行
	for _, test := range tests {
		r := bytes.NewReader([]byte(test.input))
		result := NSingle(r)
		if result != test.expect {
			t.Errorf("Expected %d but got %d", test.expect, result)
		}
	}
}
