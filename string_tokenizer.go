package stringtokenizer

import (
	"bufio"
	"io"
	"unicode/utf8"
)

type StringTokenizer struct {
	scanner           *bufio.Scanner
	includeDelimiters bool
	delimiters        []rune
}

func NewStringTokenizer(rdr io.Reader, delimiters string, include bool) *StringTokenizer {

	scanner := bufio.NewScanner(rdr)

	scanner.Split(scanTokenFunc(include, delimiters))

	return &StringTokenizer{
		delimiters:        []rune(delimiters),
		includeDelimiters: include,
		scanner:           scanner,
	}
}

func (st *StringTokenizer) HasMoreTokens() bool {
	return st.scanner.Scan()
}

func (st *StringTokenizer) NextToken() string {
	return st.scanner.Text()
}

func isDelimiter(delimiters string, r rune) bool {
	for _, d := range delimiters {
		if r == d {
			return true
		}
	}
	return false
}

func scanTokenFunc(includeDelimiters bool, delimiters string) bufio.SplitFunc {
	return func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		// Skip delimiters
		start := 0
		for width, i := 0, start; i < len(data); i += width {
			var r rune
			r, width = utf8.DecodeRune(data[i:])
			if isDelimiter(delimiters, r) {
				if includeDelimiters {
					return i + width, data[start : i+width], nil // return the delimiter
				}
				start = i + width
			} else {
				break
			}
		}
		// Scan until next delimiter
		for width, i := 0, start; i < len(data); i += width {
			var r rune
			r, width = utf8.DecodeRune(data[i:])
			if isDelimiter(delimiters, r) {
				return i, data[start:i], nil
			}
		}

		// If at EOF, return final token
		if atEOF && len(data) > start {
			return len(data), data[start:], nil
		}

		// Request more data.
		return start, nil, nil
	}
}
