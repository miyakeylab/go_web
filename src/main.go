package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	f := flag.Int("flag1", 0, "flag 1")
	flag.Parse()

	fmt.Println("os.Args: ", os.Args)
	fmt.Println("compare")
	fmt.Println("flag.Args: ", flag.Args())
	if *f == 100 {
		fmt.Println("Hello")
	}
}
