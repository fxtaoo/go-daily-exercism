package day

import "testing"

func Test_F20221112(t *testing.T) {
	t.Parallel()

	tests := []struct {
		str    string
		newStr string
	}{
		{"", ""},
		{"1", ""},
		{"1 ", "%20"},
		{"1 AB我 ！ ", "%20AB%20%20"},
	}

	for _, test := range tests {
		test := test
		t.Run("", func(t *testing.T) {
			t.Parallel()
			got := F20221112(test.str)
			if got != test.newStr {
				t.Error(test.newStr, got)
			}
		})
	}
}
