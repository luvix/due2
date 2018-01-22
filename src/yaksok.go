package main

import (
	"flag"
	"fmt"
)

func main() {
	// log.Fatal("Noway")
	fmt.Println("Hello world")
	h := flag.String("h", "h", "h")
	hello := flag.String("hello", "hello", "hello")
	cmd := flag.String("cmd", "cmd", "cmd")
	flag.Parse()

	fmt.Println(flag.Args())
	fmt.Println("???")
	fmt.Println(*h)
	fmt.Println(*hello)
	fmt.Println(*cmd)
	// for i, value := range flag.Args() {
	// 	fmt.Println("i: ", i, "value", value)
	// }
}
