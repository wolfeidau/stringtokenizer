package stringtokenizer

import (
	"strings"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := "a鸡b鸡c"
	tokenizer := NewStringTokenizer(strings.NewReader(input), "鸡", false)

	expectedTokens := []string{
		"a",
		"b",
		"c",
	}

	for i := 0; i < len(expectedTokens); i++ {

		tokenizer.HasMoreTokens()

		token := tokenizer.NextToken()
		if len(expectedTokens) == 0 {
			t.Errorf("Too many tokens returned")
		}

		if token != expectedTokens[i] {
			t.Errorf("Expected token %s but got %s", expectedTokens[i], token)
		}
	}

	if tokenizer.HasMoreTokens() {
		t.Errorf("Too many tokens returned")
	}
}

func TestNextToken_include_tokens(t *testing.T) {
	input := "a,b,c"
	tokenizer := NewStringTokenizer(strings.NewReader(input), ",", true)

	expectedTokens := []string{
		"a",
		",",
		"b",
		",",
		"c",
	}

	for i := 0; i < len(expectedTokens); i++ {

		tokenizer.HasMoreTokens()

		token := tokenizer.NextToken()
		if len(expectedTokens) == 0 {
			t.Errorf("Too many tokens returned")
		}

		if token != expectedTokens[i] {
			t.Errorf("Expected token %s but got %s", expectedTokens[i], token)
		}
	}

	if tokenizer.HasMoreTokens() {
		t.Errorf("Too many tokens returned")
	}
}
