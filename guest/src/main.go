package main

import (
	"fmt"
	"os"
	commands "sev-guest/src/commands"
)

func printUsage() {
	fmt.Println("sevtool --[command] [command opts]")
	fmt.Println("commands:")
	fmt.Println("	get_report")
	fmt.Println("		Input params:")
	fmt.Println("			filename - name of the report binary file")
}

func main() {
	args := os.Args[1:]
	if len(args) > 0 {
		switch args[0] {
		case "--get_report":
			report := commands.AttestationReport{}
			commands.GetReport([64]byte{}, &report)
			commands.PrintAttestationReport(&report)
		default:
			printUsage()
		}
	} else {
		fmt.Println("Command not recognized.")
	}
}
