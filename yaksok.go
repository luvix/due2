package main

import (
	"flag"
	"fmt"
	"os/exec"
)

// SecondlyFlagSet ...
type SecondlyFlagSet struct {
	flagSet flag.FlagSet
}

// This is  SecondlyFlagSet.init
func (I *SecondlyFlagSet) init() {
	I.flagSet.String("job", "", "command what you want to schedule. If your job command is next to keyword, option is not to be needed.")
}

func main() {
	// Subcommands
	secondlyFlagSet := flag.NewFlagSet("secondly", flag.ExitOnError)
	// Verify that a subcommand has been provided
	// os.Arg[0] is the main command
	// os.Arg[1] will be the subcommand
	// if len(os.Args) < 2 {
	// 	fmt.Println("list or count subcommand is required")
	// 	os.Exit(1)
	// }

	// Switch on the subcommand
	// Parse the flags for appropriate FlagSet
	// FlagSet.Parse() requires a set of arguments to parse as input
	// os.Args[2:] will be all arguments starting after the subcommand at os.Args[1]
	flag.Parse()

	fmt.Println("flag.Args()", flag.Args())
	fmt.Println("*listCommand: ", *secondlyFlagSet)
	fmt.Println("listCommand: ", secondlyFlagSet)

	fmt.Println("flag.args")
	flagargs := flag.Args
	for i, v := range flagargs() {
		fmt.Println(i, v)
	}

	secondlyFlagSet.Parse(flag.Args()[1:])

	fmt.Println("secondlyFlagSet.args")
	for i, v := range secondlyFlagSet.Args() {
		fmt.Println(i, v)
	}

	fmt.Println("secondlyFlagSet.Parsed", secondlyFlagSet.Parsed())

	// switch flag.Arg(0) {
	// case "count":
	// 	secondlyFlagSet.Parse(flag.Args()[1:])
	// case "list":
	// 	secondlyFlagSet.Parse(flag.Args()[1:])
	// }

	for _, arg := range secondlyFlagSet.Args() {
		// var temp flag.FlagSet
		// temp.Parse(strings.Split(temp, " "))
		fmt.Println(arg)
		cmd := exec.Command("cmd", "/C", "echo", "hello")
		// var out bytes.Buffer
		// cmd.Stdin = strings.NewReader("some input")
		// cmd.Stdout = &out
		// cmd.Run()
		output, _ := cmd.Output()
		fmt.Printf("in all caps: %q", output)
	}
}
