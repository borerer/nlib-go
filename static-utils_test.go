package nlibgo

import "testing"

func willNotPanic() {
	// no-op
}

func willPanic() {
	panic("panic")
}

func TestSafe(t *testing.T) {
	err := Safe(willNotPanic)
	if err != nil {
		t.Fatal()
	}
	err = Safe(willPanic)
	if err == nil {
		t.Fatal("expect to see error")
	}
}
