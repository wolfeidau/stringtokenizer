package stringtokenizer

import (
	"fmt"
	"strings"
)

func ExampleStringTokenizer_NextToken() {

	input := "a⌘b鸡c"
	tokenizer := NewStringTokenizer(strings.NewReader(input), "⌘鸡", false /*includeDelimiters */)

	for tokenizer.HasMoreTokens() {
		token := tokenizer.NextToken()
		fmt.Println(token)
	}

	// Output:
	// a
	// b
	// c
}

func ExampleStringTokenizer_NextToken_include_delimiters() {

	input := "name=markw,age=23,cyclist=true"
	tokenizer := NewStringTokenizer(strings.NewReader(input), "=,", true /*includeDelimiters */)

	for tokenizer.HasMoreTokens() {
		token := tokenizer.NextToken()
		fmt.Println(token)
	}

	// Output:
	// name
	// =
	// markw
	// ,
	// age
	// =
	// 23
	// ,
	// cyclist
	// =
	// true
}
