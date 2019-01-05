// Front-end in Go server
// @jeffotoni
// 04/01/2019

package util

import (
	"bufio"
	"os"
)

// write bufio to optimization
var writer *bufio.Writer

func Print(text string) {

	writer = bufio.NewWriter(os.Stdout)
	writer.WriteString(text)
	writer.Flush()
}
