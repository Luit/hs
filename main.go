// hs is a streaming hexdump
package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	p := 0
	b := make([]byte, 1024)
	for {
		n, err := os.Stdin.Read(b)
		if err == io.EOF {
			break
		}
		if err != nil {
			if p%16 != 0 {
				fmt.Println()
			}
			fmt.Printf("Error reading from Stdin: %s\n", err.Error())
			return
		}
		for i := 0; i < n; i += 1 {
			if p%16 == 0 {
				fmt.Printf("%07x: ", p)
			}
			fmt.Printf("%02x", b[i])
			p += 1
			if p%2 == 0 {
				fmt.Print(" ")
			}
			if p%16 == 0 {
				fmt.Println()
			}
		}
	}
	if p%16 != 0 {
		fmt.Println()
	}
}
