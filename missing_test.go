package main

import (
	"os"
	"testing"
)

func TestMain(t *testing.T) {
	os.Stdin, _ = os.Open("test.txt")
	main()
}
