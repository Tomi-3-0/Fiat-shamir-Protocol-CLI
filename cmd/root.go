/*
Copyright © 2024 Caleb Omoniyi <calebomoniyitomiwa@gmail.com>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "fiat-shamir",
	Short: "Fiat-shamir protocol",
	Long: `Fiat-Shamir heuristic converts interactive "proof of knowledge"
to non-interactive, letting users prove knowledge without revealing it.
Particularly useful in cryptography

This application is a tool to see the protocol in practice.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(proveCmd)
}
