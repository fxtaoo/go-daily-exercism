package day

import "testing"

func Test_F20221109(t *testing.T) {
	t.Parallel()

	tests := []struct {
		str string
		b   bool
	}{
		{"", true},
		{"1", true},
		{"11", false},
		{"123", true},
	}

	for _, test := range tests {
		test := test
		t.Run(test.str, func(t *testing.T) {
			t.Parallel()
			got := F20221109(test.str)
			if got != test.b {
				t.Error(test.b, got)
			}
		})
		t.Run(test.str, func(t *testing.T) {
			t.Parallel()
			got := F20221109_2(test.str)
			if got != test.b {
				t.Error(test.b, got)
			}
		})

		t.Run(test.str, func(t *testing.T) {
			t.Parallel()
			got := F20221109_3(test.str)
			if got != test.b {
				t.Error(test.b, got)
			}
		})

	}

}
