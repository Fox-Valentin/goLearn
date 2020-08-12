package basic

import "testing"

func TestTriangle(t *testing.T) {
	test := []struct{ a, b, c int }{
		{3, 4, 5},
		{5, 12, 13},
	}
	for _, tt := range test {
		if actual := CalculateTriangle(tt.a, tt.b); actual != tt.c {
			t.Errorf("calculating (%d %d); got %d; expected %d;", tt.a, tt.b, actual, tt.c)
		}
	}
}

func TestSubstr(t *testing.T) {
	tests := []struct {
		str string
		ans int
	}{
		{"abcabcaabbcc", 3},
	}
	for _, tt := range tests {
		if actual := LengthOfNonRepeatingSubStr(tt.str); actual != tt.ans {
			t.Errorf("got %d for input %s; expect %d;", actual, tt.str, tt.ans)
		}
	}
}

func BenchmarkSubstr(b *testing.B) {
	s := "黑化肥会会发"
	ans := 4
	for i := 0; i < b.N; i++ {
		if actual := LengthOfNonRepeatingSubStr(s); actual != ans {
			b.Errorf("got %d for input %s; expected %d", actual, s, ans)
		}
	}
}
