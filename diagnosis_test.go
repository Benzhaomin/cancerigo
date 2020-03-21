package cancerigo

import (
	"testing"
)

func TestDiagnose(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"word", 1},
		{"KappaPride Keepo word word word word word", 2},
		{"CANCER Kappa PRIDE Kappa POG POG POG POG POG POG POG", 11},
		{"", 3},
	}
	for _, c := range cases {
		got := Diagnose(c.in)
		if got != c.want {
			t.Errorf("Diagnose(%q) == %d, want %d", c.in, got, c.want)
		}
	}
}

// Ryzen 2700X / 3200Mhz DDR4
// BenchmarkDiagnose-16    	  407304	      2571 ns/op	     176 B/op	       1 allocs/op
func BenchmarkDiagnose(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Diagnose("NOW Kappa THAT'S Kappa WHAT Kappa I Kappa CALL Kappa CANCER")
	}
}
