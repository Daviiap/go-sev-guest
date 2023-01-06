package main

import (
	"fmt"
	"log"
	"os"
	"sev-guest/src/commands"
)

func printUsage() {
	fmt.Println("sevtool --[command] [command opts]")
	fmt.Println("commands:")
	fmt.Println("  --get_report")
	fmt.Println("    Input params:")
	fmt.Println("      filename - name of the report binary file to write")
	fmt.Println("  --read_report")
	fmt.Println("    Input params:")
	fmt.Println("      filename - name of the report binary file to read")
}

type commandsOpts struct {
	PrintUsage        bool
	GetReport         bool
	GetReportOptions  commands.GetReportOptions
	ReadReport        bool
	ReadReportOptions commands.ReadReportOptions
}

func isValidIndex(index int, length int) bool {
	return length > index
}

func parseOptions(cmdOpts *commandsOpts) {
	args := os.Args[1:]

	i := 0

	for _, s := range args {
		switch s {
		case "--get_report":
			cmdOpts.GetReport = true
			cmdOpts.GetReportOptions.Filename = "report.bin"

			if isValidIndex(i+1, len(args)) {
				cmdOpts.GetReportOptions.Filename = args[i+1]
			}
		case "--read_report":
			cmdOpts.ReadReport = true

			if isValidIndex(i+1, len(args)) {
				cmdOpts.ReadReportOptions.Filename = args[i+1]
			} else {
				log.Fatal("Invalid argument")
			}
		case "--help":
			cmdOpts.PrintUsage = true
		}

		i++
	}
}

func main() {
	cmds := commandsOpts{}
	parseOptions(&cmds)

	if cmds.PrintUsage {
		printUsage()
	} else if cmds.GetReport {
		commands.GetReportCommand(cmds.GetReportOptions)
	} else if cmds.ReadReport {
		commands.ReadReportCommand(cmds.ReadReportOptions)
	}
}
