package cmd

import "github.com/spf13/cobra"

var gitCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clears the terminal",
	Long:  `Clears the terminal.`,
}
