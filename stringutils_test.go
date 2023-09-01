package stringutils_test

import (
	"errors"
	"fmt"
	"testing"
	"unicode/utf8"

	"github.com/eminozkan/stringutils"
)

func FuzzReverse(f *testing.F) {
	tcs := []string{"hello", "uğur", "Präzisionsmeßgerät"}
	for _, tc := range tcs {
		f.Add(tc)
	}
	f.Fuzz(func(t *testing.T, original string) {
		reversed, err1 := stringutils.Reverse(original)
		if err1 != nil {
			return
		}

		reversedAgain, err2 := stringutils.Reverse(reversed)
		if err2 != nil {
			return
		}

		if original != reversedAgain {
			t.Errorf("want: %q got: %q", reversed, reversedAgain)
		}
		if utf8.ValidString(original) && !utf8.ValidString(reversed) {
			t.Errorf("reverse produced invalid UTF-8 string %q", reversed)
		}
	})
}

func TestReverse(t *testing.T) {
	tcs := map[string]struct {
		input []string
		want  []string
		err   []error
	}{
		"invalid utf8 strings": {
			input: []string{"\xC0\x80", "\xF4\x90\x80\x80", "\xFE\xA1\xA1\xA1\xA1\xA1"},
			want:  []string{"\xC0\x80", "\xF4\x90\x80\x80", "\xFE\xA1\xA1\xA1\xA1\xA1"},
			err:   []error{stringutils.ErrInvalidUTF8, stringutils.ErrInvalidUTF8, stringutils.ErrInvalidUTF8},
		},
		"none Turkish letters": {
			input: []string{"hello", "this is vigo"},
			want:  []string{"olleh", "ogiv si siht"},
			err:   []error{nil, nil},
		},
		"with Turkish letters": {
			input: []string{"uğur", "kırmızı şapka ve ÖĞRENCİ"},
			want:  []string{"ruğu", "İCNERĞÖ ev akpaş ızımrık"},
			err:   []error{nil, nil},
		},
		"with German letters": {
			input: []string{"Präzisionsmeßgerät"},
			want:  []string{"täregßemsnoisizärP"},
			err:   []error{nil, nil},
		},
	}

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			for i, in := range tc.input {
				got, err := stringutils.Reverse(in)

				if !errors.Is(err, tc.err[i]) {
					t.Errorf("want: %v; got: %v", tc.err[i], err)
				}

				if got != tc.want[i] {
					t.Errorf("want: %v; got: %v", tc.want[i], got)
				}
			}
		})
	}
}

var gs string

func BenchmarkReverse(b *testing.B) {
	var s string
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s, _ = stringutils.Reverse("merhaba dünya!")
	}

	gs = s
}

func ExampleReverse() {
	r, _ := stringutils.Reverse("vigo")
	fmt.Println(r)
	// Output: ogiv
}

func TestReplace(t *testing.T) {
	tcs := map[string]struct {
		input []string
		want  string
		err   error
	}{
		"invalid base string": {
			input: []string{"\xC0\x80", "test", "go"},
			want:  "\xC0\x80",
			err:   stringutils.ErrInvalidUTF8,
		},
		"invalid target string": {
			input: []string{"test", "\xC0\x80", "go"},
			want:  "test",
			err:   stringutils.ErrInvalidUTF8,
		},
		"invalid replacement string": {
			input: []string{"test go", "go", "\xC0\x80"},
			want:  "test go",
			err:   stringutils.ErrInvalidUTF8,
		},
		"contained substring": {
			input: []string{"Hello World", "World", "go"},
			want:  "Hello go",
			err:   nil,
		},
		"not contained substring": {
			input: []string{"Hello World", "go", "go"},
			want:  "Hello World",
			err:   stringutils.ErrInvalidSubstring,
		},
	}

	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			got, err := stringutils.Replace(tc.input[0], tc.input[1], tc.input[2])
			if !errors.Is(err, tc.err) {
				t.Errorf("unexpected error, want : %v, got : %v", tc.err, err)
			}

			if got != tc.want {
				t.Errorf("unexpected response want : %v, got : %v", tc.err, err)
			}
		})
	}

}
