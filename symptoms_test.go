package cancerigo

import (
	"reflect"
	"testing"
)

func TestPrecompute(t *testing.T) {
	m := "test1 two three"
	p := precompute(m)

	if p.message != m {
		t.Errorf("precompute(%s).message == %s, want %s", m, p.message, m)
	}
	if p.length != 15 {
		t.Errorf("precompute(%s).length == %d, want %d", m, p.length, 13)
	}
	want := []string{"test1", "two", "three"}
	if !reflect.DeepEqual(p.words, want) {
		t.Errorf("precompute(%s).words == %s, want %s", m, p.words, want)
	}
	if p.wordsCount != 3 {
		t.Errorf("precompute(%s).message == %d, want %d", m, p.wordsCount, 3)
	}
}

func TestMinimumWordCount(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"sentence", 1},
		{"sentence words", 0},
		{"sentence words words", 0},
		{"", 2},
	}
	for _, c := range cases {
		got := minimumWordCount(precompute(c.in))
		if got != c.want {
			t.Errorf("minimumWordCount(%q) == %d, want %d", c.in, got, c.want)
		}
	}
}

func TestMinimumMessageLength(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"1234567890", 0},
		{"123", 0},
		{"1", 1},
		{"", 1},
	}
	for _, c := range cases {
		got := minimumMessageLength(precompute(c.in))
		if got != c.want {
			t.Errorf("minimumMessageLength(%q) == %d, want %d", c.in, got, c.want)
		}
	}
}

func TestMaximumMessageLength(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"1234567890", 0},
		{"123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890", 3},
		{"", 0},
	}
	for _, c := range cases {
		got := maximumMessageLength(precompute(c.in))
		if got != c.want {
			t.Errorf("maximumMessageLength(%q) == %d, want %d", c.in, got, c.want)
		}
	}
}

func TestCapsRatio(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"Foobar", 0},
		{"FooBar", 1},
		{"FOOBAR", 2},
		{"", 0},
	}
	for _, c := range cases {
		got := capsRatio(precompute(c.in))
		if got != c.want {
			t.Errorf("capsRatio(%q) == %d, want %d", c.in, got, c.want)
		}
	}
}

func TestEmoteCountAndRatio(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"Kappa", 2},
		{"Kappa KappaPride Keepo", 4},
		{"Kappa KappaPride Keepo Keepo KappaPride", 5},
		{"", 0},
	}
	for _, c := range cases {
		got := emoteCountAndRatio(precompute(c.in))
		if got != c.want {
			t.Errorf("emoteCountAndRatio(%q) == %d, want %d", c.in, got, c.want)
		}
	}
}
func TestEchoRatio(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"lol", 0},
		{"lol lol", 1},
		{"lol lol lol", 2},
		{"lol lol lol lol", 2},
		{"lol rekt lol rekt", 1},
		{"", 0},
	}
	for _, c := range cases {
		got := echoRatio(precompute(c.in))
		if got != c.want {
			t.Errorf("echo(%q) == %d, want %d", c.in, got, c.want)
		}
	}
}
