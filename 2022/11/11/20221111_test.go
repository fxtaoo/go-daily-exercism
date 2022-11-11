package day

import (
	"fmt"
	"testing"
)

func Test_F20221111(t *testing.T) {
	t.Parallel()

	tests := []struct {
		s1 string
		s2 string
		b  bool
	}{
		{"", "", true},
		{"1", "1", true},
		{"21", "12", true},
		{"123", "124", false},
		{"123", "124", false},
		{"1   32", "   123", true},
	}

	for _, test := range tests {
		test := test
		t.Run(fmt.Sprintf("%s %s", test.s1, test.s2), func(t *testing.T) {
			t.Parallel()
			got := F20221111_2(test.s1, test.s2)
			if got != test.b {
				t.Error(test.b, got)
			}
		})
	}
}

func BenchmarkF20221111(b *testing.B) {
	for i := 0; i < b.N; i++ {
		F20221111("1   32", "   123")
	}
}

func BenchmarkF20221111_2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		F20221111_2("1   32", "   123")
	}
}
