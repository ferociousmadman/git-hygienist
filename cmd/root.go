/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

type Commit struct {
	Hash    string
	Author  string
	Message string
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "git-hygienist",
	Short: "A Go-based local CLI tool engineered for validating git commit history against Conventional Commit standards.",
	Long:  `performs pattern matching to validate standard commit types like feat:, fix:, and docs. There will be an automated Health Feedback with cobra, which will have a commit health check that evaluates individual entries and provides immediate terminal feedback on compliance status. And there's validation logic that currently processes a commitObjects collection, which should increase performance when it's going through repository history.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.git-hygienist.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
