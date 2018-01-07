package main

import (
	"flag"
	"fmt"
)

func main() {
	// Subcommands
	countCommand := flag.NewFlagSet("count", flag.ExitOnError)
	listCommand := flag.NewFlagSet("list", flag.ExitOnError)

	// Count subcommand flag pointers
	// Adding a new choice for --metric of 'substring' and a new --substring flag
	countTextPtr := countCommand.String("text", "", "Text to parse. (Required)")
	countMetricPtr := countCommand.String("metric", "chars", "Metric {chars|words|lines|substring}. (Required)")
	countSubstringPtr := countCommand.String("substring", "", "The substring to be counted. Required for --metric=substring")
	countUniquePtr := countCommand.Bool("unique", false, "Measure unique values of a metric.")

	// List subcommand flag pointers
	listTextPtr := listCommand.String("text", "", "Text to parse. (Required)")
	listMetricPtr := listCommand.String("metric", "chars", "Metric <chars|words|lines>. (Required)")
	listUniquePtr := listCommand.Bool("unique", false, "Measure unique values of a metric.")

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
	fmt.Println("*listCommand: ", *listCommand)
	fmt.Println("listCommand: ", listCommand)

	fmt.Println("flag.args")
	for i, v := range flag.Args() {
		fmt.Println(i, v)
	}

	fmt.Println("countcommand.args")
	for i, v := range countCommand.Args() {
		fmt.Println(i, v)
	}

	fmt.Println("countCommand.Parsed", countCommand.Parsed())

	switch flag.Arg(0) {
	case "count":
		countCommand.Parse(flag.Args()[1:])
	case "list":
		listCommand.Parse(flag.Args()[1:])
	}

	fmt.Println("*countCommand: ", *countCommand)
	fmt.Println("*countTextPtr: ", *countTextPtr)
	fmt.Println("*countMetricPtr: ", *countMetricPtr)
	fmt.Println("*countSubstringPtr: ", *countSubstringPtr)
	fmt.Println("*countUniquePtr: ", *countUniquePtr)
	fmt.Println("*listTextPtr: ", *listTextPtr)
	fmt.Println("*listMetricPtr: ", *listMetricPtr)
	fmt.Println("*listUniquePtr: ", *listUniquePtr)
	// switch os.Args[1] {
	// case "list":
	// 	listCommand.Parse(os.Args[2:])
	// case "count":
	// 	countCommand.Parse(os.Args[2:])
	// default:
	// 	flag.PrintDefaults()
	// 	os.Exit(1)
	// }
	fmt.Println("hell world")
}
