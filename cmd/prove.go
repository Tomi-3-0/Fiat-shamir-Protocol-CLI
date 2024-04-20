/*
Copyright © 2024 Omoniyi Caleb <calebomoniyitomiwa@gmail.com>
*/
package cmd

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

// proveCmd represents the prove command
var proveCmd = &cobra.Command{
	Use:   "prove --message 'string'",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// msg := args[0]
		msg, _ := cmd.Flags().GetString("message")

		if msg == "" {
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Enter message: ")
			input, err := reader.ReadString('\n'); if err != nil{
				fmt.Println("Error reading input", err)
				return
			}
			msg = strings.TrimSpace(input)
		}

		x := big.NewInt(0).SetBytes([]byte(msg))

		result, err := Prover(x, Secp256k1CurveOrder()); if err != nil {
			fmt.Print("Failed to Prove: ", err)
		}

		fmt.Printf("Generating proof for message: %s\n\n", msg)

		time.Sleep(2 * time.Second)
		
		result.Print(10)
		fmt.Printf("\n")

		if result.Verifier() {
			fmt.Println("Successfully generated and verified proof!")
		} else {
			fmt.Println("Verification failed")
		}

		
	},
}

func init() {
	rootCmd.AddCommand(proveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// proveCmd.PersistentFlags().String("foo", "", "A help for foo")

	proveCmd.Flags().String("message", "", "Message to prove")
}
