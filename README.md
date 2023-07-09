# stringtokenizer

This is a port of the [string tokenizer class from Java](https://docs.oracle.com/en/java/javase/11/docs/api/java.base/java/util/StringTokenizer.html) to Go. It is built on the [bufio.Scanner](https://pkg.go.dev/bufio#Scanner) in the Go standard library.

# Why

This library provides a string tokenizer for Go, which is useful for splitting a string into tokens. It can be used to build a [lexer](https://en.wikipedia.org/wiki/Lexical_analysis) or parser. 

The main features of this library are:

* It can return the delimiters as a token, which is great for lexical analysis
* It supports UTF-8 delimiters
* One or more delimiters can be provided which enables flexible tokenization

# Usage

Below is an example using `⌘` or `鸡` as a delimiter to illustrate handling multi-byte delimiters:

```golang
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
```

# License

This project is released under Apache 2.0 license and is copyright [Mark Wolfe](https://www.wolfe.id.au).