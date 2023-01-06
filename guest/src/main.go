package main

import (
	"fmt"
	"os"
	"sev-guest/src/commands"
)

func printUsage() {
	fmt.Println("sevtool --[command] [command opts]")
	fmt.Println("commands:")
	fmt.Println("  --get_report")
	fmt.Println("    Input params:")
	fmt.Println("      filename - name of the report binary file")
}

func getReport(options getReportOptions) {
	report := commands.AttestationReport{}
	reportBin, err := commands.GetReport([64]byte{}, &report)

	if err != nil {
		fmt.Println(err)
	}

	commands.WriteAttestationReport(&reportBin, options.Filename)
}

type getReportOptions struct {
	Filename     string
	DataFileName string
}

type readReportOptions struct {
	Filename string
}

type commandsOpts struct {
	PrintUsage        bool
	GetReport         bool
	GetReportOptions  getReportOptions
	ReadReport        bool
	ReadReportOptions readReportOptions
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
		getReport(cmds.GetReportOptions)
	}
}
