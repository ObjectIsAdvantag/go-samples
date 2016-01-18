// Copyright 2015, Stève Sfartz
// Licensed under the MIT License

package util

func Reverse(s string) string {
	b := make([]byte, len(s), len(s))
	for i := 0; i < len(s); i++ {
		b[i] = s[len(s)-1-i]
	}
	return string(b)
}

