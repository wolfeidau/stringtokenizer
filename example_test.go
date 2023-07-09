package stringtokenizer

import (
	"fmt"
	"strings"
)

func ExampleStringTokenizer_NextToken() {

	input := "a,b,c"
	tokenizer := NewStringTokenizer(strings.NewReader(input), ",", false /*includeDelimiters */)

	for tokenizer.HasMoreTokens() {
		token := tokenizer.NextToken()
		fmt.Println(token)
	}

	// Output:
	// a
	// b
	// c
}
