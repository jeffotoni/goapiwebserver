// Go Api server
// @jeffotoni
// 2019-01-04

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
