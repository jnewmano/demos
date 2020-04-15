package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

func printfString(a, b, c string, d int) string {
	s := fmt.Sprintf("%s - %s / %s %d", a, b, c, d)
	return s
}

func bufferString(a, b, c string, d int) string {

	var buff bytes.Buffer

	fmt.Fprintf(&buff, "%s - %s / %s %d", a, b, c, d)

	return buff.String()
}

func concatString(a, b, c string, d int) string {
	s := a + " - " + b + " / " + c + " " + strconv.Itoa(d)
	return s
}

func arrayString(a, b, c string, d int) string {

	ds := strconv.Itoa(d)

	arr := make([]byte, 0, len(a)+len(b)+len(c)+len(ds)+7)

	arr = append(arr, a...)

	arr = append(arr, " - "...)
	arr = append(arr, b...)

	arr = append(arr, " / "...)
	arr = append(arr, c...)

	arr = append(arr, " "...)
	arr = append(arr, ds...)
	return string(arr)
}

func builderString(a, b, c string, d int) string {

	var s strings.Builder
	s.Grow(len(a) + len(b) + len(c) + 14)

	_, _ = s.WriteString(a)
	_, _ = s.WriteString(" - ")
	_, _ = s.WriteString(b)
	_, _ = s.WriteString(" / ")
	_, _ = s.WriteString(c)
	_, _ = s.WriteString(" ")
	_, _ = s.WriteString(strconv.Itoa(d))

	return s.String()
}
