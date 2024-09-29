package golibIO

import (
	"bufio"
	"io"
	"strconv"
)

func NSingle(r io.Reader) int {
	var scanner = bufio.NewScanner(r)
	scanner.Scan()

	ans, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}
	return ans
}
