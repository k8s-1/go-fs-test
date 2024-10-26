package split

import (
	"reflect"
	"testing"
)

// func TestSplit(t *testing.T) {
// 	got := Split("a/b/c", "/")
// 	want := []string{"a", "b", "c"}
// 	if !reflect.DeepEqual(want, got) {
// 		t.Fatalf("expected: %v, got: %v", want, got)
// 	}
// }

func TestSplit(t *testing.T) {
	tests := []struct {
		// in
		s   string
		sep string

		// out
		want []string
	}{
		{"a/b/c", "/", []string{"a", "b", "c"}},
		{"4/3c", "/", []string{"4", "3c"}},
	}

	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			got := Split(tt.s, tt.sep)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("expected: %v, got: %v", tt.want, got)
			}
		})
	}
}
