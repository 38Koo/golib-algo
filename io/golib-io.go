package golib_io

import (
	"bufio"
	"os"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)

func init() {
	// 読み込みの区切り方を指定
	scanner.Split(bufio.ScanWords)
}

func ReadSingle() int {
	scanner.Scan()
	num, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}

	return num
}

func ReadDouble() (int, int){
	return ReadSingle(), ReadSingle()
}

func ReadNs(c int) []int {
	var res []int

	for i := 1; scanner.Scan(); i++ {	
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}

		res = append(res, num)

		if i == c {
			break
		}
	}

	return res
}
