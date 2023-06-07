package regularexpresionmatching

import "testing"

func TestExample1(t *testing.T) {
	s := "aa"
	p := "a"
	output := false
	got := isMatch(s, p)
	if got != output {
		t.Errorf("isMatch(%v,%v)=%v, wanted: %v", s, p, got, output)
	}
}

func TestExample2(t *testing.T) {
	s := "aa"
	p := "a*"
	output := true
	got := isMatch(s, p)
	if got != output {
		t.Errorf("isMatch(%v,%v)=%v, wanted: %v", s, p, got, output)
	}
}

func TestExample3(t *testing.T) {
	s := "ab"
	p := ".*"
	output := true
	got := isMatch(s, p)
	if got != output {
		t.Errorf("isMatch(%v,%v)=%v, wanted: %v", s, p, got, output)
	}
}

func TestExample4(t *testing.T) {
	s := "ab"
	p := ".*c"
	output := false
	got := isMatch(s, p)
	if got != output {
		t.Errorf("isMatch(%v,%v)=%v, wanted: %v", s, p, got, output)
	}
}

func TestExample5(t *testing.T) {
	s := "aab"
	p := "c*a*b"
	output := true
	got := isMatch(s, p)
	if got != output {
		t.Errorf("isMatch(%v,%v)=%v, wanted: %v", s, p, got, output)
	}
}
