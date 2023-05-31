package zigzagconversion

import "testing"

func TestExample1(t *testing.T) {
	s := "PAYPALISHIRING"
	rows := 3
	output := "PAHNAPLSIIGYIR"

	got := convert(s, rows)
	if got != output {
		t.Errorf("convert(%v,%v)=%v, wanted: %v", s, rows, got, output)
	}
}

func TestExample2(t *testing.T) {
	s := "PAYPALISHIRING"
	rows := 4
	output := "PINALSIGYAHRPI"

	got := convert(s, rows)
	if got != output {
		t.Errorf("convert(%v,%v)=%v, wanted: %v", s, rows, got, output)
	}
}

func TestExample3(t *testing.T) {
	s := "A"
	rows := 1
	output := "A"

	got := convert(s, rows)
	if got != output {
		t.Errorf("convert(%v,%v)=%v, wanted: %v", s, rows, got, output)
	}
}
