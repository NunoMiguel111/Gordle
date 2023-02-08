package gordle

import (
	"errors"
	"strings"
	"testing"

	"golang.org/x/exp/slices"
)

func TestGameAsk(t *testing.T) {
	tt := map[string]struct {
		input string
		want  []rune
	}{
		"5 characters in english": {
			input: "Hello",
			want:  []rune("Hello"),
		},
		"5 characters in arabic": {
			input: "مرحبا ",
			want:  []rune("مرحبا "),
		},
		"5 characters in japanese": {
			input: "こんにちは",
			want:  []rune("こんにちは"),
		},
		"3 characters in japanese": {
			input: "こんに\nこんにちは",
			want:  []rune("こんにちは"),
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			g := New(strings.NewReader(tc.input))

			got := g.ask()
			if !slices.Equal(got, tc.want) {
				t.Errorf("readRunes() got = %v, want %v", string(got), string(tc.want))
			}
		})
	}
}

func TestGameValidateGuess(t *testing.T) {
	tt := map[string]struct {
		word     []rune
		expected error
	}{
		"nominal": {
			word:     []rune("GUESS"),
			expected: nil,
		},
		"too long": {
			word:     []rune("POCKET"),
			expected: errInvalidWordLength,
		},
		"too short": {
			word: []rune("RED"),
			expected: errInvalidWordLength,
		}
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			g := New(nil)

			err := g.validateGuess(tc.word)

			if !errors.Is(err, tc.expected) {
				t.Errorf("%c, expected: %q found: %q", tc.word, tc.expected, err)
			}
		})
	}
}
