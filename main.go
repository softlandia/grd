package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("start grid convert\n")
	switch os.Args[1] {
	case "1":
		fmt.Printf("from 'x y z n' makes 'x y z'\n")
		CollExclude1(os.Args[2])
	case "2":
		fmt.Printf("from 'i n x y z' makes 'x y z'\n")
		CollExclude2(os.Args[2])
	}
	fmt.Printf("stop grid convert\n")
}
