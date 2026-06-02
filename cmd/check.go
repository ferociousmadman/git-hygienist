/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

var (
	CommitOne = Commit{Hash: "a1b2c3d", Author: "ferociousmadman", Message: "feat: initial commit"}
	CommitTwo = Commit{Hash: "e5f6g7h", Author: "ferociousmadman", Message: "fix: bug in main.go"}
)

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check commit health",
	Run: func(cmd *cobra.Command, args []string) {
		commitObjects := []Commit{CommitOne, CommitTwo}

		// Initialize tabwriter: minwidth, tabwidth, padding, padchar, flags
		// tabwriter.Debug flags tabwriter to automatically draw vertical '|' lines
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', tabwriter.Debug)

		// 1. Print Top Border
		fmt.Fprintln(w, "+---------+------------------+---------+---------+")

		// 2. Print Header Row (Columns must be separated by tabs '\t')
		fmt.Fprintln(w, " HASH\t AUTHOR\t TYPE\t STATUS\t")

		// 3. Print Header Separator
		fmt.Fprintln(w, "+---------+------------------+---------+---------+")

		for _, c := range commitObjects {
			status := "FAIL"
			commitType := "unknown"

			if strings.HasPrefix(c.Message, "feat:") {
				commitType = "feat"
				status = "PASS"
			} else if strings.HasPrefix(c.Message, "fix:") {
				commitType = "fix"
				status = "PASS"
			} else if strings.HasPrefix(c.Message, "docs:") {
				commitType = "docs"
				status = "PASS"
			}

			// 4. Print Data Row (Notice the padding spaces and trailing tab)
			fmt.Fprintf(w, " %s\t %s\t %s\t %s\t\n", c.Hash, c.Author, commitType, status)
		}

		// 5. Print Bottom Border
		fmt.Fprintln(w, "+---------+------------------+---------+---------+")

		// Flush buffers the text, calculates the column widths, and prints the clean table
		w.Flush()
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)
}
