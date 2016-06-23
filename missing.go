package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var k, m, n int
	var c [10001]int
	f := bufio.NewReader(os.Stdin)
	s := bufio.NewScanner(f)
	s.Split(bufio.ScanWords)

	s.Scan()
	m, _ = strconv.Atoi(s.Text())
	for i := 0; i < m; i++ {
		s.Scan()
		k, _ = strconv.Atoi(s.Text())
		c[k]++
	}
	s.Scan()
	n, _ = strconv.Atoi(s.Text())
	for i := 0; i < n; i++ {
		s.Scan()
		k, _ = strconv.Atoi(s.Text())
		c[k]--
	}
	for k = 0; k < len(c); k++ {
		if c[k] != 0 {
			break
		}
	}
	for i := k; i < k+101; i++ {
		if c[i] != 0 {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println()
}
