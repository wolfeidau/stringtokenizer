# stringtokenizer

This is a port of the [string tokenizer class from Java](https://docs.oracle.com/en/java/javase/11/docs/api/java.base/java/util/StringTokenizer.html) to Go. It is built on the [bufio.Scanner](https://pkg.go.dev/bufio#Scanner) in the Go standard library.

# Why

This library provides a string tokenizer for Go, which is useful for splitting a string into tokens. It can be used to build a [lexer](https://en.wikipedia.org/wiki/Lexical_analysis) or parser.

# Usage

```golang
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
```

# License

This project is released under Apache 2.0 license and is copyright [Mark Wolfe](https://www.wolfe.id.au).