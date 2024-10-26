package split

import (
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	got := Split("a/b/c", "/")
	want := []string{"a", "b", "c"}
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}

func TableTestSplit(t *testing.T) {
	tests := []struct {
		// in
		input string
		sep   string

		// out
		want []string
	}{
		{"a/b/c", "/", []string{"a", "b", "c"}},
		{"a/b/c/", "/", []string{"a", "b", "c", ""}},
		{"//", "/", []string{"", "", ""}},
		{"abc", "", []string{"abc"}},
		{"hello world", " ", []string{"hello", "world"}},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := Split(tt.input, tt.sep)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("expected: %v, got: %v", tt.want, got)
			}
		})
	}
}
