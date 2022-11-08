package day

import (
	"testing"
)

func Test_F20221108(t *testing.T) {
	want := "12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ2728"
	result := F20221108()

	if want != result {
		t.Errorf(want, result)
	}
}
