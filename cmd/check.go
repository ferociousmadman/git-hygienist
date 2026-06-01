/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var (
	CommitOne = Commit{Hash: "a1b2c3d", Author: "ferociousmadman", Message: "feat: initial commit"}
	CommitTwo = Commit{Hash: "e5f6g7h", Author: "ferociousmadman", Message: "fix: bug in main.go"}
)

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("check called")
		fmt.Printf("Git-Hygienist: Checking your commit health...\n")
		commitObjects := []Commit{CommitOne, CommitTwo}

		for i := range len(commitObjects) {
			fmt.Printf("%v\n", commitObjects[i])
			if strings.HasPrefix(commitObjects[i].Message, "feat:") {
				println("commit type feat")
			} else if strings.HasPrefix(commitObjects[i].Message, "fix:") {
				println("commit type fix")
			} else if strings.HasPrefix(commitObjects[i].Message, "docs:") {
				println("commit type docs")
			} else {
				println("fail")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// checkCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// checkCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
