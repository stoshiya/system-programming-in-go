package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	var buffer bytes.Buffer
	buffer.Write([]byte("bytes.Buffer.Write() example\n"))
	buffer.WriteString("bytes.Buffer.WriteString() example\n")
	// io.WriteString(buffer, "bytes.Buffer example\n") はビルドできない
	io.WriteString(os.Stdout, "io.WriteString() example\n")
	io.WriteString(&buffer, "io.WriteString() example\n")
	fmt.Println(buffer.String())
}
