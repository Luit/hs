// hs is a streaming hexdump
package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	p := 0
	var printable [16]byte
	b := make([]byte, 1024)
	for {
		n, err := os.Stdin.Read(b)
		for i := 0; i < n; i += 1 {
			if p%16 == 0 {
				fmt.Printf("%07x: ", p)
			}
			fmt.Printf("%02x", b[i])
			if 0x20 <= b[i] && b[i] < 0x80 {
				printable[p%16] = b[i]
			} else {
				printable[p%16] = '.'
			}
			p += 1
			if p%2 == 0 {
				fmt.Print(" ")
			}
			if p%16 == 0 {
				fmt.Printf(" %s\n", printable)
			}
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			if p%16 != 0 {
				fmt.Printf(" %s\n", printable)
			}
			fmt.Printf("Error reading from Stdin: %s\n", err.Error())
			return
		}
	}
	if p%16 != 0 {
		for i := p % 16; i < 16; i += 1 {
			fmt.Print("  ")
			if (p+i)%2 == 0 {
				fmt.Print(" ")
			}
		}
		fmt.Printf(" %s\n", printable[0:p%16])
	}
}
